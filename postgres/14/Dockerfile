FROM ubuntu:18.04 AS extension_builder

RUN apt update && \
    apt install -y wget gnupg2 && \
    wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add - && \
    echo "deb http://apt.postgresql.org/pub/repos/apt bionic-pgdg main" > /etc/apt/sources.list.d/pgdg.list && \
    apt update && \
    apt install -y postgresql-server-dev-14 build-essential libssl-dev libkrb5-dev zlib1g-dev libedit-dev libxml2-dev libxslt-dev libselinux-dev libpam-dev liblz4-dev patch

RUN wget https://github.com/2ndQuadrant/pglogical/archive/REL2_4_2.tar.gz && \
    wget https://github.com/replicase/pgcapture/archive/v0.0.56.tar.gz && \
    tar -zxvf REL2_4_2.tar.gz && \
    tar -zxvf v0.0.56.tar.gz

COPY 14/pglogical/pglogical.patch /pglogical-REL2_4_2

RUN cd /pglogical-REL2_4_2 && \
    patch -p1 < pglogical.patch && \
    make && \
    cd /pgcapture-0.0.56/hack/postgres/extension && make

RUN cd /pglogical-REL2_4_2 && make install && \
    cd /pgcapture-0.0.56/hack/postgres/extension && make install

FROM postgres:14

COPY --from=extension_builder /usr/lib/postgresql/14/lib /usr/lib/postgresql/14/lib

COPY --from=extension_builder /usr/share/postgresql/14/extension /usr/share/postgresql/14/extension
