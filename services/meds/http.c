#include "http.h"
#include "resources.h"
#include "storage.h"

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

struct input {
	char *buffer;
	int size;
	int position;
};

char consume_char(struct input* input) {
	if (input->position >= input->size)
		return 0;
	return input->buffer[input->position++];
}

enum action {
	A_ADD,
	A_END,
	A_ABORT
};

enum action next_char(char c) {
	if (c >= 'A' && c <= 'Z')
		return A_ADD;
	if (c == ' ')
		return A_END;
	return A_ABORT;
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

void render_page(char * buffer, const char * key, const char * secret)
{
	sprintf(buffer, pg_index, key, secret ? secret : "");
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

bool process_request(char *request, char *response, int *response_length)
{
	DEBUG("RECV: %s\n", request);

	char page[MAXSEND];
	bzero(page, sizeof(page));

	// printf("Verb: %s\nUrl: %s\n\n", verb.data, url.data);

	// if (!strcmp("GET", verb.data))
	// {
	// 	if (try_send_resource(url.data, response, response_length))
	// 		return true;
	// 	char * value = 0;
	// 	if (url.length == 32)
	// 		value = load_item(url.data, secret.data);
	// 	if (!value)
	// 		value = "";
	// 	render_page(page, url.data, value);
	// 	respond(response, response_length, 200, page, "text/html");
	// 	return true;
	// }
	
	// if (!strcmp("POST", verb.data))
	// {
	// 	if (!read_headers(&request, &cl))
	// 		return false;
	// 	if (cl.length != 0)
	// 	{
	// 		uint64_t content_length = atoll(cl.data);
	// 		if (strlen(request) < content_length)
	// 			return false;
	// 		read_body(request, content_length, &secret);
	// 	}
	// 	if (secret.length == 0)
	// 	{
	// 		respond(response, response_length, 400, 0, "text/html");
	// 		return true;
	// 	}

	// 	char value[64];
	// 	char key[64];
	// 	bzero(key, sizeof(key));

	// 	int i = 0;
	// 	for (i = 0; i < 64; i++)
	// 	{	
	// 		gen_key(key, i);

	// 		if (load_item(key, value))
	// 			continue;

	// 		store_item(key, secret.data);
	// 		break;
	// 	}
	// 	// fprintf(stderr, "Sending key: %s for secret: %s, i = %d\n", key, secret.data, i);

	// 	render_page(page, key, secret.data);
	// 	redirect(response, response_length, key);
	// 	return true;
	// }

	respond(response, response_length, 400, 0, "text/html");
	return true;
}