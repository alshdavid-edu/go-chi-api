package handler

import (
	"chi-example/platform/newsfeed"
	"encoding/json"
	"net/http"
)

func NewsfeedGet(feed newsfeed.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		json.NewEncoder(w).Encode(items)
	}
}
