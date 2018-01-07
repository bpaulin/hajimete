package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) getFirstList(w http.ResponseWriter, r *http.Request) {
	log.Println("get first list route")
	w.WriteHeader(http.StatusNotImplemented)
}

func (a *App) createFirst(w http.ResponseWriter, r *http.Request) {
	log.Println("create first route")
	var first First

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&first); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	first.Type = "first"

	cbKey := uuid.NewV4().String()
	if _, err := a.Bucket.Insert("first::"+cbKey, first, 0); err != nil {
		respondWithError(w, http.StatusConflict, err.Error())
		return
	}

	w.Header().Set("location", "/first/"+cbKey)
	w.WriteHeader(http.StatusCreated)
}

func (a *App) getFirst(w http.ResponseWriter, r *http.Request) {
	log.Println("get first route")
	var first First
	params := mux.Vars(r)

	if _, err := a.Bucket.Get("first::"+params["id"], &first); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	first.Type = ""

	respondWithJSON(w, http.StatusOK, first)
}

func (a *App) updateFirst(w http.ResponseWriter, r *http.Request) {
	log.Println("update first route")
	w.WriteHeader(http.StatusNotImplemented)
}

func (a *App) deleteFirst(w http.ResponseWriter, r *http.Request) {
	log.Println("delete first route")
	w.WriteHeader(http.StatusNotImplemented)
}
