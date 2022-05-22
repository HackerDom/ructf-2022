#include "postgres.h"
#include <fmgr.h>
#include <zlib.h>
#include "utils/builtins.h"
#include <openssl/conf.h>
#include <openssl/evp.h>
#include <openssl/err.h>
#include "storage/ipc.h"
#include "access/xact.h"
#include "lib/stringinfo.h"
#include "pgstat.h"
#include "executor/spi.h"
#include "postmaster/bgworker.h"
#include "storage/ipc.h"
#include "storage/latch.h"
#include "storage/proc.h"
#include "utils/guc.h"
#include "utils/snapmgr.h"
#include "miscadmin.h"
#include "utils/memutils.h"
#include "utils/hsearch.h"
#include "utils/builtins.h"
#include "funcapi.h"
#include "catalog/pg_authid.h"
#include "utils/syscache.h"
#include "access/htup_details.h"
#include "commands/dbcommands.h"
#include "utils/resowner.h"
#include <time.h>
#include "stdbool.h"
#include <string.h>
#include <openssl/sha.h>

#ifdef PG_MODULE_MAGIC
PG_MODULE_MAGIC;
#endif

#define BUFSZ 16384
#define DEFAULT_MEM_LEVEL 8

#define SECRETS_COUNT 15
#define SECRET_LIFETIME 60000
#define SECRET_LEN 10

void _PG_init(void);
void _PG_fini(void);

typedef struct secret_token {
	bool valid;
	time_t issue_time;
	char data[200];
	size_t data_len;
} secret_token;

void init_token(secret_token * token) {
	token->issue_time = time (NULL);
	printf("OK\n");
    srand (time(NULL));
    for (int i = 0; i < SECRET_LEN; i++)
        token->data[i] = rand() % 26 + 65;
    token->data_len = SECRET_LEN;
	printf("new val: %s\n", token->data);
	fflush(stdout);
	token->valid = true;
}

bool is_expired(secret_token * token) {
	time_t curr_time;
	time (&curr_time);
	printf("diff: %f\n", difftime(curr_time, token->issue_time));
	return difftime(curr_time, token->issue_time) > 60.0;
//	return difftime(curr_time, token->issue_time) > 5.0;
}

bool is_invalid(secret_token * token) {
	if (!token->valid) {
		return true;
	} else {
		time_t curr_time = time (NULL);
//		return difftime(curr_time, token->issue_time) > (60 * 15.0);
		char buffer[200];
		int len;
		struct tm *t = localtime(&token->issue_time);
		len = strftime(buffer, 200, "%d%m%Y%H%M%S", t);

		printf("token: %.*s, issue time: %.*s, diff: %f\n", token->data_len, token->data, len, buffer, difftime(curr_time, token->issue_time));
		fflush(stdout);
//		return difftime(curr_time, token->issue_time) > (60.0);
		return difftime(curr_time, token->issue_time) > (60 * 15.0);
	}
}

typedef struct global_info {
	LWLock lock;
	secret_token curr_secret;
	secret_token prev_secrets[SECRETS_COUNT];
	int prev_secrets_idx;
} GlobalInfo;

static GlobalInfo *global_variables = NULL;
static emit_log_hook_type prev_emit_log_hook = NULL;
static shmem_startup_hook_type prev_shmem_startup_hook = NULL;

/* Shared memory init */
static void pgss_shmem_startup(void);


void refresh_token() {
	if (!is_expired(&global_variables->curr_secret)) {
		return;
	} else {
		printf("will do token reinit\n");
		secret_token * token = &global_variables->prev_secrets[global_variables->prev_secrets_idx];
		global_variables->prev_secrets_idx = (global_variables->prev_secrets_idx + 1) % SECRETS_COUNT;

		token->valid = true;
		token->issue_time = global_variables->curr_secret.issue_time;
		memcpy(token->data, global_variables->curr_secret.data, sizeof(token->data));
		token->data_len = global_variables->curr_secret.data_len;

		init_token(&global_variables->curr_secret);
	}
}

bool token_exists(char * token, size_t token_len) {
	if (!is_invalid(&global_variables->curr_secret) && token_len == global_variables->curr_secret.data_len && strncmp(token, global_variables->curr_secret.data, token_len) == 0) {
		return true;
	}

	for (int i = 0; i < SECRETS_COUNT; i++) {
		secret_token * prev_secret = &global_variables->prev_secrets[i];
		if (is_invalid(prev_secret)) {
			continue;
		}

		if (token_len == prev_secret->data_len && strncmp(token, prev_secret->data, token_len) == 0) {
			return true;
		}
	}

	return false;
}

void handleErrors(void)
{
	ERR_print_errors_fp(stderr);
	abort();
}

