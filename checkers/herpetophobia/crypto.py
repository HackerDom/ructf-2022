def multiply(left_perm, right_perm):
    new_perm = [0 for _ in range(len(left_perm))]
    for i, element in enumerate(right_perm):
        new_perm[i] = left_perm[element]
    
    return new_perm


def invert(perm):
    new_perm = [0 for _ in range(len(perm))]
    for i in range(len(perm)):
        new_perm[perm[i]] =  i
    
    return new_perm


def exponentiation(perm, power):
    result = [i for i in range(len(perm))]

    if power < 0:
        perm = invert(perm)
        power = -power

    while power > 0:
        if power & 1 == 1:
            result = multiply(result, perm)
        
        perm = multiply(perm, perm)
        power  = power >> 1
    
    return result
