package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
)


func main() {
	fmt.Println("hellower")

	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString ==""{
		log.Fatal("PORT  is not found in the environemtn")
	}
	
	router := chi.NewRouter()
	srv := &http.Server{
		Handler :router,
		Addr: ":"+portString,
	}
	
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz",handlerReadiness)
	v1Router.Get("/err",handlerErr)
	router.Mount("/v1",v1Router)
	
	log.Printf("Server Starting on port %v",portString)
	err:=srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("PORT:",portString)
}