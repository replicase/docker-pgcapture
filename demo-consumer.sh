#!/bin/bash

export PG_VERSION=${PG_VERSION:-14}
export DECODE_PLUGIN=${DECODE_PLUGIN:-pgoutput}

docker-compose run --rm wait-demo-consumer-deps

docker exec pulsar sh -c "bin/pulsar-admin namespaces create public/pgcapture; bin/pulsar-admin topics create persistent://public/pgcapture/postgres"
# consumer would auto create subscription, you do not need manual create subscription
# docker exec pulsar sh -c "bin/pulsar-admin topics create-subscription --subscription postgres_cdc persistent://public/pgcapture/postgres"
docker exec pulsar sh -c "bin/pulsar-admin clusters update --broker-url http://pulsar:6605 --url http://pulsar:8080 standalone"
CSRF_TOKEN=$(curl localhost:9527/pulsar-manager/csrf-token)
curl \
    -H "X-XSRF-TOKEN: $CSRF_TOKEN" \
    -H "Cookie: XSRF-TOKEN=$CSRF_TOKEN;" \
    -H 'Content-Type: application/json' \
    -X PUT http://localhost:9527/pulsar-manager/users/superuser \
    -d '{"name": "admin", "password": "apachepulsar", "description": "test", "email": "username@test.org"}'

docker-compose up -d --build --force-recreate pg2pulsar pulsar2pg
