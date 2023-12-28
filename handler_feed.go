package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/AzizAhsaan/rssagg/internal/database"
)

func (apiCfg *apiConfig)handlerCreateFeed(w http.ResponseWriter, r *http.Request,user database.User){
	type parameters struct{
		Name string `json:"name"`
		URL string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params :=parameters{}
	err:=decoder.Decode(&params)
	if err != nil{
		respondeWithError(w,400,fmt.Sprintf("errorrs pasing Json: %v",err))
		return 
	}

	feed,err:= apiCfg.DB.CreatedFeed(r.Context(), database.CreatedFeedParams{
		ID:uuid.New(),
		CreatedAt:time.Now().UTC(),
		UpdatedAt:time.Now().UTC(),
		Name:params.Name,
		Url:params.URL,
		UserID: user.ID,

	})
	if err != nil{
		respondeWithError(w,400,fmt.Sprintf("Couldn't create feed %s",err))
	}
	respondWithJSON(w,201,databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig)handlerGetFeeds(w http.ResponseWriter, r *http.Request){
	feeds,err:= apiCfg.DB.GetFeeds(r.Context())
	if err != nil{
		respondeWithError(w,400,fmt.Sprintf("Couldn't get feeds %s",err))
	}
	respondWithJSON(w,201,databaseFeedsToFeeds(feeds))
}

