# docker-pgcapture

## Demo cdc consumer
1. ```bash
   ./demo-consumer.sh
   go run consumer/main.go
   ```
2. connect localhost:5432 postgres and run
   ```sql
   create table users (id int primary key, name text not null);
   insert into users(id, name) values (1, 'kenny'); 
   ```

## Demo cdc consumer with scheduler
1. ```bash
   ./demo-scheduler.sh
   go run consumer/main.go
   ```
2. connect localhost:5432 postgres and run
   ```sql
   create table users (id int primary key, name text not null);
   insert into users(id, name) values (1, 'kenny'); 
   ```
3. run scheduler to dump change to consumer
   ```bash
   go run scheduler/main.go
   ```
