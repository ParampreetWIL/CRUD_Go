package handlers

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/ParampreetWIL/CRUD_Go/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func initDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("[+] Database Connected")
	return conn
}

var conn = initDB()
var ctx = context.Background()
var db = database.New(conn)

// @Summary Show all tasks
// @Description get list of all tasks
// @Tags tasks
// @Produce  json
// @Success 200 {object} database.Task
// @Failure 400 {object} error
// @Router / [get]
func GetAllTasksHandler(c *fiber.Ctx) error {
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
func AddNewTaskHandler(c *fiber.Ctx) error {
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
func EditTaskHandler(c *fiber.Ctx) error {
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
func DeleteTaskHandler(c *fiber.Ctx) error {
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
