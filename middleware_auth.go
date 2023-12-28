package main

import (
	"net/http"
	"fmt"
	"github.com/AzizAhsaan/rssagg/internal/auth"

	"github.com/AzizAhsaan/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		apiKey,err := auth.GetAPIKey(r.Header)
		if err != nil{
			respondeWithError(w,400,fmt.Sprintf("Auth error:%v",err))
			return
		}
		user,err := apiCfg.DB.GetUserByApiKey(r.Context(),apiKey)
		if err != nil{
			respondeWithError(w,400,fmt.Sprintf("Couldn't get user:%v",err))
			return
		}

		handler(w,r,user)
	}
}