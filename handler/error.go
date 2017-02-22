package handler

import (
	"fmt"
	"net/http"
)

func Error(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This page is an error.")
}
