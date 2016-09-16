package main

import (
	"log"

	"github.com/hearb/hearb/mysql"
	"github.com/hearb/hearb/redis"
	"github.com/hearb/hearb/route"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("cannot load \".env\" file")
		panic(err)
	}
	if err := mysql.NewDB(); err != nil {
		log.Println("cannot connect to mysql")
		panic(err)
	}
	if err := redis.NewClient(); err != nil {
		log.Println("cannot connect to redis")
		panic(err)
	}

	app := route.Init()
	app.Run()
}
