# RuCTF 2022 | Ambulance

## Vulnerability

Calling `free()` on arbitrary pointer. There is a bug in [fastecdsa](https://github.com/AntonKueltz/fastecdsa) python library.

### Function `curvemath_mul` in [curvemath.c](https://github.com/AntonKueltz/fastecdsa/blob/master/src/curveMath.c#L210)

fastecdsa can perform elliptic point multiplication (`d * P`), which uses `curvemath_mul` function internally:

```c
static PyObject * curvemath_mul(PyObject *self, PyObject *args) {
    char * x, * y, * d, * p, * a, * b, * q, * gx, * gy;

    if (!PyArg_ParseTuple(args, "sssssssss", &x, &y, &d, &p, &a, &b, &q, &gx, &gy)) {
        return NULL;
    }

    PointZZ_p result;
    mpz_t scalar;
    mpz_init_set_str(scalar, d, 10);
    CurveZZ_p * curve = buildCurveZZ_p(p, a, b, q, gx, gy, 10);;

    PointZZ_p * point = buildPointZZ_p(x, y, 10);
    pointZZ_pMul(&result, point, scalar, curve);
    destroyPointZZ_p(point);
    destroyCurveZZ_p(curve);

    char * resultX = mpz_get_str(NULL, 10, result.x);
    char * resultY = mpz_get_str(NULL, 10, result.y);
    mpz_clears(result.x, result.y, scalar, NULL);

    PyObject * ret = Py_BuildValue("ss", resultX, resultY);
    free(resultX);
    free(resultY);
    return ret;
}
```

- One can notice `PointZZ_p result` variable, which is not initialized properly.
- Later this variable is cleared by `mpz_clears(result.x, result.y, scalar, NULL);`.
- `mpz_clears` internally calls `free()` on the chunk contains `mpz_t`.
- If we could save `result` uninitialized, we could call `free()` on fake address.

### How to remain `result` unitialized

Look at `pointZZ_pMul()` function which uses `result` variable:

```c
void pointZZ_pMul(PointZZ_p * rop, const PointZZ_p * point, const mpz_t scalar, const CurveZZ_p * curve) {
    // handle the identity element
    if(pointZZ_pIsIdentityElement(point)) {
        return pointZZ_pSetToIdentityElement(rop);
    }

    PointZZ_p R0, R1, tmp;
    mpz_inits(R1.x, R1.y, tmp.x, tmp.y, NULL);
    mpz_init_set(R0.x, point->x);
    mpz_init_set(R0.y, point->y);
    pointZZ_pDouble(&R1, point, curve);

    int dbits = mpz_sizeinbase(scalar, 2), i;

    for(i = dbits - 2; i >= 0; i--) {
        if(mpz_tstbit(scalar, i)) {
            mpz_set(tmp.x, R0.x);
            mpz_set(tmp.y, R0.y);
            pointZZ_pAdd(&R0, &R1, &tmp, curve);
            mpz_set(tmp.x, R1.x);
            mpz_set(tmp.y, R1.y);
            pointZZ_pDouble(&R1, &tmp, curve);
        }
        else {
            mpz_set(tmp.x, R1.x);
            mpz_set(tmp.y, R1.y);
            pointZZ_pAdd(&R1, &R0, &tmp, curve);
            mpz_set(tmp.x, R0.x);
            mpz_set(tmp.y, R0.y);
            pointZZ_pDouble(&R0, &tmp, curve);
        }
    }

    mpz_init_set(rop->x, R0.x);
    mpz_init_set(rop->y, R0.y);
    mpz_clears(R0.x, R0.y, R1.x, R1.y, tmp.x, tmp.y, NULL);
}
```

One could see that `result` still remains unitialized if the first condition holds:

```c
if(pointZZ_pIsIdentityElement(point)) {
    return pointZZ_pSetToIdentityElement(rop);
}
```

And `pointZZ_pIsIdentityElement` function:

```c
int pointZZ_pIsIdentityElement(const PointZZ_p * op) {
    return mpz_cmp_ui(op->x, 0) == 0 && mpz_cmp_ui(op->y, 0) == 0 ? 1 : 0;
}
```

So, if the point has coordinates `(0, 0)`, it will be considered as identity element, and `result` variable will remain unitialized.

### Writing address inside `result` variable

We need to write our controlled address into stack.

After some searching I've found [gmpy2](https://github.com/aleaxit/gmpy) python library, which performs stack allocation (`alloca()`) in `GMPy_MPZ_To_Binary` function, located in [gmpy2_binary.c](https://github.com/aleaxit/gmpy/blob/9127042b1240f560274af96fe6a187dc9b33d9a7/src/gmpy2_binary.c#L248) file:

```c
static PyObject *
GMPy_MPZ_To_Binary(MPZ_Object *self)
{
    size_t size = 2;
    int sgn;
    char *buffer;
    PyObject *result;

    sgn = mpz_sgn(self->z);
    if (sgn == 0) {
        TEMP_ALLOC(buffer, size);
        buffer[0] = 0x01;
        buffer[1] = 0x00;
        goto done;
    }

    size = ((mpz_sizeinbase(self->z, 2) + 7) / 8) + 2;

    TEMP_ALLOC(buffer, size);
    buffer[0] = 0x01;
    if (sgn > 0)
        buffer[1] = 0x01;
    else
        buffer[1] = 0x02;
    mpz_export(buffer+2, NULL, -1, sizeof(char), 0, 0, self->z);

  done:
    result = PyBytes_FromStringAndSize(buffer, size);
    TEMP_FREE(buffer, size);
    return result;
}
```

It uses `TEMP_ALLOC` macro, defined in [gmpy2.h](https://github.com/aleaxit/gmpy/blob/9127042b1240f560274af96fe6a187dc9b33d9a7/src/gmpy2.h#L410):

```c
#define TEMP_ALLOC(B, S)     \
  if(S < ALLOC_THRESHOLD) {  \
      B = alloca(S);         \
  } else {                   \
      if(!(B = malloc(S))) { \
          PyErr_NoMemory();  \
          return NULL;       \
      }                      \
  }
#define TEMP_FREE(B, S) if(S >= ALLOC_THRESHOLD) free(B)
```

`GMPy_MPZ_To_Binary` is used when serializing `gmpy2.mpz` python object, and raw binary data will be written onto stack:

```python
gmpy2.to_binary(gmpy2.mpz(number))
```

### Exploitation

1. Call `gmpy2.to_binary` to write controlled address onto stack
2. Make point `(0, 0)` and multiply it by something
3. During the multiplication process `free()` on controlled address will be triggered

### Spawning shell with fake chunk

The rest is just a heap feng shui. For example:

1. The service interface allows an attacker to spawn `bytes` object. The allocator will reuse the same chunk for two objects, which have equal size. Using this primitive we can rewrite data at some location: repeatedly spawn objects of the same size. When someone frees a fake chunk inside an existing object, there will be **two overlapped chunks**.

2. The service interface allows an attacker to create `list` object. The attacker could create `list` and `bytes` overlapped and rewrite few pointers of `list` elements.

3. Also, the attacker could create fake python object, which base class is also fake python object. If he set object's **deallocator function** to controlled address, this address will be called when object destroys.

### How to leak address

The service contains the class `Disease` and the object of this class:

```python
class Disease:
    pass

...

NoDisease = Disease()
```

Python default `__repr__` leaks the address of the object:

```
>>> print(NoDisease)
<__main__.Disease at 0x1037392a0>
```

By default this object is stored in `User` structure, but it could be bypassed if the attacker updates disease in the other session. Then he could calculate offset to any python heap chunk.

### How to make `(0, 0)` point

Look at the `SecureCurve`:

```python
SecureCurve = Curve(
    p  = 0xa0fca03a870f6e3fc52aeef0d61f198fddc7a2c6bd414b3e5a1afc5a4a82009d,
    a  = 0x3458be7671950c6b01bed2734056c9217012fd1f07ee085afd504b412061e63c,
    b  = 0x0,
    q  = 0xa0fca03a870f6e3fc52aeef0d61f19915ca241a1b2e1cb33cb1434415514a902,
    gx = 0x6a0ea6b596c2adb773a821e9c6799a0e8ab03e355560a64ac1eecb6df8bd92ba,
    gy = 0x9e337b7d04c686771d18cd12a9b5174cb5b134be7ab09176c418bce4ff265de9,
    oid = b's\xee\xccur\xee',
    name = 'SecureCurve',
)
```

It has `b = 0`, therefore the curve contains `(0, 0)` point. The attacker could set his public key to `(0, 0)` using `q // 2` as private key (`recovery_key`):

```
>>> SecureCurve.G * (SecureCurve.q // 2)
<POINT AT INFINITY>
```

### The exploit TLDR

1. Leak address of `Disease()` object and calculate offset to some _chunk_
2. Set public key to `(0, 0)`
3. Create a _fake chunk_ inside the controlled _chunk_
4. Write the address on the stack using `gmpy2.to_binary`
5. Free the _fake chunk_ using point muptiplication (verifying password)
6. Create list at _fake chunk_
7. Create a fake object with deallocator function and rewrite some list pointers using _chunk_
8. Destroy this object and drop shell

[Example sploit.py](/sploits/ambulance/sploit.py)

## Fix

Change anything that affects python memory layout. For example, switch from Python 3.9 to Python 3.10.

## Note

This solution may not work on arbitrary machine and/or patched service, but it was tested on vuln image and works perfectly.
