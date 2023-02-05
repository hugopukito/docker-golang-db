package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
)

var err error
var db *sql.DB

func main() {
	db, err = sql.Open("mysql", "user1:user1@tcp(mysql:3306)/docker_db")
	if err != nil {
		log.Println("error mysql connection init" + err.Error())
	}

	opt, err := redis.ParseURL("redis://redis:6379")
	if err != nil {
		log.Println("error connection redis parse")
	}
	redisClient := redis.NewClient(opt)
	if err := redisClient.Ping(); err.String() != "ping: PONG" {
		log.Println("error redis connection init: " + err.String())
	}

	_, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672")
	if err != nil {
		log.Println("error rabbitmq connection init" + err.Error())
	}

	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/test-sql", TestSQL)
	http.ListenAndServe(":8095", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, ni****")
}

func TestSQL(w http.ResponseWriter, r *http.Request) {
	db.Exec("CREATE TABLE IF NOT EXISTS containers (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL)")

	db.Exec("INSERT INTO containers (name) VALUE ('Armand')")

	fmt.Fprintf(w, "sql")
}
