# Base image preparation

Before start generate new pair of ssh keys: `ssh-keygen`

* make VM
* install Ubuntu Server from iso, use advanced installer, choose minimal installation (by F4)
* make ructfe user during process
* use separate partition for /home (because of docker) and then /var/lib/docker as bind mount to home

After boot:
* `swapoff /swapfile && rm /swapfile` (before it gets used)
* set password for root and login as one
* delete ructfe user: `deluser ructfe && rm -rf /home/ructfe`
* enable root login with password from SSH
* `ssh-copy-id` copy new keys into VM (for root)

* edit `/etc/default/grub`, set: `GRUB_CMDLINE_LINUX="net.ifnames=0 biosdevname=0"`
* edit `/etc/netplan/*.yml`, rename interface to `eth0` there
* `update-grub`
* `apt update && apt upgrade`

* `ansible-playbook -i ansible_hosts -t base image.yml`
* delete logs, poweroff, make snapshot


* check it =)
