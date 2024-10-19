package main

import (
	"example/controller"
	"example/repository"
	"example/service"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	host     string
	port     string
	dbname   string
	user     string
	password string
	sslMode string
	timeZone string
)

func main() {
	time.Sleep(time.Second * 10)
	conf := repository.DbConf{
		Host: host,
		Port: port,
		DbName: dbname,
		User: user,
		Password: password,
		SslMode: sslMode,
		TimeZone: timeZone,
	}

	repo, err := repository.New(conf)
	if err != nil{
		panic(err)
	}
	service := service.New(repo)
	controller := controller.New(service)

	app := fiber.New()

	app.Post("/login", controller.Login)

	app.Listen(":8080")
}

func init() {
	godotenv.Load("../.env")

	host = os.Getenv("DB_HOST")
	port = os.Getenv("POSTGRES_PORT")
	dbname = os.Getenv("POSTGRES_DB")
	user = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	sslMode = os.Getenv("SSL_MODE")
	timeZone = os.Getenv("TIME_ZONE")
}
