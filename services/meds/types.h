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

#define MAXRECV 1024
#define MAXSEND 4096

#define max(a,b) ({ __typeof__ (a) _a = (a); __typeof__ (b) _b = (b); _a > _b ? _a : _b; })
#define min(a,b) ({ __typeof__ (a) _a = (a); __typeof__ (b) _b = (b); _a < _b ? _a : _b; })

#if DEBUG_ON
#define DEBUG(...) { fprintf(stderr, __VA_ARGS__); fflush(stderr); }
#else
#define DEBUG(...) {}
#endif