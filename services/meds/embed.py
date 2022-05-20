import sys
import htmlmin

with open('index.html', 'r') as f:
	index = f.read()
	index = htmlmin.minify(index, remove_empty_space=True).encode("utf-8")

with open('favicon.ico', 'rb') as f:
	favicon = f.read()

output = [ 
	'#pragma once', 
	'', 
	'byte pg_index[] = { %s };' % ', '.join([hex(c) for c in index]),
	'byte res_favicon[] = { %s };' % ', '.join([hex(c) for c in favicon]),
	'int size_favicon = { %d };' % len(favicon)
]


for l in output:
	print(l)