package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/lunatictiol/clean-architecture-in-go/configs"
)

func main() {
	var server configs.Server
	dsn := "postgres://postgres:postgres@localhost:5432/myapp_db"
	var err error
	DB, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server s80")
	server.New(":8080", DB)
	fmt.Println("server running on port 8080")
	err = server.Run()
	if err != nil {
		log.Fatalf("unable to start the server error: %v", err)
		return
	}
}
