package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/hiteshjain48/RSA-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err!=nil {
		respondWithErr(w, 400, fmt.Sprintf("Error parsing JSON: %s",err))
		return
	}
	feed_follow, err := apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:			uuid.New(),
		CreatedAt: 	time.Now().UTC(),
		UpdatedAt: 	time.Now().UTC(),
		FeedID:		params.FeedID,
		UserID: 	user.ID,
	})
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error following feed: %s",err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feed_follow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feed_follows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err!=nil {
		respondWithErr(w, 400, fmt.Sprintf("Couldn't get feeds following: %s",err))
	}

	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feed_follows))
}


func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Couldn't parse feed_id: %s",err))
		return
	}
	
	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:		feedFollowID,
		UserID: user.ID,
	})
	if err!=nil {
		respondWithErr(w, 400, fmt.Sprintf("Error in deleting: %s", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}