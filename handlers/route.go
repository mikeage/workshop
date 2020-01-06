package handlers

import (
	"encoding/json"
	"net/http"
)

// BasicHandler does something
func BasicHandler(w http.ResponseWriter, r *http.Request) {
	ret := map[string]string{
		"key": "Hello world",
	}
	retRaw, err := json.Marshal(ret)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(retRaw))

}
