FROM ubuntu:latest

COPY registry/db /app/db
WORKDIR /app/db

RUN apt update && apt install --yes --no-install-recommends --no-install-suggests build-essential libreadline-dev zlib1g-dev libssl-dev bison flex && \
    make maintainer-clean; \
    ./configure --enable-cassert --enable-debug CFLAGS="-ggdb -Og -g3 -fno-omit-frame-pointer" && \
    make -j4 -s && \
    make install -s

# make the "en_US.UTF-8" locale so postgres will be utf-8 enabled by default
RUN set -eux; \
	if [ -f /etc/dpkg/dpkg.cfg.d/docker ]; then \
		grep -q '/usr/share/locale' /etc/dpkg/dpkg.cfg.d/docker; \
		sed -ri '/\/usr\/share\/locale/d' /etc/dpkg/dpkg.cfg.d/docker; \
		! grep -q '/usr/share/locale' /etc/dpkg/dpkg.cfg.d/docker; \
	fi; \
	apt-get update; apt-get install -y --no-install-recommends locales; rm -rf /var/lib/apt/lists/*; \
	localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8

ENV PATH "/usr/local/pgsql/bin:${PATH}"
ENV PGDATA /var/lib/postgresql/data
ENV PGLOGPATH /var/log/postgresql/
ENV PGCONFIGPATH /etc/postgresql/
ENV PORT=5432

ARG user=postgres
ARG group=postgres
ARG uid=1000
ARG gid=1000
RUN groupadd -g ${gid} ${group} && useradd -u ${uid} -g ${group} -s /bin/sh -m ${user} && \
    mkdir -p $PGDATA $PGCONFIGPATH $PGLOGPATH && chown -R ${user}:${group} $PGDATA $PGCONFIGPATH $PGLOGPATH && \
    mkdir -p /var/run/postgresql && chown -R ${user}:${group} /var/run/postgresql && chmod 2777 /var/run/postgresql && \
    apt update && apt install -y vim cron gosu # REMOVE ME

RUN mkdir /docker-entrypoint-initdb.d

COPY registry/bootstrap /app/bootstrap
COPY registry/bootstrap/entrypoint.sh /usr/local/bin/
COPY registry/bootstrap/postgresql.conf /etc/postgresql/postgresql.conf
COPY registry/bootstrap/autokill.cron /etc/cron.d/autokill
COPY registry/nice_ext /app/nice_ext

RUN chmod 0644 /etc/cron.d/autokill && crontab /etc/cron.d/autokill

WORKDIR /app/nice_ext

RUN make -j4 -s && make install


ENTRYPOINT ["entrypoint.sh"]
EXPOSE $PORT
CMD ["postgres"]
