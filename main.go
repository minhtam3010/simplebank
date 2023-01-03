package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/minhtam3010/simplebank/api"
	db "github.com/minhtam3010/simplebank/db/sqlc"
	"github.com/minhtam3010/simplebank/db/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
