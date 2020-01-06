package handlers

import (
	"encoding/json"
	"net/http"
	"workshop/crud"
	"workshop/types"
)

// UserHandler does something
func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user types.User
		json.NewDecoder(r.Body).Decode(&user)
		if user.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("A user must have a name"))
			return
		}
		ret, err := crud.Insert("users", user, w)
		if err != nil {
			return
		}
		json.NewEncoder(w).Encode(ret)
	}
	if r.Method == http.MethodGet {
	}
}
