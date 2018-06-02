package main

import (
	"fmt"
	"net/http"
)

func err(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Woops! The requested page doesn't exist yet!")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Flare, a place dedicated discovering (and collaborating on) new projects... :)")
	fmt.Fprintf(w, "\n\nThis is --obviously-- a work in progress.")

}
