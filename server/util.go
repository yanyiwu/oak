package server

import (
	"encoding/json"
	"net/http"
)

func ResponseError(w http.ResponseWriter, err string) {
	result := struct {
		Error string `json:"error"`
	}{
		Error: err,
	}
	b, _ := json.Marshal(result)
	w.Write(b)
}
