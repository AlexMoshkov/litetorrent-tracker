package responses

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		ResponseServerError(w, err)
	}
}
