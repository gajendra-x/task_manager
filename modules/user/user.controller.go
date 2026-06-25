package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var reqBody CreateUserPayload
	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadGateway)
		return
	}

	response, err := CreateUSerService(reqBody)

	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create user", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(User{
		ID:        response.ID,
		Name:      response.Name,
		Email:     response.Email,
		CreatedAt: response.CreatedAt,
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	email := strings.TrimPrefix(r.URL.Path, "/api/user/")

	response, err := GetUSerService(email)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(User{
		ID:        response.ID,
		Name:      response.Name,
		Email:     response.Email,
		CreatedAt: response.CreatedAt,
	})
}
