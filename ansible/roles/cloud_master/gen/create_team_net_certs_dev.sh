#!/bin/bash

#NUM="$1"

cd "$(dirname "${BASH_SOURCE[0]}")"

if [ -d net_certs_dev ]; then
 echo "Remove net_certs_dev first"
 exit 1
fi

mkdir net_certs_dev
cd net_certs_dev || exit 1

echo "This will take time"

gen_conf() {
  NUM="$1"
  ../easyrsa/easyrsa --pki-dir="team${NUM}-net" init-pki
  EASYRSA_REQ_CN="team${NUM}-net" ../easyrsa/easyrsa --batch --pki-dir="team${NUM}-net" build-ca nopass
  ../easyrsa/easyrsa --batch --pki-dir="team${NUM}-net" build-client-full "team${NUM}-client" nopass
  ../easyrsa/easyrsa --batch --pki-dir="team${NUM}-net" build-server-full "team${NUM}-server" nopass
  ../easyrsa/easyrsa --batch --pki-dir="team${NUM}-net" gen-dh
}

for i in {1..767}; do
 gen_conf "$i"
done
