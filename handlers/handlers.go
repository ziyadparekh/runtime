package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hi there, welcome to Runtime\n")
}
