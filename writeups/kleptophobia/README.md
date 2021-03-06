# Vuln

You are given a simple SPN. Encryption is in CBC-mode.

S-box from this service is entirely linear. You can [find](./get_traces.py) the input-output bits' relationship:

```
python3 get_traces.py
S(x)[2] = x[(1, 2)] ^ 1
S(x)[1] = x[3] ^ 1
S(x)[3] = x[5] ^ 1
S(x)[4] = x[(6, 7)] ^ 1
S(x)[7] = x[(3, 4)]
S(x)[5] = x[6]
S(x)[6] = x[0]
S(x)[0] = x[(2, 5)]
```
Every output bit depends on 1-2 input bits.

That means that attacker only need one pt-ct pair of blocks to find the initial key.

Public user info contains enough information to define the length of serialized protobuf structure. Attacker knows how many blocks was encrypted and the value of all padding bytes.

Every generated by checksystem user has only 0-2 bytes in the last block + padding. Using [sploit](../../sploits/kleptophobia/main.py) you can brute other unknown bytes and find the right encryption key (using z3 to solve linear equations).

# Fix

[Example](./fix.py)

Attacker need to know the length of serialized protobuf and at least one pair of pt-ct blocks to exploit this vulnerability. And the easiest fixing method is to append some random bytes at the start. It would expand the total length and move known pt bytes to unknown places.

When you write one field in serialized stuct twice, the deserializer would write the first entry and then rewrite it with the second one. The final structure would be successfully deserialized and would't contain the first entry.
