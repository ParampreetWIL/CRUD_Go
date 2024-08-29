package handlers

import (
	"context"

	"github.com/ParampreetWIL/CRUD_Go/database"
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/vault/api"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupRoutes(app *fiber.App, viper *viper.Viper, db *database.Queries, vault_client *api.Client) {
	oauthConfig := &oauth2.Config{
		ClientID:     viper.GetString("CLIENT_ID"),
		ClientSecret: viper.GetString("CLIENT_SECRET"),
		RedirectURL:  viper.GetString("REDIRECT_URI"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	ctx := context.Background()

	app.Get("/", func(c *fiber.Ctx) error {
		return GetAllTasksHandler(c, db, ctx)
	})
	app.Post("/", func(c *fiber.Ctx) error {
		return AddNewTaskHandler(c, db, ctx)
	})
	app.Post("/edit", func(c *fiber.Ctx) error {
		return EditTaskHandler(c, db, ctx)
	})
	app.Delete("/:id", func(c *fiber.Ctx) error {
		return DeleteTaskHandler(c, db, ctx)
	})
	app.Get("/login", func(c *fiber.Ctx) error {
		return GoogleOAuthRedirect(c, db, ctx, oauthConfig)
	})
	app.Get("/oauth/redirect", func(c *fiber.Ctx) error {
		return GoogleOAuthLogin(c, db, ctx, oauthConfig, vault_client, viper)
	})
	app.Post("/profile", func(c *fiber.Ctx) error {
		return GetUserProfile(c, db, ctx, viper)
	})
}
