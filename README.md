# docker-pgcapture

Demo [pgcapture](https://github.com/replicase/pgcapture) amazing library by docker !

## Demo cdc consumer
1. ```bash
   # build pgcapture image if you do not have pgcapture image in your local environment.
   (cd pgcapture && ./dockerbuild.sh) && \
   # default postgres version is 14 and decode plugin is pgoutput
   # you can specify postgres version and decode plugin by setting environment variables
   # example: POSTGRES_VERSION=11 DECODE_PLUGIN=pglogical_output ./demo-consumer.sh
   ./demo-consumer.sh && \
   go run consumer/main.go
   ```
2. connect localhost:5432 postgres and create users table and insert data
   ```sql
   create table users (id int primary key, name text not null, uid uuid not null, info jsonb not null, addresses text[] not null);
   insert into users(id, name, uid, info) values (1, 'kenny', 'bc03d615-8afb-452d-b0cc-340087def732', '{"myAge": 18}', '{"taipei", "hsinchu"}'); 
   ```

## Demo cdc consumer and scheduler dump
1. ```bash
   # build pgcapture image if you do not have pgcapture image in your local environment.
   (cd pgcapture && ./dockerbuild.sh) && \
   # default postgres version is 14 and decode plugin is pgoutput
   # you can specify postgres version and decode plugin by setting environment variables
   # example: POSTGRES_VERSION=11 DECODE_PLUGIN=pglogical_output ./demo-consumer.sh
   ./demo-scheduler.sh && \
   go run consumer/main.go
   ```
2. connect localhost:5432 postgres and create users table and insert data
   ```sql
   create table users (id int primary key, name text not null, uid uuid not null, info jsonb not null, addresses text[] not null);
   insert into users(id, name, uid, info) values (1, 'kenny', 'bc03d615-8afb-452d-b0cc-340087def732', '{"myAge": 18}', '{"taipei", "hsinchu"}'); 
   ```
3. run scheduler to dump change to consumer
   ```bash
   go run scheduler/main.go
   ```

## How to change Postgres image version
You can use [postgres folder](postgres) to custom your postgres version with pglogcial and pgcapture extensions.

## How to change pgcapture image version
You can use [dockerbuild.sh](pgcapture/dockerbuild.sh) to custom your pgcapture version and change it in [docker-compose.yml](docker-compose.yml).
Since the pgcapture is still in development and is not stable, I recommend you always use the latest version of pgcapture.
Currently, the default pgcapture version is v0.0.56.
