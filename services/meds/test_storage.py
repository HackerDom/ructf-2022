import uuid
import subprocess
import sys
import math

count = int(sys.argv[1])
command = sys.argv[2]

def call_storage(arg, count):

	uuids = [str(uuid.uuid4()) for _ in range(count)]

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
			output, err = call_storage('height', count)
			height = int(output)
			hmax = max(hmax, height)
			havg += height
		havg /= runs
		print('%d: max: %d, avg: %.2f' % (count, hmax, havg))


output, err = call_storage(command, count)

print(output)

if err:
	raise err

if command == "dump":
	with open("tree/tree.html", "r") as f:
		with open("bin/tree.html", "w") as fout:
			fout.write(f.read().replace("JSON", output))