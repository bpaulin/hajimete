package main

import (
	"github.com/gorilla/mux"
	"gopkg.in/couchbase/gocb.v1"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	Bucket *gocb.Bucket
}

func (a *App) Initialize(address, user, password, bucketName string) {
	cluster, err := gocb.Connect(address)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("connected to couchbase!")

	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: user,
		Password: password,
	})
	a.Bucket, err = cluster.OpenBucket(bucketName, "")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("connected to bucket!")

	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/first", a.getFirstList).Methods("GET")
	a.Router.HandleFunc("/first", a.createFirst).Methods("POST")
	a.Router.HandleFunc("/first/{id}", a.getFirst).Methods("GET")
	a.Router.HandleFunc("/first/{id}", a.updateFirst).Methods("PUT")
	a.Router.HandleFunc("/first/{id}", a.deleteFirst).Methods("DELETE")
}

func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}