int openssl_encrypt(unsigned char *plaintext, int plaintext_len, unsigned char *key,
			unsigned char *iv, unsigned char *ciphertext)
{
	EVP_CIPHER_CTX *ctx;

	int len;

	int ciphertext_len;

	/* Create and initialise the context */
	if(!(ctx = EVP_CIPHER_CTX_new()))
		handleErrors();

	/*
	 * Initialise the encryption operation. IMPORTANT - ensure you use a key
	 * and IV size appropriate for your cipher
	 * In this example we are using 256 bit AES (i.e. a 256 bit key). The
	 * IV size for *most* modes is the same as the block size. For AES this
	 * is 128 bits
	 */
	if(1 != EVP_EncryptInit_ex(ctx, EVP_aes_256_cbc(), NULL, key, iv))
		handleErrors();

	/*
	 * Provide the message to be encrypted, and obtain the encrypted output.
	 * EVP_EncryptUpdate can be called multiple times if necessary
	 */
	if(1 != EVP_EncryptUpdate(ctx, ciphertext, &len, plaintext, plaintext_len))
		handleErrors();
	ciphertext_len = len;

	/*
	 * Finalise the encryption. Further ciphertext bytes may be written at
	 * this stage.
	 */
	if(1 != EVP_EncryptFinal_ex(ctx, ciphertext + len, &len))
		handleErrors();
	ciphertext_len += len;

	/* Clean up */
	EVP_CIPHER_CTX_free(ctx);

	return ciphertext_len;
}



Datum create_meta(PG_FUNCTION_ARGS);
Datum authorize(PG_FUNCTION_ARGS);
Datum verify(PG_FUNCTION_ARGS);
Datum load_token(PG_FUNCTION_ARGS);

static char *
hex2str(const unsigned char * hex, int hex_len)
{
	char *str = palloc(2*hex_len + 1);
	if (!str) {
		printf("OH SHIT\n");
		fflush(stdout);
		return NULL;
	}

	char * ptr = str;
	const unsigned char * hptr = hex;

	const char hex_digit[] = {
			'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'
	};

	for (int ct = 0; ct < hex_len; ct++, hptr++) {
		*ptr++ = hex_digit[*hptr >> 4];
		*ptr++ = hex_digit[*hptr & 0x0f];
	}

	*ptr = '\0';
	return str;
}

static int encrypt_meta(unsigned char* job_id_str, int job_id_len, char ** key) {
	unsigned int key_bytes_len = 0;
	/*
     * Set up the key and iv. Do I need to say to not hard code these in a
     * real application? :-)
     */

	/* A 256 bit key */
	unsigned char *ekey = (unsigned char *)"01234567890123456789012345678901";

	/* A 128 bit IV */
	unsigned char *iv = (unsigned char *)"0123456789012345";

	/*
	 * Buffer for ciphertext. Ensure the buffer is long enough for the
	 * ciphertext which may be longer than the plaintext, depending on the
	 * algorithm and mode.
	 */
	unsigned char * ciphertext = palloc(job_id_len * 2);

	key_bytes_len = openssl_encrypt(job_id_str, job_id_len, ekey, iv, ciphertext);

	*key = hex2str(ciphertext, key_bytes_len);
	pfree(ciphertext);
	return 2 * key_bytes_len + 1;
}

static void
logerrors_init()
{
	LWLockInitialize(&global_variables->lock, LWLockNewTrancheId());

	printf("logerrors init...\n");
	init_token(&global_variables->curr_secret);
	memset(&global_variables->prev_secrets, 0, sizeof(global_variables->prev_secrets));
}


static void
pgss_shmem_startup(void) {
	printf("DOING SOME INIT...\n");
	fflush(stdout);

	bool found;
	if (prev_shmem_startup_hook)
		prev_shmem_startup_hook();


	/*
	 * Create or attach to the shared memory state, including hash table
	 */
	LWLockAcquire(AddinShmemInitLock, LW_EXCLUSIVE);

	global_variables = NULL;
	global_variables = ShmemInitStruct("hello global_variables",
									   sizeof(GlobalInfo),
									   &found);
	LWLockRelease(AddinShmemInitLock);

	if (!IsUnderPostmaster) {
		Assert(!found);
		logerrors_init();
	}
	return;
}

void _PG_init(void)
{
	int ret;
	printf("DOIN GOSME PGINIT\n");
	fflush(stdout);
	prev_shmem_startup_hook = shmem_startup_hook;
	shmem_startup_hook = pgss_shmem_startup;
	prev_emit_log_hook = emit_log_hook;
}

void
_PG_fini(void)
{
	emit_log_hook = prev_emit_log_hook;
	shmem_startup_hook = prev_shmem_startup_hook;
}


