package main

import "fmt"

var currentId int
var todos Todos // Todo 
var paths Path // Path

func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
	RepoCreatePath(Location{Latitude: "0.53", Longitude: "43.54", Timestamp: "dsfdf"})
	RepoCreatePath(Location{Latitude: "0.653", Longitude: "44.54", Timestamp: "dwedwed"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	return Todo{} // return empty Todo if not found
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t //this is bad, I don't think it passes race condtions
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}

func RepoCreatePath(l Location) Location {
	paths = append(paths, l)
	return l
}
