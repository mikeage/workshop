package handlers

import (
	"context"
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
		users, err := getAllUsers(w)
		if err != nil {
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}

func getAllUsers(w http.ResponseWriter) ([]types.User, error) {
	var users []types.User
	cursor, err := crud.GetAll("users", w)

	if err != nil {
		return users, err
	}

	err = cursor.All(context.TODO(), &users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	return users, err

}
