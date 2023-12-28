package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/AzizAhsaan/rssagg/internal/database"
)

func (apiCfg *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params :=parameters{}
	err:=decoder.Decode(&params)
	if err != nil{
		respondeWithError(w,400,fmt.Sprintf("errorrs pasing Json: %v",err))
		return 
	}

	user,err:=apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:uuid.New(),
		CreatedAt:time.Now().UTC(),
		UpdatedAt:time.Now().UTC(),
		Name:params.Name,

	})
	if err != nil{
		respondeWithError(w,400,fmt.Sprintf("Couldn't create user %s",err))
	}
	respondWithJSON(w,201,databaseUserToUser(user))
}

func (apiCfg *apiConfig)GetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User){
	respondWithJSON(w,200,databaseUserToUser(user))
}