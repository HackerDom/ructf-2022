import sys
import htmlmin

with open('index.html', 'r') as f:
	index = f.read()
	index = htmlmin.minify(index, remove_empty_space=True).encode("utf-8")

# with open('bg.png', 'rb') as f:
# 	bg = f.read()

output = [ 
	'#pragma once', 
	'', 
	'byte pg_index[] = { %s };' % ', '.join([hex(c) for c in index]),
	# 'byte res_bg[] = { %s };' % ', '.join([hex(c) for c in bg]),
	# 'uint64 size_bg = { %d };' % len(bg)
]


for l in output:
	print(l)