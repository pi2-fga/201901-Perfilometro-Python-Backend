package main

import "time"

type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

type Location struct {
	Latitude 	float32		`json:"latitude"`
	Longitude 	float32 	`json:"longitude"`
	Timestamp   time.Time 	`json:"timestamp"`
}

type Path []Location