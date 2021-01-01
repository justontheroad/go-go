package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func Show(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Show!")
}
