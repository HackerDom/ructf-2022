import os
import sys

N = 64
C = 10

SERVER_HEADER = """[Interface]
PrivateKey = %s
Address = 10.%d.%d.254/24
ListenPort = %d
"""

SERVER_PEER = """
[Peer]
PublicKey = %s
AllowedIPs = 10.%d.%d.%d/32
PersistentKeepalive=30
"""

if __name__ != "__main__":
    print("I am not a module")
    sys.exit(0)

# gen client configs
os.chdir(os.path.dirname(os.path.realpath(__file__)))
try:
    os.mkdir("server_wg_dev")
except FileExistsError:
    print("Remove ./server_wg_dev dir first")
    sys.exit(1)

for i in range(1, N+1):
    srv_priv_key = open("net_keys_wg_dev/%d.srv.privkey.txt" % i).read().strip()
    server = SERVER_HEADER % (srv_priv_key, 60+i//256, i%256, 31000+i)

    for c in range(1, C+1):
        clt_pub_key = open("net_keys_wg_dev/%d.clt%d.pubkey.txt" % (i, c)).read().strip()

        server += SERVER_PEER % (clt_pub_key, 60+i//256, i%256, 200+c)

    open("server_wg_dev/%d.conf" % i, "w").write(server)

print("Finished, check ./server_wg_dev dir")
