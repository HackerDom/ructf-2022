FROM ubuntu:20.04

COPY deploy/registry/db /app/db
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
COPY nice_ext /app/nice_ext
WORKDIR /app/nice_ext

RUN make -j16 -s && make install

