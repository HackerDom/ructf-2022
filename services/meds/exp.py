import random
import math
import sys

def compute_path(n):
	while (n > 2047):
		if n % 2 == 0:
			n = (n - 2) // 2
			print('right', n)
		else:
			n = (n - 1) // 2
			print('left', n)

compute_path(10_000_012)
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
