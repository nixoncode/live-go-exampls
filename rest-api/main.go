// Package rest_api
// Created by GoLand.
// User: nixon
// Date: 17/5/2023
// Time: 12:00
package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/nixoncode/live-go-exampls/database"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	err := initApp()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		collection := database.GetCollection("todos")
		sampleDoc := bson.M{"name": "run 10,000 kilometers"}
		nDoc, err := collection.InsertOne(context.TODO(), sampleDoc)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error inserting todo")
		}

		return c.JSON(nDoc)
	})

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func initApp() error {
	err := loadEnv()
	if err != nil {
		return err
	}
	err = database.StartMongoDB()
	if err != nil {
		return err
	}
	return nil

}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
