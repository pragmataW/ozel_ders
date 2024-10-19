package main

import (
	"example/controller"
	"example/pkg"
	"example/repository"
	"example/service"
	"os"

	"github.com/gofiber/fiber/v2"
)

var (
	host string
	port string
	dbname string
	user string
	password string
	sslMode string
	timeZone string
	secretKey string
)

func main() {
	repo, err := repository.New(repository.DbConf{
		Host: host,
		Port: port,
		DbName: dbname,
		User: user,
		Password: password,
		SslMode: sslMode,
		TimeZone: timeZone,
	})

	if err != nil{
		panic(err)
	}

	encryptor:= pkg.NewEncryptor(secretKey)
	jwtGenerator := pkg.NewJwtGenerator(secretKey)
	service := service.New(service.WithEncryptor(encryptor),
							service.WithRepo(repo),
							service.WithJwtGenerator(jwtGenerator))

	controller := controller.New(service)

	app := fiber.New()

	app.Post("/login", controller.Login)
	app.Post("/register", controller.Register)

	app.Listen(":8080")
}

func init() {
	host = os.Getenv("DB_HOST")
	port = os.Getenv("POSTGRES_PORT")
	dbname = os.Getenv("POSTGRES_DB")
	user = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	sslMode = os.Getenv("SSL_MODE")
	timeZone = os.Getenv("TIME_ZONE")
	secretKey = os.Getenv("SECRET_KEY")
}