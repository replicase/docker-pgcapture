package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/replicase/pgcapture/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"
)

func main() {
	ctx := context.Background()
	pgConn, err := pgx.Connect(ctx, "postgres://postgres@localhost:5433/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer pgConn.Close(ctx)

	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewDBLogControllerClient(conn)

	table := "users"
	pages, err := getTablePages(ctx, pgConn, table)
	if err != nil {
		panic(err)
	}

	batch := uint32(1)
	var dumps []*pb.DumpInfoResponse

	for i := uint32(0); i < uint32(pages); i += batch {
		dumps = append(dumps, &pb.DumpInfoResponse{Schema: "public", Table: table, PageBegin: i, PageEnd: i + batch - 1})
	}

	uri := "postgres_cdc"
	if _, err = client.Schedule(context.Background(), &pb.ScheduleRequest{Uri: uri, Dumps: dumps}); err != nil {
		panic(err)
	}

	if _, err = client.SetScheduleCoolDown(context.Background(), &pb.SetScheduleCoolDownRequest{
		Uri: uri, Duration: durationpb.New(time.Second * 5),
	}); err != nil {
		panic(err)
	}
}

func getTablePages(ctx context.Context, conn *pgx.Conn, table string) (int, error) {
	if _, err := conn.Exec(ctx, "analyze "+table); err != nil {
		return 0, err
	}

	var pages int
	if err := conn.QueryRow(ctx, fmt.Sprintf("select relpages from pg_class where relname = '%s'", table)).Scan(&pages); err != nil {
		return 0, err
	}
	return pages, nil
}
