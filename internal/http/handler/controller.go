package handler

import (
	"encoding/json"
	"net/http"
	"usermanager/internal/domain/request"
	"usermanager/internal/services"
	"usermanager/internal/validation"
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

	err := validation.ValidateNewUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	errServ := s.Services.Create(user)
	if errServ != nil {
		http.Error(w, errServ.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User created successfully")

}

func (s *UsersManagerServices) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user request.Users

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	errServ := s.Services.Get(user.Email, user.Password)
	if errServ != nil {
		http.Error(w, errServ.Error(), http.StatusUnauthorized)
		return
	}

	resonse := map[string]string{
		"message": "Login successful",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resonse)
}
