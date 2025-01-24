package main

import (
	"log"
	"net/http"

	"github.com/sufimalek/jwtapi/internal/api"
	"github.com/sufimalek/jwtapi/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	router := api.NewRouter(db)
	log.Printf("Server started on :%s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}
