package main

import (
    "context"
    "fmt"
    "strconv"
    "time"

    "github.com/jackc/pgtype"
    "github.com/rueian/pgcapture/pkg/pgcapture"
    "google.golang.org/grpc"
)

type User struct {
    ID   pgtype.Int4 `pg:"id"`
    Name pgtype.Text `pg:"name"`
}

func (u *User) TableName() (schema, table string) {
    return "public", "users"
}

func (u *User) DebounceKey() string {
    return strconv.Itoa(int(u.ID.Int))
}

func main() {
    conn, err := grpc.Dial("localhost:10001", grpc.WithInsecure())
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    consumer := pgcapture.NewConsumer(context.Background(), conn, pgcapture.ConsumerOption{
        URI:              "postgres_cdc",
        TableRegex:       "users",
        DebounceInterval: time.Second,
    })
    defer consumer.Stop()

    err = consumer.Consume(map[pgcapture.Model]pgcapture.ModelHandlerFunc{
        &User{}: func(change pgcapture.Change) error {
            fmt.Println(change.New)
            return nil
        },
    })
    if err != nil {
        panic(err)
    }
}
