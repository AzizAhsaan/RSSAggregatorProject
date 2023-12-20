package main

import (
	"encoding/json"
	"log"
	"net/http"
)
func respondeWithError(w http.ResponseWriter, code int ,msg string){
	if code > 499 {
		log.Println("Reponsding with 5XX error:", msg)
	}
	type errResponse struct{
		Error string `json:"error"`
		Success string `json:"success"`
	}
	respondWithJSON(w,code,errResponse{
		Error: "Always wrong",
		Success:msg,
	})
}

func respondWithJSON(w http.ResponseWriter,code int, payload interface{}){
	dat, err := json.Marshal(payload)
	if err != nil{
		log.Println("Falid to marshal JSON responsive: %v",payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(dat)
}