package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/bmi", ApiHandler)
	http.ListenAndServe(":3000", nil)
}