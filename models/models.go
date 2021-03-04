package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ReqPost struct {
	ID        bson.ObjectId `bson:"id" json:"id"`
	URL       string        `bson:"url" json:"url"`
	Hash      RetHash
	CreatedAt time.Time `bson:"createdAt" json:"created_at"`
}

type RetHash struct {
	Hash string `bson:"hash" json:"hash"`
}
