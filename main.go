package main

import (
	"cook-helper/handler"
	"cook-helper/repository"
	"cook-helper/service"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const port = "3000"

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(
		repository.Config{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
			Password: os.Getenv("DB_PASSWORD"),
		})

	if err != nil {
		log.Fatalf("error initializing db: %s", err.Error())
	}

	repo := repository.New(db)
	services := service.New(repo)
	handlers := handler.New(services)
	router := handlers.InitRoutes()

	log.Printf("Serving on port:%s", port)
	http.ListenAndServe(net.JoinHostPort("localhost", port), router)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
