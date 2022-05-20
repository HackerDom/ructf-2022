FROM ubuntu:20.04

RUN apt update \
    && apt install -y \
        libboost-dev \
        libboost-thread-dev \
        libboost-log-dev \
        libboost-program-options-dev \
        libpq-dev

RUN useradd --no-create-home --shell /bin/false --uid 1000 --user-group schizovm

RUN mkdir /var/schizovm && chown schizovm:schizovm /var/schizovm

COPY schizovm /

RUN chmod +x /schizovm

USER schizovm

CMD ["/schizovm", "--address", "0.0.0.0", "--port", "31337"]
