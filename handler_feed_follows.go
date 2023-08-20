package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

// func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
// 	feeds, err := apiCfg.DB.GetFeed(r.Context())
// 	if err!=nil {
// 		respondWithErr(w, 400, fmt.Sprintf("Couldn't get feed: %s",err))
// 	}

// 	respondWithJSON(w, 201, databaseFeedsToFeeds(feeds))
// }
