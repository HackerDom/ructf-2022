FROM ubuntu:20.04

RUN apt update \
    && apt install -y \
        libboost-dev \
        libboost-thread-dev \
        libboost-log-dev \
        libboost-program-options-dev \
        libpq-dev

RUN useradd --no-create-home --shell /bin/false --uid 1000 --user-group prosopagnosia

RUN mkdir /var/prosopagnosia && chown prosopagnosia:prosopagnosia /var/prosopagnosia

COPY prosopagnosia /

RUN chmod +x /prosopagnosia

USER prosopagnosia

CMD ["/prosopagnosia", "--address", "0.0.0.0", "--port", "31337"]
