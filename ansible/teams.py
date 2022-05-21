#!/usr/bin/python3

import pycountry
import requests
import os

user = os.getenv('SITE_USER', 'checker')
password = os.getenv('SITE_PASSWORD')
r = requests.get('https://ructf.org/checker_json/', auth=(user, password))
teams = r.json()

for team in teams:
    a = 60 + int(team['id'] / 256)
    b = team['id'] % 256

    network = "10.{a}.{b}.0/24".format(a=a, b=b)
    host    = "10.{a}.{b}.2".format(a=a, b=b)

    name = team['name'].replace("'", "\\'") + " " + \
        pycountry.countries.search_fuzzy(team['country'])[0].flag

    if not team['logo']:
      logo = "https://ructfe.org/ctf-static/dummy.69777ab438ae.png"
    else:
      logo = "https://ructf.org{}".format(team['logo'])

    print("  {{name => '{name}', network => '{network}', host => '{host}', token => '{token}', country => '{country}', logo => '{logo}'}},".format(name=name, network=network, host=host, token=team['checker_token'], country=team['country'], logo=logo))
