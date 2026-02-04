package handler

import (
	"encoding/json"
	"net/http"
	"usermanager/internal/domain/request"
	"usermanager/internal/services"
	"usermanager/internal/validation"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UsersManagerServices struct {
	Services services.UsersManagerServices
}

func (s *UsersManagerServices) RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	var user request.Users

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := validation.ValidateData(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.Services.Create(user); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

func (s *UsersManagerServices) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user request.Users

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := s.Services.Login(user.Email, user.Password); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	resonse := map[string]string{
		"message": "Login successful",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resonse)
}

func (s *UsersManagerServices) AllUsers(w http.ResponseWriter, r *http.Request) {
	response, err := s.Services.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
