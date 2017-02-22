package handler

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func Show(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    qualifier := vars["qualifier"]
    fmt.Fprintln(w, "Show", qualifier)
}
