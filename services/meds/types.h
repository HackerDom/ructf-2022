#pragma once

#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <strings.h>
#include <unistd.h>
#include <netinet/in.h>
#include <fcntl.h>
#include <uuid/uuid.h>

#define true 1
#define false 0

typedef int bool;
typedef unsigned char byte;
typedef char value_t[256];

#define MAXRECV 1024
#define MAXSEND 8192

#define MAXMEDS 32
#define MAXDIAG 200

#define max(a,b) ({ __typeof__ (a) _a = (a); __typeof__ (b) _b = (b); _a > _b ? _a : _b; })
#define min(a,b) ({ __typeof__ (a) _a = (a); __typeof__ (b) _b = (b); _a < _b ? _a : _b; })

#if DEBUG_ON
#define DEBUG(...) { fprintf(stderr, __VA_ARGS__); fflush(stderr); }
#else
#define DEBUG(...) {}
#endif

#define reverse_bytes(a) (((a & 0xff) << 24) | ((a & 0xff00) << 8) | ((a & 0xff0000) >> 8) | ((a & 0xff000000) >> 24))