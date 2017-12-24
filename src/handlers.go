package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (a *App) sdk(w http.ResponseWriter, r *http.Request) {
	log.Println("sdk route")
	a.Bucket.Upsert("u:kingarthur",
		User{
			Id:        "kingarthur",
			Email:     "kingarthur@couchbase.com",
			Interests: []string{"Holy Grail", "African Swallows"},
		}, 0)

	var inUser User
	a.Bucket.Get("u:kingarthur", &inUser)
	response, _ := json.Marshal(inUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
