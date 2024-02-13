package main

import (
	"database/sql"
	"fmt"
	"learn-until-die/api"
	db "learn-until-die/db/sqlc"
	"learn-until-die/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load env file")
	}
	conn, err := sql.Open(config.DbDriver, config.DbSource)
	fmt.Println(err)
	if err != nil {
		log.Fatal("cannot connect db")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddr)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
