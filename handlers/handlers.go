package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/ParampreetWIL/CRUD_Go/auth"
	"github.com/ParampreetWIL/CRUD_Go/database"
	structures "github.com/ParampreetWIL/CRUD_Go/structs"
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/vault/api"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// @Summary Show all tasks
// @Description get list of all tasks
// @Tags tasks
// @Produce  json
// @Success 200 {object} database.Task
// @Failure 400 {object} error
// @Router / [get]
func GetAllTasksHandler(c *fiber.Ctx, db *database.Queries, ctx context.Context) error {
	tasks, err := db.GetAllTasks(ctx)
	fmt.Println(tasks)
	if err != nil {
		fmt.Println(err)
		c.SendStatus(500)
		return err
	}

	return c.JSON(tasks)
}

// @Summary Adds a task
// @Description Add new task to ToDo List as not done
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body database.AddTaskParams true "Task object"
// @Success 200 {object} database.Task
// @Failure 400 {object} error
// @Router / [post]
func AddNewTaskHandler(c *fiber.Ctx, db *database.Queries, ctx context.Context) error {
	// if id is valid integer in body then edit the current record with the given id
	task := new(database.AddTaskParams)
	if err := c.BodyParser(task); err != nil {
		c.SendStatus(500)
		return err
	}

	t, err := db.AddTask(ctx, *task)
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return c.JSON(t)
}

// @Summary Edit a task
// @Description Edit the task with the given id
// @Tags tasks
// @Accept  json
// @Param task body database.UpdateTaskParams true "Update Task Params"
// @Success 200
// @Failure 400 {object} error
// @Router /edit [post]
func EditTaskHandler(c *fiber.Ctx, db *database.Queries, ctx context.Context) error {
	task := new(database.UpdateTaskParams)
	if err := c.BodyParser(task); err != nil {
		c.SendStatus(500)
		return err
	}
	fmt.Println(task)
	err := db.UpdateTask(ctx, *task)
	if err != nil {
		c.SendStatus(500)
		return err
	}
	return c.SendStatus(200)
}

// @Summary Delete a task
// @Description Delete a task with given id
// @Tags tasks
// @Param id path int true "Task ID"
// @Success 200
// @Failure 400 {object} error
// @Router /{id} [delete]
func DeleteTaskHandler(c *fiber.Ctx, db *database.Queries, ctx context.Context) error {
	_id := c.Params("id")
	id, err := strconv.ParseInt(_id, 10, 64)

	if err != nil {
		c.SendStatus(500)
		return err
	}

	err = db.DeleteTasks(ctx, id)
	if err != nil {
		c.SendStatus(500)
		return err
	}
	return c.SendStatus(200)
}

func GoogleOAuthLogin(c *fiber.Ctx, db *database.Queries, ctx context.Context, oauthConfig *oauth2.Config, vault_client *api.Client, viper *viper.Viper) error {
	code := c.Query("code") //get code from query params for generating token
	if code == "" {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token")
	}
	token, err := oauthConfig.Exchange(context.Background(), code) //get token
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token: " + err.Error())
	}
	client := oauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user info: " + err.Error())
	}

	defer response.Body.Close()
	var user structures.User
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error reading response body: " + err.Error())
	}
	err = json.Unmarshal(bytes, &user) //unmarshal user info
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error unmarshal json body " + err.Error())
	}

	vault_token, err := auth.Tokenize(vault_client, user.Email)

	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
	}

	jwt_token, err := auth.GenerateJWT(user, viper.GetString("JWT_SECRET_KEY"))

	if err != nil {
		fmt.Println("Error in JWT Generate Token: ", err)
	}

	db.AddUser(ctx, database.AddUserParams{
		EmailToken: vault_token,
		JwtToken:   jwt_token,
		Name:       user.Name,
	})

	return c.Status(fiber.StatusOK).JSON(structures.JWTToken{
		AccessToken: jwt_token,
	}) //return user info
}

// @Summary Get the profile data without password
// @Description Get profile data with JWT Token
// @Tags JWT
// @Param jwt body structures.JWTToken true "JWT Token"
// @Accept json
// @Produce json
// @Success 302 {object} structures.User "User details"
// @Failure 401
// @Router /profile [post]
func GetUserProfile(c *fiber.Ctx, db *database.Queries, ctx context.Context, viper *viper.Viper) error {
	jwt_token := new(structures.JWTToken)
	c.BodyParser(jwt_token)

	user, err := auth.DecryptJWT(jwt_token.AccessToken, viper.GetString("JWT_SECRET_KEY"))

	if err == nil {
		return c.JSON(user)
	}

	fmt.Println(err)

	return c.SendStatus(401)
}

// @Summary Redirect to Google OAuth
// @Description Redirects the user to Google's OAuth 2.0 authentication page to start the login process.
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 302 {string} string "Redirected to Google OAuth"
// @Failure 500 {object} error "Internal Server Error"
// @Router /login [get]
func GoogleOAuthRedirect(c *fiber.Ctx, _ *database.Queries, _ context.Context, oauthConfig *oauth2.Config) error {
	url := oauthConfig.AuthCodeURL("state")
	fmt.Println("Login request")
	return c.Redirect(url)
}
