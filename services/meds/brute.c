#include "types.h"
#include "farmhash_adapter.h"
#include "diag.h"

char alpha[] = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&\'()*+,-./:;<=>?@[\\]^_`{|}~ ";
int alpha_length;

uuid_t target;
int target_length;

void check(char* comb) {
	value_t value;
	bzero(value, sizeof(value));
	strcpy(value, comb);

	char meds[MAXMEDS];
	prescribe(value, meds);
	strcat(value, "|");
	strcat(value, meds);

	uuid_t key;
	fingerprint128(value, key);
	if (memcmp(key, target, target_length) == 0)	{
		char buf[37];
		uuid_unparse_lower(key, buf);
		printf("%s|%s -> %s (%d)\n", comb, meds, buf, target_length);
		exit(0);
	}
}

void brute(char* comb, int pos) {
	if (pos < 0) {
		check(comb);
		return;
	}
	for (int i = 0; i < alpha_length; i++) {
		comb[pos] = alpha[i];
		brute(comb, pos - 1);
	}
}

int main(int argc, char** argv) {
	alpha_length = strlen(alpha);
	char comb[32];
	bzero(comb, sizeof(comb));

	uuid_parse(argv[1], target);
	target_length = atoi(argv[2]);

	for (int length = 0; length < sizeof(comb); length++) {
		brute(comb, length);
	}

	return 0;
}