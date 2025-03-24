package internal

import (
	"fmt"
	"net/http"
)

func (s Server) PostApiLoginUser(w http.ResponseWriter, r *http.Request){
	// get username and password from form

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// check if user exists
	var exists bool
	err := s.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)", username).Scan(&exists)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error checking user existence", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "Invalid ether username or password", http.StatusUnauthorized)
		return
	}

	// verify password
	var hashed []byte
	err = s.DB.QueryRow("SELECT password_hash FROM users WHERE username=$1", username).Scan(&hashed)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error retrieving password hash", http.StatusInternalServerError)
		return
	}
	ok := VerifyPassword(hashed, password)
	if !ok {
		http.Error(w, "Invalid ether username or password", http.StatusUnauthorized)
		return
	}

	// ok
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}