#!/bin/bash

python3 gen_keys_dev.py
python3 gen_conf_server_dev.py
python3 gen_conf_client_dev.py

python3 gen_keys_prod.py
python3 gen_conf_server_prod.py
python3 gen_conf_client_prod.py