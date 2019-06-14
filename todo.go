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
	Latitude 	string		`json:"latitude"`
	Longitude 	string 	`json:"longitude"`
	Timestamp   string 	`json:"timestamp"`
}

type Path []Location