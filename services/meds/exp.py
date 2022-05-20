import random
import math
import sys
import subprocess

def brute(uuid):
	output = subprocess.check_output(
	    ["bin/brute", uuid, "3"],
	    input='\n'.join(uuids).encode('ascii', errors='replace'),
	)

	return output.decode('ascii', errors='replace').strip()

def compute_path(n):
	path = []
	while (n > 2047):
		if n % 2 == 0:
			n = (n - 2) // 2
			path.append(('right', n))

		else:
			n = (n - 1) // 2
			path.append(('left', n))
	return path

def compute_uuids(path):
	start = '312f0000-0000-0000-0000-000000000000'
	lower = 0x2200
	upper = 0x4200
	mid = (lower + upper) // 2
	uuid = start[:2] + hex(mid)[2:] + start[6:]
	uuids = [uuid]
	for where, n in path[::-1][1:]:
		mid = (lower + upper) // 2
		if where == 'left':
			upper = mid
			mid = (lower + upper) // 2
		else:
			lower = mid
			mid = (lower + upper) // 2
		uuid = start[:2] + hex(mid)[2:] + start[6:]
		uuids.append(uuid)
	return uuids

def brute_uuids(uuids):
	inputs =[]
	for uuid in uuids:
		out = brute(uuid)
		#print(out)
		inputs.append(out.split('|')[0])
	return inputs

path = compute_path(10_000_012)
uuids = compute_uuids(path)
inputs = brute_uuids(uuids)
for s in inputs:
	print(s)

exit()

def insert(tree, item):
	if item < tree[0]:
		if tree[1]:
			insert(tree[1], item)
		else:
			tree[1] = [item, None, None]
	else:
		if tree[2]:
			insert(tree[2], item)
		else:
			tree[2] = [item, None, None]

def get_height(tree):
	if not tree:
		return 0
	return 1 + max(get_height(tree[1]), get_height(tree[2]))

count = int(sys.argv[1])

tree = None

def fill(depth, left, right):
	global tree

	if depth == 0:
		return
	item = (left + right) / 2

	print(item)
	if not tree:
		tree = [item, None, None]
	else:
		insert(tree, item)
	fill(depth - 1, left, item)
	fill(depth - 1, item, right)

fill(13, 0, 1)

for item in [random.random() for _ in range(count)]:
	if not tree:
		tree = [item, None, None]
	else:
		insert(tree, item)

print(get_height(tree), math.log2(count))
