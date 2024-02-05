package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB
var driverName = "postgres"
var dataSource = "postgresql://postgres:cuankipintar@localhost:5434/simple_bank?sslmode=disable"

func TestMain(m *testing.M) {
	var err error

	testDb, err = sql.Open(driverName, dataSource)

	if err != nil {
		log.Fatal("Error to connect db")
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
