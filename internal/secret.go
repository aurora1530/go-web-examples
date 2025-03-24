package internal

import (
	"fmt"
	"net/http"
)

func (s Server) GetApiSecretPage(w http.ResponseWriter, r *http.Request) {
	session, _ := s.Store.Get(r, "session")

	if auth,ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	username,ok := session.Values["username"].(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Welcome %s!", username)
}