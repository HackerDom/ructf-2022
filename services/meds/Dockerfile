FROM ubuntu:18.04

RUN set -eux; \
	groupadd -r meds --gid=999; \
	useradd -r -g meds --uid=999 --home-dir=/app --shell=/bin/bash meds;

COPY bin/meds /app/
COPY bin/run.sh /app/

WORKDIR /app

CMD ["/app/run.sh"]
