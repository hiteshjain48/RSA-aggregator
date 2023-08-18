package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/hiteshjain48/RSA-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerReadiness(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err!=nil {
		respondWithErr(w, 400, fmt.Sprintf("Error parsing JSON: %s",err))
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error creating user: %s",err))
		return
	}


	respondWithJSON(w, 200, databaseUserToUser(user))
}