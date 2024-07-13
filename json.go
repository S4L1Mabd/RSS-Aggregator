// This File will contain the functions That Handle The JSON Resport  , S4L1M //

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responsewithError(w http.ResponseWriter, code int, msg string) {

	if code > 499 {

		log.Printf("Ooooops , 500 Error ")
		return
	}
	type errortype struct {
		Error string `json:"error"` //  This is m3natha in Json we will respose with {  error : "the error"} error is the key of our errortype.Error
	}

	responsewithJSON(w, code, errortype{
		Error: msg,
	})

}

func responsewithJSON(w http.ResponseWriter, code int, payload interface{}) {

	dat, err := json.Marshal(payload) // nakatbo data fa dat in json format

	if err != nil {

		log.Printf("Ooooops , 500 Error ")
		w.WriteHeader(500) // retourner code 500 if we have error to say that the error is internel
		return
	}
	log.Printf("a response is getting")
	w.Header().Add("Content-Type", "application/json") // we say that the respose is json
	w.WriteHeader(code)
	w.Write(dat) // we return the content

}
