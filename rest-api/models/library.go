// Package models
// Created by GoLand.
// User: nixon
// Date: 17/5/2023
// Time: 13:18
package models

type Library struct {
	ID      string `json:"id" bson:"_id"`
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
	Books   []Book `json:"books" bson:"books"`
}
