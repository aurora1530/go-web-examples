package internal

import (
	"fmt"
	"net/http"
)

func (s Server) PostApiCreateUser(w http.ResponseWriter, r *http.Request){
	// get username and password from form

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// check if user already exists
	var exists bool
	err := s.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)", username).Scan(&exists)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error checking user existence", http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	// hash password

	hashed,err := HashPassword(password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// insert user into database
	_, err = s.DB.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, hashed)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}