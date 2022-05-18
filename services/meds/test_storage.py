import uuid
import subprocess
import sys
import math

count = int(sys.argv[1])

uuids = [str(uuid.uuid4()) for _ in range(count)]

output = subprocess.check_output(
    ["bin/test_storage"],
    input='\n'.join(uuids).encode('ascii'),
).decode('ascii')

print(output.strip())

with open("tree/tree.html", "r") as f:
	with open("bin/tree.html", "w") as fout:
		fout.write(f.read().replace("JSON", output))