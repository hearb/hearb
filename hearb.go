package main

import (
	"log"

	"github.com/hearb/hearb/route"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("cannot load \".env\" file")
		panic(err)
	}

	app := route.Init()
	app.Run()
}
