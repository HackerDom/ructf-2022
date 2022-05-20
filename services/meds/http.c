#include "http.h"
#include "resources.h"
#include "storage.h"
#include "diag.h"

struct strbuf
{
	size_t length;
	size_t limit;
	char data[512];
};

void init_strbuf(struct strbuf * buf, size_t limit)
{
	bzero(buf, sizeof(*buf));
	if (limit >= sizeof(buf->data))
		limit = sizeof(buf->data) - 1;
	buf->limit = limit;
}

bool try_add_char(struct strbuf * buf, char c)
{
	if (buf->length < buf->limit)
	{
		buf->data[buf->length++] = c;
		return true;
	}
	return false;
}

struct request_info {
	struct strbuf verb;
	struct strbuf url;
	struct strbuf body;
	int content_length;
};

struct input {
	char *buffer;
	int size;
	int position;
};

char peek_char(struct input* input) {
	if (input->position >= input->size)
		return 0;
	return input->buffer[input->position];
}

char consume_char(struct input* input) {
	if (input->position >= input->size)
		return 0;
	return input->buffer[input->position++];
}

enum action {
	A_SKIP,
	A_ADD,
	A_END,
	A_ABORT
};

enum state {
	S_NONE,
	S_WS,
	S_VERB,
	S_URL,
	S_PROTO,
	S_HEADER_KEY,
	S_HEADER_VALUE,
	S_BODY,
	S_END
};

#define RS(x) case x: return #x;
char * render_state(enum state state) {
	switch (state) {
		RS(S_NONE)
		RS(S_WS)
		RS(S_VERB)
		RS(S_URL)
		RS(S_PROTO)
		RS(S_HEADER_KEY)
		RS(S_HEADER_VALUE)
		RS(S_BODY)
		RS(S_END)
	}
	return "<unk>";
}

struct parsing_info {
	struct strbuf* buf;
	int newline_count;
	bool is_cl;
	enum state next_state;
};

bool is_ws(char c) {
	return c == ' ' || c == '\r' || c == '\n';
}

enum action next_char(char c, bool init, enum state* state, struct parsing_info* parsing_info, struct request_info* request_info) {
	switch (*state) {
		case S_WS:
			if (init)
				parsing_info->newline_count = 0;
			if (c == '\n')
				parsing_info->newline_count++;
			if (is_ws(c))
				return A_SKIP;

			if (parsing_info->next_state == S_NONE)
				return A_ABORT;

			*state = parsing_info->next_state;
			return A_END;

		case S_VERB:
			if (init) {
				init_strbuf(&request_info->verb, 8);
				parsing_info->buf = &request_info->verb;
				parsing_info->next_state = S_URL;
			}
			if (c >= 'A' && c <= 'Z')
				return A_ADD;
			if (c == ' ') {
				*state = S_WS;
				return A_END;
			}
			return A_ABORT;

		case S_URL:
			if (init) {
				init_strbuf(&request_info->url, 64);
				parsing_info->buf = &request_info->url;
				parsing_info->next_state = S_PROTO;
			}
			if (c == '/' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '-' || c >= '0' && c <= '9' || c == '.')
				return A_ADD;
			if (c == ' ') {
				*state = S_WS;
				return A_END;
			}
			return A_ABORT;

		case S_PROTO:
			if (init) {
				if (parsing_info->newline_count == 1) {
					*state = S_HEADER_KEY;
					return A_END;
				}
				if (parsing_info->newline_count == 2) {
					*state = S_BODY;
					return A_END;
				}
				parsing_info->next_state = S_PROTO;
			}
			if (c == '/' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '.')
				return A_SKIP;
			if (is_ws(c)) {
				*state = S_WS;
				return A_END;
			}
			return A_ABORT;

		case S_HEADER_KEY:
			if (init) {
				parsing_info->next_state = S_HEADER_VALUE;
			}
			if (c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '-' || c == ':')
				return A_ADD;
			if (c == ' ') {
				parsing_info->is_cl = !strcasecmp(parsing_info->buf->data, "Content-Length:");
				*state = S_WS;
				return A_END;
			}
			return A_ABORT;

		case S_HEADER_VALUE:
			if (init) {
				if (parsing_info->newline_count == 1) {
					*state = S_HEADER_KEY;
					return A_END;
				}
				if (parsing_info->newline_count == 2) {
					*state = S_BODY;
					return A_END;
				}
				parsing_info->next_state = S_HEADER_VALUE;
			}
				
			if (parsing_info->is_cl) {
				if (c >= '0' && c <= '9')
					return A_ADD;
			}
			if (is_ws(c)) {
				if (parsing_info->is_cl)
					request_info->content_length = atoi(parsing_info->buf->data);
				*state = S_WS;
				return A_END;
			}
			return A_SKIP;

		case S_BODY:
			if (init) {
				init_strbuf(&request_info->body, 512);
				parsing_info->buf = &request_info->body;
			}
			if (parsing_info->buf->length >= request_info->content_length) {
				*state = S_END;
				return A_END;
			}
			if (!c)
				return A_ABORT;
			return A_ADD;
	}
}

enum parse_result {
	PR_SUCCESS,
	PR_INVALID,
	PR_INCOMPLETE
};

enum parse_result parse_request(struct input* input, struct request_info* info) {
	enum state state = S_VERB;
	
	struct strbuf default_buf;
	init_strbuf(&default_buf, 512);

	struct parsing_info parsing_info;
	parsing_info.buf = &default_buf;
	parsing_info.newline_count = -1;
	parsing_info.is_cl = false;
	parsing_info.next_state = S_NONE;
	
