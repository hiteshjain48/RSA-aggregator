package main

import (
	"fmt"
	"net/http"

	"github.com/hiteshjain48/RSA-aggregator/internal/auth"
	"github.com/hiteshjain48/RSA-aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User) 

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithErr(w, 403, fmt.Sprintf("Auth error: %s",err))
			return
		}
		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err!=nil {
			respondWithErr(w, 400, fmt.Sprintf("Couldn't get user: %v",err))
		}
		handler(w,r,user)
	}
}