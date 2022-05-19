#!/usr/bin/env python3

import requests
import traceback
import random
import string
import time

from bs4 import BeautifulSoup
from gornilo import CheckRequest, Verdict, Checker, PutRequest, GetRequest
from gornilo.models.verdict.verdict_codes import *

from helpers import get_diagnosis, get_prescription

checker = Checker()


@checker.define_check
def check_service(request: CheckRequest) -> Verdict:
    rand_name = ''.join(random.choice(string.ascii_uppercase) for _ in range(31)) + '='
    diag = get_diagnosis(rand_name)
    meds = get_prescription(diag)

    url = "http://" + request.hostname + ":16780/"
    try:
        response = requests.post(url, data = "diag=" + diag, allow_redirects = True)
        soup = BeautifulSoup(response.text, features="html.parser")
        actual_meds = soup.find(id="meds").text

        if actual_meds != meds:
            print("Prescription mismatch! Expected %s, got %s. Diagnosis was: '''%s'''" % (meds, actual_meds, diag))
            return Verdict.MUMBLE("Prescription mismatch!")

        return Verdict.OK()
    except:
        traceback.print_exc()
        return Verdict.MUMBLE("Couldn't get a meaningful response!")


@checker.define_put(vuln_num=1, vuln_rate=1)
def put_flag(request: PutRequest) -> Verdict:
    diag = get_diagnosis(request.flag)

    url = "http://" + request.hostname + ":16780/"
    try:
        for i in range(3):
            response = requests.post(url, data = "diag=" + diag, allow_redirects = False)
            key = response.headers['Location'][1:]
            if len(key) == 0:
                print("Couldn't get flag id. Response was: ", vars(response))
                time.sleep(i + 1)
                continue
            print("Saved flag " + request.flag)
            return Verdict.OK(key)
        return Verdict.MUMBLE("Couldn't put flag!")
    except:
        traceback.print_exc()
        return Verdict.MUMBLE("Couldn't get a meaningful response!")


@checker.define_get(vuln_num=1)
def get_flag(request: GetRequest) -> Verdict:
    url = "http://" + request.hostname + ":16780/" + request.flag_id.strip()
    try:
        response = requests.get(url)
        soup = BeautifulSoup(response.text, features="html.parser")
        diag = soup.find(id="diag")
        if diag and request.flag in diag.text:
            return Verdict.OK()
        print("Diagnosis doesn't contain a correct flag. '''%s''' doesn't contain %s. Flag id = %s." % (diag, request.flag, request.flag_id.strip()))
        print("Response: ", vars(response))
        return Verdict.CORRUPT("Flag is missing!")
    except:
        traceback.print_exc()
        return Verdict.MUMBLE("Couldn't get a meaningful response!")



if __name__ == '__main__':
    checker.run()
