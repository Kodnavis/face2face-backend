package handler

import (
	"fmt"
	"net/http"
)

type User struct{}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Created a user")
}

func (u *User) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all users")
}

func (u *User) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a user by ID")
}

func (u *User) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a user by ID")
}

func (u *User) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a user by ID")
}
