package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/hiteshjain48/RSA-aggregator/internal/database"
)

type User struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	ApiKey string `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	user := User{
		ID:			dbUser.ID,
		CreatedAt: 	dbUser.CreatedAt,
		UpdatedAt: 	dbUser.UpdatedAt,
		Name: 		dbUser.Name,
		ApiKey: 	dbUser.ApiKey,
	}
	return user
}

type Feed struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	URL string `json:"url"`
	UserID uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	feed := Feed{
		ID:			dbFeed.ID,
		CreatedAt: 	dbFeed.CreatedAt,
		UpdatedAt: 	dbFeed.UpdatedAt,
		Name: 		dbFeed.Name,
		URL: 		dbFeed.Url,
		UserID: 	dbFeed.UserID,
	}
	return feed
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dBfeed := range dbFeeds {
		feeds = append(feeds,databaseFeedToFeed(dBfeed))
	}
	return feeds
}