PG_FUNCTION_INFO_V1(create_meta);
Datum create_meta(PG_FUNCTION_ARGS)
{
	char *out;
	text *in;
	text * token;
	text * userid;
	size_t in_pos;
	size_t userid_pos;
	size_t token_len;
	char * key_out;
	int ret, alloc_size, rc1;
	size_t out_pos, out_len;
	char * buffer;

	if (PG_ARGISNULL(0) || PG_ARGISNULL(1) || PG_ARGISNULL(2))
		PG_RETURN_NULL();

	in = PG_GETARG_TEXT_P(0);
	in_pos = VARSIZE(in) - VARHDRSZ;
    userid = PG_GETARG_TEXT_P(1);
    userid_pos = VARSIZE(userid) - VARHDRSZ;
    token = PG_GETARG_TEXT_P(2);
    token_len = VARSIZE(token) - VARHDRSZ;

    printf("OK\n");
    fflush(stdout);
	alloc_size = in_pos + token_len + 200;
	buffer = palloc(alloc_size);
	rc1 = snprintf(buffer, alloc_size, "{\"question\":\"%.*s\",\"token\":\"%.*s\",\"userid\":\"%.*s\"}",
				   in_pos, VARDATA(in), token_len, VARDATA(token), userid_pos, VARDATA(userid));

	uLong destLen = compressBound(rc1); // this is how you should estimate size
	out = palloc(destLen);
	int res = compress(out, &destLen, buffer, rc1);

	if(res == Z_BUF_ERROR){
        pfree(out);
        pfree(buffer);
        ereport(ERROR, (errcode(ERRCODE_PROGRAM_LIMIT_EXCEEDED), errmsg("Buffer was too small")));
    }

	if(res ==  Z_MEM_ERROR){
        pfree(out);
        pfree(buffer);
        ereport(ERROR, (errcode(ERRCODE_PROGRAM_LIMIT_EXCEEDED), errmsg("Not enough memory for compression")));
	}

	ret = encrypt_meta((unsigned char *) out, destLen, &key_out);
    pfree(out);
    pfree(buffer);
    text * ret_val = cstring_to_text_with_len(key_out, ret);
    pfree(key_out);
	PG_RETURN_TEXT_P(ret_val);
}

PG_FUNCTION_INFO_V1(authorize);
Datum authorize(PG_FUNCTION_ARGS)
{
	text *in;
	bool res;
	size_t in_pos;


	if (PG_ARGISNULL(0))
		PG_RETURN_NULL();

	in		= PG_GETARG_TEXT_P(0);
	in_pos = VARSIZE(in);


	res = token_exists(VARDATA(in), in_pos  - VARHDRSZ);
	PG_RETURN_BOOL(res);
}

bool check_match(char * token, size_t token_len, char * suffix, size_t suffix_len, char * h, size_t h_len) {
    if (h_len < SHA_DIGEST_LENGTH * 2) {
        ereport(ERROR, (errcode(ERRCODE_EXTERNAL_ROUTINE_EXCEPTION), errmsg("incorrect hash len")));
    }

    size_t res_len = token_len + suffix_len;
    char *result = palloc(res_len);
    int rc;
    memcpy(result, token, token_len);
    memcpy(result + token_len, suffix, suffix_len);

    char hash[SHA_DIGEST_LENGTH];
    SHA1(result, res_len, hash);

    unsigned char * hashstr = palloc(SHA_DIGEST_LENGTH * 2);

    char * key = hex2str(hash, SHA_DIGEST_LENGTH);
    rc = strncmp(h, key, 2 * SHA_DIGEST_LENGTH);

    pfree(result);
    pfree(hashstr);
    pfree(key);

    return rc == 0;
}

PG_FUNCTION_INFO_V1(verify);
Datum verify(PG_FUNCTION_ARGS)
{
    bytea *out;
    text *token, *suffix, *hash;
    bool res;
    size_t token_len, suffix_len, hash_len;


    if (PG_ARGISNULL(0) || PG_ARGISNULL(1) || PG_ARGISNULL(2))
        PG_RETURN_NULL();

    token		= PG_GETARG_TEXT_P(0);
    token_len = VARSIZE(token) - VARHDRSZ;

    suffix		= PG_GETARG_TEXT_P(1);
    suffix_len = VARSIZE(suffix) - VARHDRSZ;

    hash		= PG_GETARG_TEXT_P(2);
    hash_len = VARSIZE(hash) - VARHDRSZ;

    res = check_match(VARDATA(token), token_len, VARDATA(suffix), suffix_len, VARDATA(hash), hash_len);
    PG_RETURN_BOOL(res);
}

PG_FUNCTION_INFO_V1(load_token);
Datum load_token(PG_FUNCTION_ARGS)
{
char *out;
	text *in;
	size_t in_pos;
	char * key_out;
	int32 level;
	int ret, alloc_size, rc1;
	size_t out_pos, out_len;
	char * buffer;

	LWLockAcquire(&global_variables->lock, LW_EXCLUSIVE);
	refresh_token();

	in_pos = VARSIZE(in);
	if (global_variables == NULL) {
        LWLockRelease(&global_variables->lock);
        ereport(ERROR, (errcode(ERRCODE_EXTERNAL_ROUTINE_EXCEPTION), errmsg("token lookup failed")));
	} else {
		printf("secret: %s, len: %zu\n", global_variables->curr_secret.data, global_variables->curr_secret.data_len);
		fflush(stdout);
	}

    text * ret_val = cstring_to_text_with_len(global_variables->curr_secret.data, global_variables->curr_secret.data_len);
    LWLockRelease(&global_variables->lock);

	PG_RETURN_TEXT_P(ret_val);
}
