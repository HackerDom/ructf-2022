import requests
import time
import random
import string
import uuid

reqs = 0
timeacc = 0

import sys
host = sys.argv[1] if len(sys.argv) > 1 else 'localhost'

while True:
	ts = time.time()

    key = str(uuid.uuid4())
	url = "http://" + hostname + ":16780/" + key;
    response = requests.post(url, data = b"diag=FOO", allow_redirects = False)

	delta = time.time() - ts

	reqs += 1
	timeacc += delta

	if reqs % 100 == 0:
		avg = timeacc / reqs
		rate = 1 / avg
		print("avg: %.2fs, rate: %.2f/sec, total: %d" % (avg, rate, reqs))