// Package models
// Created by GoLand.
// User: nixon
// Date: 17/5/2023
// Time: 13:21
package models

type Book struct {
	ID     string `json:"id" bson:"_id"`
	Name   string `json:"name" bson:"name"`
	Author string `json:"author" bson:"author"`
	ISBN   string `json:"isbn" bson:"isbn"`
}
