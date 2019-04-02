package handler

import (
	"chi-example/platform/mock_http"
	"chi-example/platform/newsfeed"
	"net/http"
	"testing"
)

func TestNewsfeedGet(t *testing.T) {
	feed := newsfeed.New()
	feed.Add(newsfeed.Item{"hello", "world"})

	handler := NewsfeedGet(feed)

	w := &mock_http.ResponseWriter{}
	r := &http.Request{}

	handler(w, r)

	result := w.GetBodyJSONArray()

	if len(result) != 1 {
		t.Errorf("Item was not added to the datastore")
	}

	if result[0]["title"] != "hello" {
		t.Errorf("Item was not properly set")
	}
}
