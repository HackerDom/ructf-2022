#include <sys/epoll.h>
#include <time.h>

#include "types.h"
#include "storage.h"
#include "http.h"

#define PORT 16780u

int create_listener()
{
	int sock = socket(AF_INET, SOCK_STREAM, 0);
	if (sock < 0)
	{
		perror("fuck socket\n");
		exit(1);
	}

	int enable = 1;
	if (setsockopt(sock, SOL_SOCKET, SO_REUSEADDR, &enable, sizeof(int)) < 0)
    {
		perror("fuck reuseaddr\n");
		exit(1);
	}

	struct sockaddr_in name;
	bzero(&name, sizeof(name));
	name.sin_family = AF_INET;
	name.sin_addr.s_addr = INADDR_ANY;
	name.sin_port = htons(PORT);

	if (bind(sock, (struct sockaddr *)&name, sizeof(name)))
	{
		perror("fuck bind\n");
		exit(1);
	}

	return sock;
}

void make_nonblocking(int sfd)
{
	int flags = fcntl(sfd, F_GETFL, 0);

	if (flags < 0)
	{
		perror("fuck fcntl\n");
		exit(1);
	}

	flags |= O_NONBLOCK;
	if (fcntl(sfd, F_SETFL, flags) < 0)
	{
		perror("fuck fcntl\n");
		exit(1);
	}
}

#define MAXEVENTS 64
struct epoll_event event;
struct epoll_event events[MAXEVENTS];

#define MAXFDS 16 * 1024
struct client_state {
	char recvbuf[MAXRECV];
	char sendbuf[MAXSEND];
	int transferred;
	int connected_at;
	int to_send;
};
struct client_state clients[MAXFDS];

void collect_garbage()
{
	int now = time(0);
	for (int i = 0; i < MAXFDS; i++)
	{
		if (clients[i].connected_at && now - clients[i].connected_at > 2)
		{
			close(i);
			clients[i].connected_at = 0;
			clients[i].transferred = 0;
			clients[i].to_send = 0;
			bzero(clients[i].sendbuf, sizeof(clients[i].sendbuf));
			bzero(clients[i].recvbuf, sizeof(clients[i].recvbuf));
		}
	}
}

void run()
{
	int listener = create_listener();

	make_nonblocking(listener);

	listen(listener, 100);

	int efd = epoll_create1(0);
	if (efd < 0)
	{
		perror("fuck epoll\n");
		exit(1);
	}

	event.data.fd = listener;
	event.events = EPOLLIN;
	if (epoll_ctl(efd, EPOLL_CTL_ADD, listener, &event) < 0)
	{
		perror("fuck epoll_ctl\n");
		exit(1);
	}

	while (true)
	{
		int n = epoll_wait(efd, events, MAXEVENTS, -1);

		int i;
		for (i = 0; i < n; i++)
		{
			if ((events[i].events & EPOLLERR) ||
				(events[i].events & EPOLLHUP) ||
				(!(events[i].events & EPOLLIN) && !(events[i].events & EPOLLOUT)))
			{
				printf("epoll error\n");
				close(events[i].data.fd);
				continue;
			}
			else if (listener == events[i].data.fd)
			{
				while (true)
				{
					struct sockaddr_in in_addr;
					int in_len = sizeof(in_addr);
					bzero(&in_addr, in_len);
					
					int cli = accept(listener, (struct sockaddr *)&in_addr, &in_len);
					if (cli < 0)
					{
						if ((errno == EAGAIN) || (errno == EWOULDBLOCK))
							break;
						else
						{
							printf("accept error: %d\n", errno);
							if (errno == EMFILE)
							{
								collect_garbage();
								break;
							}
							exit(1);
						}
					}
					if (cli >= MAXFDS)
					{
						printf("client fd %d is too high!", cli);
						exit(1);
					}

					make_nonblocking(cli);

					bzero(&event, sizeof(event));
					event.data.fd = cli;
					event.events = EPOLLIN | EPOLLOUT;
					if (epoll_ctl(efd, EPOLL_CTL_ADD, cli, &event) < 0)
					{
						perror("fuck epoll_ctl\n");
						exit(1);
					}

					clients[cli].connected_at = time(0);
					clients[cli].transferred = 0;
					clients[cli].to_send = 0;
				}
			}
			else
			{
				int cli = events[i].data.fd;
				if (events[i].events & EPOLLIN)
				{
					// printf("Reading from %d\n", cli);
					int bytes_read = read(cli, 
						clients[cli].recvbuf + clients[cli].transferred, 
						sizeof(clients[cli].recvbuf) - clients[cli].transferred - 1);
					if (bytes_read == 0)
					{
						close(cli);
						continue;
					}
					clients[cli].transferred += bytes_read;
					clients[cli].recvbuf[clients[cli].transferred] = 0;
					// printf("Request from %d, errno = %d:\n%s\n\n", cli, errno, clients[cli].recvbuf);
					if (process_request(clients[cli].recvbuf, clients[cli].sendbuf, &clients[cli].to_send)) {
						clients[cli].transferred = 0;
						bzero(clients[i].recvbuf, sizeof(clients[i].recvbuf));
					}
				}
				else if (clients[cli].to_send > 0)
				{
					int bytes_written = write(cli, 
						clients[cli].sendbuf + clients[cli].transferred, 
						clients[cli].to_send - clients[cli].transferred);
					clients[cli].transferred += bytes_written;
					if (clients[cli].transferred >= clients[cli].to_send || bytes_written == 0) {
						close(cli);
						bzero(clients[i].sendbuf, sizeof(clients[i].sendbuf));
					}
					//printf("Response:\n%s\n\n", clients[cli].sendbuf);
				}
			}
		}
	}


	close(listener);
}

int main()
{
	init_storage("data/storage");

	run();

	return 0;
}
