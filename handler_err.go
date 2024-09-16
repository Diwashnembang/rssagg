package main

import "net/http"

func handlerErr(w http.ResponseWriter,req *http.Request){
	respondWithError(w,500,"error of 500")
}