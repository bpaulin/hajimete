package main

type User struct {
	Id        string   `json:"uid"`
	Email     string   `json:"email"`
	Interests []string `json:"interests"`
}
