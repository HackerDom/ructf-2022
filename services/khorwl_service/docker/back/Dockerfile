FROM ubuntu:21.10

RUN apt update \
    && apt install -y \
        libboost-dev \
        libboost-thread-dev \
        libboost-log-dev \
        libboost-program-options-dev \
        libpq-dev

COPY backend /

RUN chmod +x /backend

CMD ["/backend"]
