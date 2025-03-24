package internal

import "net/http"

func (s Server) LogoutUser(w http.ResponseWriter, r *http.Request) {
	session,_ := s.Store.Get(r, "session")
	session.Values["username"] = nil
	session.Values["authenticated"] = false
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, "Error saving session", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}