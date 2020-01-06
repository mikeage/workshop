package handlers

import (
	"encoding/json"
	"net/http"
)

//BasicHandler says "Hello World"
func BasicHandler(w http.ResponseWriter, r *http.Request) {
	ret := map[string]string{
		"key": "Hello World",
	}

	retRaw, err := json.Marshal(ret)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(retRaw)

}
