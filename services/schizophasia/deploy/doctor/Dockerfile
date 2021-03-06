FROM alpine:latest

ARG BUILD_DATE
ARG VCS_REF
ARG VERSION

ENV LANG en_US.utf8
RUN set -eux; \
	addgroup -g 70 -S postgres; \
	adduser -u 70 -S -D -G postgres -H -h /var/lib/postgresql -s /bin/sh postgres; \
	mkdir -p /var/lib/postgresql; \
	chown -R postgres:postgres /var/lib/postgresql


COPY registry/db /usr/src/postgresql
WORKDIR /app/db

# echo "$PG_SHA256 *postgresql.tar.bz2" | sha256sum -c -; \
RUN set -eux; \
    apk add --no-cache --virtual .build-deps \
		llvm10-dev \
		bison \
		coreutils \
		dpkg-dev dpkg \
		flex \
		gcc \
		libc-dev \
		libedit-dev \
		libxml2-dev \
		libxslt-dev \
		linux-headers \
	    clang g++ \
		make \
		util-linux-dev \
		zlib-dev \
        icu-dev \
    ;\
    \
    mkdir -p /docker-entrypoint-initdb.d ; \
    cd /usr/src/postgresql ; \
	gnuArch="$(dpkg-architecture --query DEB_BUILD_GNU_TYPE)"; \
	(make maintainer-clean || true) && \
	./configure \
		--build="$gnuArch" \
		--enable-integer-datetimes \
		--enable-thread-safety \
		--disable-rpath \
		--with-uuid=e2fs \
		--with-gnu-ld \
		--with-system-tzdata=/usr/share/zoneinfo \
		--prefix=/usr/local \
		--with-includes=/usr/local/include \
		--with-libraries=/usr/local/lib \
		;\
	make -j8 -s;\
	make install -j8 -s \
	; \
	runDeps="$( \
		scanelf --needed --nobanner --format '%n#p' --recursive /usr/local \
			| tr ',' '\n' \
			| sort -u \
			| awk 'system("[ -e /usr/local/lib/" $1 " ]") == 0 { next } { print "so:" $1 }' \
	)"; \
	apk add --no-cache --virtual .postgresql-rundeps \
		$runDeps \
		bash \
		tzdata \
	;\
    apk del --no-network .build-deps; \
	cd /; \
	rm -rf \
		/usr/src/postgresql \
		/usr/local/share/doc \
		/usr/local/share/man \
	; \
	\
	postgres --version

ENV GIN_MODE=release
ENV PORT=18181

COPY --from=golang:1.18-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH /go

ENV CGO_ENABLED=1
ENV GOOS=linux
# Run the two commands below to install git and dependencies for the project.
RUN apk update && apk add --no-cache git libc6-compat build-base


COPY doctor /go/src/ructf.org/doctor
WORKDIR /go/src/ructf.org/doctor

ARG USER=appuser
ARG GROUP=appgroup
ARG UID=1337
ARG GID=1337
RUN addgroup --gid ${GID} ${GROUP} && adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --ingroup ${GROUP} \
    --no-create-home \
    --uid ${UID} \
    ${USER} && \
    chown -R ${USER}:${GROUP} /go

USER ${UID}:${GID}

RUN go get ./... && go build -v .

EXPOSE $PORT

ENTRYPOINT ["./doctor-service"]
