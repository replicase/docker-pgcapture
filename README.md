# docker-pgcapture

It's a guidance to setup [pgcapture](https://github.com/replicase/pgcapture) environment by docker.

## Build pgcapture image
Build pgcapture image if you do not have pgcapture image in your local environment.
And you can change the pgcapture version in [dockerbuild.sh](pgcapture/dockerbuild.sh).
```bash
  (cd pgcapture && ./dockerbuild.sh)
```

## Demo CDC Consumer
1. Run the script
   ```bash
   # default postgres version is 14 and decode plugin is pgoutput
   # you can specify postgres version and decode plugin by setting environment variables
   # example: POSTGRES_VERSION=11 DECODE_PLUGIN=pglogical_output ./demo-consumer.sh
   ./demo-consumer.sh
   ```
2. Run the consumer
   ```bash
   go run consumer/main.go
   ```
3. Connect localhost:5432 postgres and create users table and insert data
   ```sql
   create table users (id int primary key, name text not null, uid uuid not null, info jsonb not null, addresses text[] not null);
   insert into users(id, name, uid, info) values (1, 'foo', 'bc03d615-8afb-452d-b0cc-340087def732', '{"myAge": 18}', '{"taipei", "hsinchu"}'); 
   ```
4. See the consumer output from postgres source change

## Demo CDC Consumer and Scheduler Dump
1. Run the script
   ```bash
   # default postgres version is 14 and decode plugin is pgoutput
   # you can specify postgres version and decode plugin by setting environment variables
   # example: POSTGRES_VERSION=11 DECODE_PLUGIN=pglogical_output ./demo-scheduler.sh
   ./demo-scheduler.sh
   ```
2. Run the consumer
    ```bash
    go run consumer/main.go
    ```
3. Connect localhost:5432 postgres and create users table and insert data
   ```sql
   create table users (id int primary key, name text not null, uid uuid not null, info jsonb not null, addresses text[] not null);
   insert into users(id, name, uid, info) values (1, 'foo', 'bc03d615-8afb-452d-b0cc-340087def732', '{"myAge": 18}', '{"taipei", "hsinchu"}'); 
   ```
4. See the consumer output from postgres source change
5. Run scheduler to dump change to consumer
   ```bash
   go run scheduler/main.go
   ```
6. See the consumer output from scheduler dump change 

## How to change Postgres image version
You can use [postgres folder](postgres) to custom your postgres version with pglogcial and pgcapture extensions.
