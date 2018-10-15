package main

import (
	"gin-example/app"
	"gin-example/app/data"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

func init() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "debug"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	viper.SetConfigFile("./config/config." + env + ".json")
	viper.ReadInConfig()
	viper.Set("port", port)
}

func main() {
	db, err := data.NewDB(viper.GetString("database.meta_connection"))
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	handler := app.Handler{
		UserService: db,
		BookService: db,
	}

	ginEngine := app.GinEngine(&handler)
	if err := http.ListenAndServe(":"+viper.GetString("port"), ginEngine); err != nil {
		log.Fatal(err)
	}
}
