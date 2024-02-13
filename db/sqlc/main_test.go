package db

import (
	"database/sql"
	"fmt"
	"learn-until-die/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")

	if err != nil {
		fmt.Println(err)
		log.Fatal("Error when load config")
	}

	testDb, err = sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal("Error to connect db")
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
