package main

import (
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {

	responsewithError(w, 400, "something wrong .. ")
}
