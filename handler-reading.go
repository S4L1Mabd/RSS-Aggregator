package main

import (
	"net/http"
)

func handlerReader(w http.ResponseWriter, r *http.Request) {

	responsewithJSON(w, 220, struct{}{})

}
