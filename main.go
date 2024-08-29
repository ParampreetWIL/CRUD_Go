package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ParampreetWIL/CRUD_Go/auth"
	"github.com/ParampreetWIL/CRUD_Go/database"
	"github.com/ParampreetWIL/CRUD_Go/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
)

func initDB(viper *viper.Viper) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), viper.GetString("POSTGRESQL_URI"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("[+] Database Connected")
	return conn
}

//	@title			CRUD GO API
//	@version		1.0
//	@description	API for CRUD Operations or a TODO List.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Parampreet Singh Rai
//	@contact.email	parampreets.rai@thewitslab.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:3000
// @BasePath	/
func main() {
	app := fiber.New()
	viper := viper.New()
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	vault_client, err := auth.InitVault(viper)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var conn = initDB(viper)
	var db = database.New(conn)

	app.Use(cors.New())
	app.Static("/swagger", "./swagger")
	app.Get("/swagger/*", swagger.HandlerDefault)
	handlers.SetupRoutes(app, viper, db, vault_client)
	err = app.Listen(":" + viper.GetString("SERVER_PORT"))

	if err != nil {
		fmt.Println(err)
	}
}
