package route

import "net/http"

type Route struct {
  Name        string  `json:"name"`
  Method      string  `json:"method"`
  Pattern     string  `json:"pattern"`
    HandlerFunc http.HandlerFunc
}

type Routes []Route
