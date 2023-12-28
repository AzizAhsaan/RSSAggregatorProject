package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/AzizAhsaan/rssagg/internal/database"
)

func (apiCfg *apiConfig)handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request,user database.User){
	type parameters struct{
		FeedID uuid.UUID `json:"feed_id`
	}
	decoder := json.NewDecoder(r.Body)
	params :=parameters{}
	err:=decoder.Decode(&params)
	if err != nil{
		respondeWithError(w,400,fmt.Sprintf("errorrs pasing Json: %v",err))
		return 
	}

	feedFollow,err:= apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:uuid.New(),
		CreatedAt:time.Now().UTC(),
		UpdatedAt:time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,

	})
	if err != nil{
		respondeWithError(w,400,fmt.Sprintf("Couldn't create feed follow %s",err))
	}
	respondWithJSON(w,201,databaseFeedFollowToFeedFollow(feedFollow))
}


func (apiCfg *apiConfig)handlerGetFeedFollows(w http.ResponseWriter, r *http.Request,user database.User){

	feedFollows,err:= apiCfg.DB.GetFeedFollows(r.Context(),user.ID)
	if err != nil{
		respondeWithError(w,400,fmt.Sprintf("Couldn't get feed follow %s",err))
	}
	respondWithJSON(w,201,databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig)handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request,user database.User){
	feedFollowIDStr := chi.URLParam(r,"feedFollowID")
	feedFollowID,err := uuid.Parse(feedFollowIDStr)
	if err != nil{
		respondeWithError(w,400,fmt.Sprintf("Couldn't parse feed follow id: %v",err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:feedFollowID,
		UserID:user.ID,
	})
	if err != nil{
		respondeWithError(w,400,fmt.Sprintf("couldn't delete feed follow:%v"))
	}
	respondWithJSON(w,200,struct{}{})


}
