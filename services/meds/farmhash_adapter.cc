#include "farmhash.h"

extern "C" {
	void fingerprint128(const char* value, char hash[16]) {
		util::uint128_t result = util::Fingerprint128(value, strlen(value));
		memcpy(hash, &result, 16);
	}
}