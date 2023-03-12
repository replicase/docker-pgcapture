#!/bin/bash

docker-compose run --rm wait-demo-scheduler-deps
docker exec pulsar sh -c "bin/pulsar-admin namespaces create public/pgcapture; bin/pulsar-admin topics create persistent://public/pgcapture/postgres"

# use configure service to poke agent to enable pulsar2pg
# since gateway use StaticAgentPulsarResolver, and Dumper function implemented by AgentSourceDumper
docker-compose up -d --build --force-recreate pg2pulsar configure
