package main

import (
	"encoding/json"
	"net/http"
	"log"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Falied to marshal JSON respnes",payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter,code int , msg string){
	if code > 499 {
		log.Println("Responding with 5XX error",msg)
	}

	type errResponse struct{
		Error string `json:"error"`
	}

	respondWithJSON(w,code,errResponse{msg})
}