import requests
import time
import random
import string

reqs = 0
timeacc = 0

import sys
host = sys.argv[1] if len(sys.argv) > 1 else 'localhost'

while True:
	ts = time.time()

	url = 'http://' + host + ':4280/' + ''.join([random.choice(string.ascii_uppercase) for _ in range(32)])
	requests.post(url, data = { "secret": "FOO" })

	delta = time.time() - ts

	reqs += 1
	timeacc += delta

	if reqs % 100 == 0:
		avg = timeacc / reqs
		rate = 1 / avg
		print("avg: %.2fs, rate: %.2f/sec, total: %d" % (avg, rate, reqs))