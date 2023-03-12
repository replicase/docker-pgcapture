# docker-pgcapture

Demo [pgcapture](https://github.com/rueian/pgcapture) amazing library by docker !

## Demo cdc consumer
1. ```bash
   # build pgcapture image if you do not have pgcapture image in your local.
   ./pgcapture/dockerbuild.sh
   ./demo-consumer.sh
   go run consumer/main.go
   ```
2. connect localhost:5432 postgres and create users table and insert data
   ```sql
   create table users (id int primary key, name text not null);
   insert into users(id, name) values (1, 'kenny'); 
   ```

## Demo cdc consumer with scheduler
1. ```bash
   # build pgcapture image if you do not have pgcapture image in your local.
   ./demo-scheduler.sh
   go run consumer/main.go
   ```
2. connect localhost:5432 postgres and create users table and insert data
   ```sql
   create table users (id int primary key, name text not null);
   insert into users(id, name) values (1, 'kenny'); 
   ```
3. run scheduler to dump change to consumer
   ```bash
   go run scheduler/main.go
   ```

## How to change Postgres image version
You can use [Dockerfile](postgres/Dockerfile) to custom your postgres version with pglogcial and pgcapture extensions.

## How to change pgcapture image version

You can use [dockerbuild.sh](pgcapture/dockerbuild.sh) to custom your pgcapture version. Currently, the default pgcapture version is v0.0.40.
