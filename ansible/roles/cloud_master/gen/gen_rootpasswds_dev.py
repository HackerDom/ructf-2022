#!/usr/bin/python3

import sys
import os
import secrets

N = 64
S = 16

def genpass(n=12):
 abc = "abcdefgkmnrtxyzABCDEFGKMNRTXYZ23456789"
 return "".join([secrets.choice(abc) for i in range(n)])

os.chdir(os.path.dirname(os.path.realpath(__file__)))

try:
    os.mkdir("passwds_dev")
except FileExistsError:
    print("Remove ./passwds_dev dir first")
    sys.exit(1)


for t in range(1, N+1):
    for s in range(1, S+1):
        open("passwds_dev/team%d_serv%d_root_passwd.txt" % (t, s), "w").write(genpass()+"\n")
