package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Kodnavis/face2face-backend/user-service/model"
	"github.com/Kodnavis/face2face-backend/user-service/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	Repo *repository.UserRepo
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Firstname string `json:"firstname" validate:"required,min=2,max=50"`
		Lastname  string `json:"lastname" validate:"required,min=2,max=50"`
		Login     string `json:"login" validate:"required,min=5,max=100"`
		Password  string `json:"password" validate:"required,min=8,max=72"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	validate := validator.New()

	if err := validate.Struct(body); err != nil {
		errors := err.(validator.ValidationErrors)

		var errorMessages []string
		for _, e := range errors {
			errorMessages = append(errorMessages, fmt.Sprintf(
				"Field %s failed validation: %s", e.Field(), e.Tag()))
		}

		http.Error(w, strings.Join(errorMessages, "; "), http.StatusBadRequest)
		return
	}

	user := model.User{
		ID:        uuid.New(),
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Login:     body.Login,
		Password:  body.Password,
	}

	err := u.Repo.Insert(r.Context(), user)
	if err != nil {
		log.Printf("user create error: %v", err)

		if isDuplicateKeyError(err) {
			http.Error(w, "Login already exists", http.StatusConflict)
			return
		}

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := struct {
		ID        string `json:"id"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Login     string `json:"login"`
	}{
		ID:        user.ID.String(),
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Login:     user.Login,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
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

func isDuplicateKeyError(err error) bool {
	var pgErr *pq.Error
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}
