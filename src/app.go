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
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: user,
		Password: password,
	})
	log.Print("connected to couchbase!")

	a.Bucket, err = cluster.OpenBucket(bucketName, "")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("connected to bucket!")

	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/sdk", a.sdk).Methods("GET")
}

func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}
