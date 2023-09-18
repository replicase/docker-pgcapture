package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	pgtypeV4 "github.com/jackc/pgtype"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/replicase/pgcapture/pkg/pgcapture"
	"google.golang.org/grpc"
)

type User struct {
	ID        pgtype.Int4   `pg:"id"`
	Name      pgtypeV4.Text `pg:"name"`
	Uid       uuid.UUID     `pg:"uid"`
	Info      Info          `pg:"info"`
	Addresses []string      `pg:"addresses"`
}

type Info struct {
	MyAge int `json:"myAge"`
}

func (u *User) TableName() (schema, table string) {
	return "public", "users"
}

func (u *User) DebounceKey() string {
	return strconv.Itoa(int(u.ID.Int32))
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
			nu := change.New.(*User)
			ou := change.Old.(*User)
			fmt.Printf("id: %d, name: %s, uid: %s, info: %v addresses: %v\n", nu.ID.Int32, nu.Name.String, nu.Uid, nu.Info, nu.Addresses)
			fmt.Printf("id: %d, name: %s, uid: %s, info: %v addresses: %v\n", ou.ID.Int32, ou.Name.String, ou.Uid, ou.Info, ou.Addresses)
			return nil
		},
	})
	if err != nil {
		panic(err)
	}
}
