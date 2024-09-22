package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	Id        int    `json:"id"`
	Completed bool   `json:"completed"`
	Name      string `json:"name"`
}

func main() {
	var myName string = "John"
	myLastName := "Doe"
	fmt.Println("Hello " + myName + " " + myLastName)

	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world!"})
	})

	todos := []Todo{}

	// Get TODOS
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// Create TODO
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Name == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo name is required."})
		}

		todo.Id = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	// Update TODO
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.Id) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(400).JSON(fiber.Map{"error": "Todo is not found."})
	})

	// Delete TODO
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.Id) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": true})
			}
		}
		return c.Status(400).JSON(fiber.Map{"error": "Todo is not found."})
	})

	PORT := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + PORT))
}
