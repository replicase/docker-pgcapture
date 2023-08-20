FROM ubuntu:18.04 AS extension_builder

RUN apt update && \
    apt install -y wget gnupg2 && \
    wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add - && \
    echo "deb http://apt.postgresql.org/pub/repos/apt bionic-pgdg main" > /etc/apt/sources.list.d/pgdg.list && \
    apt update && \
    apt install -y postgresql-server-dev-11 build-essential libssl-dev libkrb5-dev zlib1g-dev libedit-dev libxml2-dev libxslt-dev libselinux-dev libpam-dev

RUN wget https://github.com/rueian/pglogical/archive/REL2_3_4_no_filter.tar.gz && \
    wget https://github.com/rueian/pgcapture/archive/v0.0.22.tar.gz && \
    tar -zxvf REL2_3_4_no_filter.tar.gz && \
    tar -zxvf v0.0.22.tar.gz

RUN cd /pglogical-REL2_3_4_no_filter && make && \
    cd /pgcapture-0.0.22/hack/postgres/extension && make

RUN cd /pglogical-REL2_3_4_no_filter && make install && \
    cd /pgcapture-0.0.22/hack/postgres/extension && make install

FROM postgres:11

COPY --from=extension_builder /usr/lib/postgresql/11/lib /usr/lib/postgresql/11/lib

COPY --from=extension_builder /usr/share/postgresql/11/extension /usr/share/postgresql/11/extension
