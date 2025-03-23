// Code generated by goa v3.20.0, DO NOT EDIT.
//
// HTTP request path constructors for the todo service.
//
// Command:
// $ goa gen backend/design

package server

import (
	"fmt"
)

// CreateTodoPath returns the URL path to the todo service create HTTP endpoint.
func CreateTodoPath(userID int) string {
	return fmt.Sprintf("/users/%v/todos", userID)
}

// ListTodoPath returns the URL path to the todo service list HTTP endpoint.
func ListTodoPath(userID int) string {
	return fmt.Sprintf("/users/%v/todos", userID)
}

// GetTodoPath returns the URL path to the todo service get HTTP endpoint.
func GetTodoPath(userID int, todoID int) string {
	return fmt.Sprintf("/users/%v/todos/%v", userID, todoID)
}

// UpdateTodoPath returns the URL path to the todo service update HTTP endpoint.
func UpdateTodoPath(userID int, todoID int) string {
	return fmt.Sprintf("/users/%v/todos/%v", userID, todoID)
}

// DeleteTodoPath returns the URL path to the todo service delete HTTP endpoint.
func DeleteTodoPath(userID int, todoID int) string {
	return fmt.Sprintf("/users/%v/todos/%v", userID, todoID)
}
