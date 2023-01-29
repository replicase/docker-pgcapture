#!/bin/bash

docker-compose up -d --build --force-recreate pulsar
sleep 10
docker exec pulsar sh -c "bin/pulsar-admin namespaces create public/pgcapture; bin/pulsar-admin topics create persistent://public/pgcapture/postgres"

docker-compose up -d --build --force-recreate postgres_source
docker-compose up -d --build --force-recreate postgres_sink
sleep 10

# use configure service to poke agent to enable pulsar2pg
# since gateway use StaticAgentPulsarResolver, and Dumper function implemented by AgentSourceDumper
docker-compose up -d --build --force-recreate pg2pulsar agent controller gateway configure
sleep 5
