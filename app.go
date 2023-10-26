package main

import (
	context "GoCrudORM/context"
	"GoCrudORM/router"
	mainRoutes "GoCrudORM/router/main"
	"GoCrudORM/types"
	"log"
	"net"
)

func main() {
	dbContext, sqlDB, err := context.SetupDbContext()
	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()
	dbContext.AutoMigrate(&types.User{})

	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	servevrRouter := router.NewRouter()
	mainRoutes.AddRoutes(servevrRouter)

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go router.HandleRequest(conn, servevrRouter)
	}
}
