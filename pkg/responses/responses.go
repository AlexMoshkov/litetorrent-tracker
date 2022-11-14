package responses

import (
	"log"
	"net/http"
)

func ResponseError(w http.ResponseWriter, code int, err error) {
	ResponseJSON(w, code, map[string]string{"error": err.Error()})
}

func ResponseOK(w http.ResponseWriter) {
	ResponseJSON(w, 200, map[string]string{"status": "ok"})
}

func ResponseServerError(w http.ResponseWriter, err error) {
	log.Println(err)
	ResponseJSON(w, 500, map[string]string{"error": "internal server error"})
}
