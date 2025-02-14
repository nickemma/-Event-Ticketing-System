package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"

	"log"
	"os"
	"testing"
)

const (
	dbSource = "postgresql://root:securepassword@localhost:5432/event-ticketing?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}
	defer conn.Close()

	testQueries = New(conn)

	os.Exit(m.Run())
}
