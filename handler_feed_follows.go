package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	//"github.com/kush-316/RSS_Agg/internal/auth"
	"github.com/kush-316/RSS_Agg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct{
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err:= decoder.Decode(&params)
	if err!=nil {
		RespondWithError(w, 400, fmt.Sprintf("Error passing JSON: %v", err))
		return
	}
	
	feedFollow, err :=apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err!= nil{
		RespondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}


func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User){
	
	feedFollows, err :=apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err!= nil{
		RespondWithError(w, 400, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User){
	feedFollowStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowStr)
	if err!= nil{
		RespondWithError(w, 400, fmt.Sprintf("Couldn't Parse feed follow id: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}

