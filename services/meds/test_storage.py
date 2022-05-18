import uuid
import subprocess
import sys
import math

count = int(sys.argv[1])
command = sys.argv[2]

uuids = [str(uuid.uuid4()) for _ in range(count)]

def call_storage(arg):
	try:
		output = subprocess.check_output(
		    ["bin/test_storage", arg],
		    input='\n'.join(uuids).encode('ascii'),
		)
		err = None
	except Exception as e:
		output = e.output
		err = e

	return (output.decode('ascii').strip(), err)

output, err = call_storage(command)

print(output)

if err:
	raise err


if command == "dump":
	with open("tree/tree.html", "r") as f:
		with open("bin/tree.html", "w") as fout:
			fout.write(f.read().replace("JSON", output))