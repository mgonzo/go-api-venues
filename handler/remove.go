package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Remove(w http.ResponseWriter, r *http.Request) {
	// validation
	// connect to db
	// set delete boolean to true
	// commit
	// close

	vars := mux.Vars(r)
	qualifier := vars["qualifier"]
	fmt.Fprintln(w, "Remove", qualifier)
}
