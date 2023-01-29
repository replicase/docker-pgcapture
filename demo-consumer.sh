#!/bin/bash

docker-compose up -d --build --force-recreate pulsar
sleep 10
docker exec pulsar sh -c "bin/pulsar-admin namespaces create public/pgcapture; bin/pulsar-admin topics create persistent://public/pgcapture/postgres"

docker-compose up -d --build --force-recreate postgres_source postgres_sink
sleep 10

docker-compose up -d --build --force-recreate pg2pulsar pulsar2pg gateway
sleep 5
