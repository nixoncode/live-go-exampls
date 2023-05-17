// Package rest_api
// Created by GoLand.
// User: nixon
// Date: 17/5/2023
// Time: 12:00
package main

import (
	"github.com/joho/godotenv"
	"github.com/nixoncode/live-go-exampls/database"
)

func main() {

	err := initApp()
	if err != nil {
		panic(err)
	}

	app := generateApp()

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
