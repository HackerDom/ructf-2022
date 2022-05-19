import os
import sys

N = 64
C = 10

SERVER_NAME = "team%d.ctf.hitb.org"

CLIENT_HEADER = """# wg-quick up ./game.conf

[Interface]
# you should uncomment only one pair here
# do not use the same pair for several team members
"""

CLIENT_INTERFACE_INFO = """
# PrivateKey = %s
# Address = 10.%d.%d.%d/24
"""

CLIENT_FOOTER = """
[Peer]
PublicKey = %s
Endpoint = %s:%d
PersistentKeepalive = 30
AllowedIPs = 10.60.0.0/14, 10.80.0.0/14, 10.10.10.0/24
"""

if __name__ != "__main__":
    print("I am not a module")
    sys.exit(0)

# gen client configs
os.chdir(os.path.dirname(os.path.realpath(__file__)))
try:
    os.mkdir("client_wg_dev")
except FileExistsError:
    print("Remove ./client_wg_dev dir first")
    sys.exit(1)

for i in range(1, N+1):
    srv_pub_key = open("net_keys_wg_dev/%d.srv.pubkey.txt" % i).read().strip()
    client = CLIENT_HEADER

    for c in range(1, C+1):
        clt_priv_key = open("net_keys_wg_dev/%d.clt%d.privkey.txt" % (i, c)).read().strip()
        client += CLIENT_INTERFACE_INFO % (clt_priv_key, 60+i//256, i%256, 200+c)

    client += CLIENT_FOOTER % (srv_pub_key, SERVER_NAME%i, 31000+i)

    open("client_wg_dev/%d.conf" % i, "w").write(client)

print("Finished, check ./client_wg_dev dir")
