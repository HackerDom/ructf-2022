#!/bin/bash

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
cd "$BASE_DIR"

VM_NAME="ructfe2020-base"
OUTPUT_IMAGE="images/ructfe2020-deploy_$(date +"%Y-%m-%d_%H-%M-%S").ova"
LATEST_IMAGE="images/ructfe2020-deploy.ova"
SSH_PORT=2222
SSH_HOST=127.0.0.1

vboxmanage snapshot "$VM_NAME" restore "base_image"

# Start vm
VBoxManage modifyvm "$VM_NAME" --natpf1 "deploy,tcp,127.0.0.1,$SSH_PORT,,22"
VBoxManage modifyvm "$VM_NAME" --memory 6144 --cpus 4
vboxmanage startvm "$VM_NAME" --type headless

echo "Waiting SSH up"

SSH_OPTS="-o StrictHostKeyChecking=no -o CheckHostIP=no -o NoHostAuthenticationForLocalhost=yes"
SSH_OPTS="$SSH_OPTS -o BatchMode=yes -o LogLevel=ERROR -o UserKnownHostsFile=/dev/null"
SSH_OPTS="$SSH_OPTS -o ConnectTimeout=2 -o User=root -i keys/id_rsa"
SSH="ssh $SSH_OPTS -p $SSH_PORT $SSH_HOST"

while ! $SSH echo "SSH CONNECTED"; do
    echo "Still waiting"
    sleep 1
done

# Deploy updates
ansible-playbook -i ansible_hosts --key-file keys/id_rsa --skip-tags "start_service" image.yml

## for debug
#exit 1

# Power off VM
$SSH poweroff || echo 'OK'
while VBoxManage list runningvms | grep -q "$VM_NAME"; do
    echo "Waiting for vm stop"
    sleep 1.2
done


echo "Deleting port-forwarding for deploy"
# note(@xelez): there is some strange race where virtual machine is locked... so we just retry in 2 seconds)
VBoxManage modifyvm "$VM_NAME" --natpf1 delete deploy || (sleep 2 && VBoxManage modifyvm "$VM_NAME" --natpf1 delete deploy)


echo "VM stopped, exporting to $OUTPUT_IMAGE"
VBoxManage export "$VM_NAME" -o "$OUTPUT_IMAGE"

echo "Changing latest symlink"
ln -f -s $(realpath "$OUTPUT_IMAGE") "$LATEST_IMAGE"

echo "Done"
