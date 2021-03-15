package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ReqPost struct {
	ID        bson.ObjectId `bson:"id" json:"id"`
	Url       string        `bson:"url" json:"url"`
	NewURL	  string 		`bson:"new_url" json:"newurl"` 
	Hash      RetHash
	CreatedAt time.Time `bson:"createdAt" json:"created_at"`
}

type RetHash struct {
	Hash string `bson:"hash" json:"hash"`
}
