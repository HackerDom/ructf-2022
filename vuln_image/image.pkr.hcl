packer {
  required_plugins {
    digitalocean = {
      version = ">= 1.0.0"
      source  = "github.com/hashicorp/digitalocean"
    }
  }
}

variable "api_token" {
  type = string
}

source "digitalocean" "vuln_image" {
  droplet_name  = "ructf-2022-{{timestamp}}"
  snapshot_name = "ructf-2022-{{timestamp}}"
  api_token     = var.api_token
  image         = "ubuntu-20-04-x64"
  region        = "ams3"
  size          = "s-4vcpu-8gb"
  ssh_username  = "root"
}

build {
  sources = ["source.digitalocean.vuln_image"]

  provisioner "shell" {
    inline_shebang = "/bin/sh -ex"
    environment_vars = [
      "DEBIAN_FRONTEND=noninteractive",
    ]
    inline = [
      # Wait apt-get lock
      "while ps -opid= -C apt-get > /dev/null; do sleep 1; done",
      "apt-get clean",
      # apt-get update sometime may fail
      "for i in `seq 1 3`; do apt-get update && break; sleep 10; done",

      # Wait apt-get lock
      "while ps -opid= -C apt-get > /dev/null; do sleep 1; done",

      "apt-get dist-upgrade -y -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold'",
      "for i in `seq 1 3`; do apt-get update && break; sleep 10; done",
      "apt-get upgrade -y -q -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold'",

      # Install docker and docker-compose
      "apt-get install -y -q apt-transport-https ca-certificates nfs-common",
      "curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -",
      "add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\"",
      "for i in `seq 1 3`; do apt-get update && break; sleep 10; done",
      "apt-get install -y -q docker-ce",
      "curl -L \"https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)\" -o /usr/local/bin/docker-compose",
      "chmod +x /usr/local/bin/docker-compose",
      
      # Install haveged, otherwise docker-compose may hang: https://stackoverflow.com/a/68172225/1494610
      "apt-get install -y -q haveged",

      # Add users for services
      "useradd -m -s /bin/bash ambulance",
      "useradd -m -s /bin/bash herpetophobia",
      "useradd -m -s /bin/bash kleptophobia",
      "useradd -m -s /bin/bash meds",
      "useradd -m -s /bin/bash prosopagnosia",
      "useradd -m -s /bin/bash psycho_clinic",
      "useradd -m -s /bin/bash schizophasia",
    ]
  }

  ### Ructf motd
  provisioner "shell" {
    inline = [
      "rm -rf /etc/update-motd.d/*",
    ]
  }
  provisioner "file" {
    source = "motd/ructf-banner.txt"
    destination = "/ructf-banner.txt"
  }
  provisioner "file" {
    source = "motd/00-header"
    destination = "/etc/update-motd.d/00-header"
  }
  provisioner "file" {
    source = "motd/10-help-text"
    destination = "/etc/update-motd.d/10-help-text"
  }
  provisioner "shell" {
    inline = [
      "chmod +x /etc/update-motd.d/*",
    ]
  }

  ## Onboot docker-compose run service
  provisioner "file" {
    source = "service-boot/ructf-service@.service"
    destination = "/etc/systemd/system/ructf-service@.service"
  }

  # Copy services
  provisioner "file" {
    source = "../services/ambulance/"
    destination = "/home/ambulance/"
  }

  provisioner "file" {
    source = "../services/herpetophobia/"
    destination = "/home/herpetophobia/"
  }

  provisioner "file" {
    source = "../services/kleptophobia/"
    destination = "/home/kleptophobia/"
  }

  provisioner "file" {
    source = "../services/meds/deploy/"
    destination = "/home/meds/"
  }

  provisioner "file" {
    source = "../services/prosopagnosia/deploy/"
    destination = "/home/prosopagnosia/"
  }

  provisioner "file" {
    source = "../services/psycho-clinic/"
    destination = "/home/psycho_clinic/"
  }

  provisioner "file" {
    source = "../services/schizophasia/deploy/"
    destination = "/home/schizophasia/"
  }

  # Build and run services for the first time
  provisioner "shell" {
    inline = [
      "cd ~ambulance",
      "docker-compose build",
      "cd ~herpetophobia",
      "docker-compose build",
      "cd ~kleptophobia",
      "docker-compose build",
      "cd ~meds",
      "docker-compose build",
      "cd ~prosopagnosia",
      "docker-compose build",
      "cd ~psycho_clinic",
      "docker-compose build",
      "cd ~schizophasia",
      "docker-compose build",

      "systemctl daemon-reload",
      "systemctl enable ructf-service@ambulance",
      "systemctl enable ructf-service@herpetophobia",
      "systemctl enable ructf-service@kleptophobia",
      "systemctl enable ructf-service@meds",
      "systemctl enable ructf-service@prosopagnosia",
      "systemctl enable ructf-service@psycho_clinic",
      "systemctl enable ructf-service@schizophasia",
    ]
  }

  # Fix some internal digitalocean+cloud-init scripts to be compatible with our cloud infrastructure
  provisioner "shell" {
    script = "digital_ocean_specific_setup.sh"
  }
}
