package main

import (
	c "app/controllers"
	"app/services"
	"log"

	"github.com/joho/godotenv"
	"github.com/srutherhub/web-app/server"
)

func main() {
	initEnv()

	cfg := server.ServerConfig{Port: "5555"}

	s := server.New()

	supabaseClient := services.NewSupabaseClient()
	db, err := supabaseClient.Connect()

	if err != nil {
		panic(err)
	}

	supabaseAuthService := services.NewSupabaseAuthService(db)

	appServices := services.AppServices{Auth: supabaseAuthService}

	s.RegisterController(c.BaseController(appServices))
	s.RegisterController(c.ApiController(appServices))

	s.Start(cfg)
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func initDb() {

}
