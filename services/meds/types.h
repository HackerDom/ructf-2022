#pragma once

#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <netinet/in.h>
#include <fcntl.h>

typedef int int32;
typedef long long int64;
typedef unsigned int uint32;
typedef unsigned long long uint64;
typedef unsigned char byte;
typedef int64 intptr;
typedef int64 bool;
typedef unsigned short uint16;
typedef intptr string;

#define false 0
#define true 1

#define MAXRECV 1024
#define MAXSEND 4096


#define max(a,b) ({ __typeof__ (a) _a = (a); __typeof__ (b) _b = (b); _a > _b ? _a : _b; })