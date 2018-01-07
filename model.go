package main


type First struct {
	Type	   string 	`json:"type,omitempty"`
	Label      string   `json:"label"`
	Spectators []string `json:"spectators"`
}
