package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, 400, "Something went wrong")
}