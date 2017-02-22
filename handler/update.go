package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	// validation
	// connect to db
	// replace the fields
	// commit
	// close

	vars := mux.Vars(r)
	qualifier := vars["qualifier"]
	fmt.Fprintln(w, "Update", qualifier)
}