	bool init = true;

	while (state != S_END) {
		char c = peek_char(input);
		// DEBUG("!! consuming c = %c w/ state = %s, init = %d\n", c, render_state(state), init);
		enum action action = next_char(c, init, &state, &parsing_info, info);
		if (init)
			init = false;
		switch (action) {
			case A_ADD:
				if (!try_add_char(parsing_info.buf, c))
					return PR_INVALID;
			case A_SKIP:
				consume_char(input);
				break;
			case A_END:
				init = true;
				init_strbuf(&default_buf, 512);
				parsing_info.buf = &default_buf;
				break;
			default:
				return c ? PR_INVALID : PR_INCOMPLETE;
		}
	}

	return PR_SUCCESS;
}


void respond_bytes(char *response, int *response_length, int code, const char *text, int length, const char *content_type)
{
	char * code_msg;
	switch (code)
	{
		case 200:
			code_msg = "HTTP/1.1 200 OK";
			break;
		case 400:
			code_msg = "HTTP/1.1 400 Bad Request";
			break;
		case 404:
			code_msg = "HTTP/1.1 404 Not Found";
			break;
		default:
			code_msg = "HTTP/1.1 500 Internal Server Error";
			break;
	}

	int pos = sprintf(response, 
		"%s\r\n"
		"Server: meds\r\n"
		"Content-Length: %d\r\n"
		"Content-Type: %s\r\n"
		"Connection: close\r\n\r\n", 
		code_msg, length, content_type);
	if (text)
		memcpy(response + pos, text, length);
	*response_length = pos + length;
}

void respond(char *response, int *response_length, int code, const char *text, const char *content_type)
{
	respond_bytes(response, response_length, code, text, text ? strlen(text) : 0, content_type);
}

void redirect(char *response, int *response_length, const char *location)
{
	*response_length = sprintf(response, 
		"HTTP/1.1 303 See Other\r\n"
		"Location: /%s\r\n"
		"Connection: close\r\n\r\n", location);
}

void render_page(char * buffer, const char* value, bool has_meds)
{
	char diag[MAXDIAG + 1];
	char meds[MAXMEDS];
	bzero(diag, sizeof(diag));
	bzero(meds, sizeof(meds));

	if (has_meds) {
		char *split = strrchr(value, '|');
		if (split) {
			strncpy(diag, value, split - value);
			strcpy(meds, split + 1);
		} else {
			strcpy(diag, value);
		}
	} else {
		strcpy(diag, value);
	}

	sprintf(buffer, pg_index, diag, has_meds ? "" : "hidden", has_meds ? meds : "");
}

bool try_send_resource(const char *name, char *response, int *response_length)
{
	// if (!strcmp("A", name))
	// {
	// 	respond_bytes(response, response_length, 200, res_bg, size_bg, "image/png");
	// 	return true;
	// }
	return false;
}

bool extract_key(struct strbuf* url, uuid_t key) {
	if (url->length != 37)
		return false;
	return uuid_parse(url->data + 1, key) == 0;
}

bool extract_value(struct strbuf* body, value_t value) {
	const int key_length = 5;
	if (body->length < key_length + 1 || body->length - key_length > MAXDIAG)
		return false;
	if (strncmp("diag=", body->data, key_length))
		return false;
	bzero(value, sizeof(value_t));
	memcpy(value, body->data + key_length, body->length - key_length);
	return strlen(value) > 0;
}

bool process_request(char *request, char *response, int *response_length)
{
	DEBUG("RECV: %s\n", request);

	struct input input;
	input.buffer = request;
	input.size = MAXRECV;
	input.position = 0;

	struct request_info info;
	bzero(&info, sizeof(info));

	enum parse_result result = parse_request(&input, &info); 

	DEBUG("Parsed: %d, verb = %s, url = %s, cl = %d, body = %s\n", result, info.verb.data, info.url.data, info.content_length, info.body.data);

	if (result == PR_INCOMPLETE)
		return false;

	char page[MAXSEND];
	bzero(page, sizeof(page));

	if (result == PR_SUCCESS) {

		if (!strcmp("GET", info.verb.data))
		{
			if (try_send_resource(info.url.data, response, response_length))
				return true;

			if (!strcmp("/", info.url.data)) {
				render_page(page, "Enter your diagnosis here.", false);
				respond(response, response_length, 200, page, "text/html");
				return true;
			} else {
				uuid_t key;
				if (extract_key(&info.url, key)) {
					value_t value;
					if (load_item(key, value)) {
						render_page(page, value, true);
						respond(response, response_length, 200, page, "text/html");
						return true;
					}
				}
			}

			respond(response, response_length, 404, 0, "text/html");
			return true;
		}
	}
	
	if (!strcmp("POST", info.verb.data))
	{
		value_t value;
		if (extract_value(&info.body, value)) {
			char meds[MAXMEDS];
			prescribe(value, meds);
			strcat(value, "|");
			strcat(value, meds);

			uuid_t key;
			if (!extract_key(&info.url, key)) {
				respond(response, response_length, 400, 0, "text/html");
				return true;
			}
			DEBUG("!! extracted key: %s\n", render_uuid(key));

			if (store_item(key, value)) {
				render_page(page, value, true);
				
				char key_str[37];
				bzero(key_str, sizeof(uuid_t));
				uuid_unparse_lower(key, key_str);
				redirect(response, response_length, key_str);
				return true;
			}

			respond(response, response_length, 500, 0, "text/html");
			return true;
		}
	}

	respond(response, response_length, 400, 0, "text/html");
	return true;
}