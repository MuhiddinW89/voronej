package main

import (
	"app/api"
	"app/config"
	consoleinput "app/consoleInput"
	"app/storage/postgresql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	consoleinput.ConsoleInput()

	cfg := config.Load()

	store, err := postgresql.NewConnectPostgresql(&cfg)
	if err != nil {
		log.Panic("Connection to Postgresql failed")
		return
	}
	defer store.CloseDB()

	r := gin.New()

	r.Use(gin.Recovery(), gin.Logger())

	api.NewApi(r, &cfg, store)

	go func() {

		fmt.Println("Server runnung on port", cfg.PostgresHost+cfg.ServerPort)
		err = r.Run(cfg.ServerHost + cfg.ServerPort)
		if err != nil {
			log.Panic("Error listening server: ")
			return
		}
	}()

	consoleinput.ConsoleOutput()
	consoleinput.ConsoleOutput()
	consoleinput.ConsoleOutput()

}
