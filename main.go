package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
)

var err error

func main() {
	_, err = sql.Open("mysql", "user1:user1@tcp(127.0.0.1:3307)/docker_db")
	if err != nil {
		log.Println("error mysql connection init")
	} else {
		log.Println("mysql connection initiated")
	}

	opt, err := redis.ParseURL("redis://localhost:6380")
	if err != nil {
		log.Println("error connection redis parse")
	}
	redisClient := redis.NewClient(opt)
	if err := redisClient.Ping(); err != nil {
		log.Println("error redis connection init")
	} else {
		log.Println("redis connection initiated")
	}

	_, err = amqp.Dial("amqp://guest:guest@localhost:5673/")
	if err != nil {
		log.Println("error rabbitmq connection init")
	} else {
		log.Println("rabbitmq connection init")
	}

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, ni****")
}
