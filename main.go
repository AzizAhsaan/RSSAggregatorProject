package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"database/sql"

	"github.com/AzizAhsaan/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct{
	DB *database.Queries
}

func main() {
	fmt.Println("hellower")

	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString ==""{
		log.Fatal("PORT  is not found in the environemtn")
	}

	godotenv.Load(".env")
	db_URL := os.Getenv("DB_URL")
	if db_URL ==""{
		log.Fatal("db_URL  is not found in the environemtn")
	}

	conn,err:=sql.Open("postgres",db_URL)
	if err != nil{
		log.Fatal("Can't connect ot database",err)
	}

	apiCfg :=apiConfig{
		DB:database.New(conn),
	}
	
	router := chi.NewRouter()
	srv := &http.Server{
		Handler :router,
		Addr: ":"+portString,
	}
	
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz",handlerReadiness)
	v1Router.Get("/err",handlerErr)
	v1Router.Post("/users",apiCfg.handlerCreateUser)
	v1Router.Get("/users",apiCfg.middlewareAuth(apiCfg.GetUserByApiKey))
	v1Router.Post("/feed",apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/getfeeds",apiCfg.handlerGetFeeds)
	v1Router.Post("/feed_follows",apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollows))
	v1Router.Get("/feed_follows",apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}",apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollows))

	router.Mount("/v1",v1Router)
	
	log.Printf("Server Starting on port %v",portString)
	err=srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("PORT:",portString)
}