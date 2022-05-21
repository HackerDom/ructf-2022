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
  image         = "fedora-36-x64"
  region        = "ams3"
  size          = "s-4vcpu-8gb"
  ssh_username  = "root"
}

build {
  sources = ["source.digitalocean.vuln_image"]

  provisioner "shell" {
    inline_shebang = "/bin/sh -ex"
    inline = [
      "dnf -y update",

      # Install docker and docker-compose
      "dnf -y install dnf-plugins-core",
      "dnf config-manager --add-repo https://download.docker.com/linux/fedora/docker-ce.repo",
      "dnf -y install docker-ce docker-ce-cli containerd.io docker-compose-plugin",
      "systemctl start docker",
      "systemctl enable docker",
      
      # Install haveged, otherwise docker-compose may hang: https://stackoverflow.com/a/68172225/1494610
      "dnf -y install haveged",

      # Add users for services
      "useradd -m -s /bin/bash ambulance",
      "useradd -m -s /bin/bash herpetophobia",
      "useradd -m -s /bin/bash kleptophobia",
      "useradd -m -s /bin/bash meds",
      "useradd -m -s /bin/bash prosopagnosia",
      "useradd -m -s /bin/bash psycho-clinic",
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
    source = "../services/prosopagnosia/"
    destination = "/home/prosopagnosia/"
  }

  provisioner "file" {
    source = "../services/psycho-clinic/"
    destination = "/home/psycho-clinic/"
  }

  provisioner "file" {
    source = "../services/schizophasia/"
    destination = "/home/schizophasia/"
  }

  # Build and run services for the first time
  provisioner "shell" {
    inline = [
      "cd ~ambulance",
      "docker-compose up --build -d || true",
      "cd ~herpetophobia",
      "docker-compose up --build -d || true",
      "cd ~kleptophobia",
      "docker-compose up --build -d || true",
      "cd ~meds",
      "docker-compose up --build -d || true",
      "cd ~prosopagnosia",
      "docker-compose up --build -d || true",
      "cd ~psycho-clinic",
      "docker-compose up --build -d || true",
      "cd ~schizophasia",
      "docker-compose up --build -d || true",
    ]
  }

  # Fix some internal digitalocean+cloud-init scripts to be compatible with our cloud infrastructure
  provisioner "shell" {
    script = "digital_ocean_specific_setup.sh"
  }
}
