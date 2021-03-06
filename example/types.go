package example

import (
	"fmt"
)

type NotAuthorized struct {
	Message string `json:"message"`
}

func (e NotAuthorized) Error() string {
	return fmt.Sprintf("error: message=%q", e.Message)
}

type NotFound struct {
	Message string `json:"message"`
}

func (e NotFound) Error() string {
	return fmt.Sprintf("error: message=%q", e.Message)
}

type PatchPostRequest struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

type PatchUserRequest struct {
	Name string `json:"name"`
}

type Post struct {
	Body  string `json:"body"`
	Id    string `json:"id"`
	Title string `json:"title"`
}

type PutPostRequest struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

type PutUserRequest struct {
	Name string `json:"name"`
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
