package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gothello/go-encurtador/utils"
	"gopkg.in/mgo.v2/bson"
)

type reqPost struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	URL       string        `bson:"url" json:"url"`
	Hash      retHash
	CreatedAt time.Time `bson:"createdAt" json:"created_at"`
}

type retHash struct {
	Hash string `bson:"hash" json:"hash"`
}

func GenerateHash(lenght int) retHash {
	alfanum := []rune("abcdefghijklmnopqrstuvxwyzABCDEFGHIJKLMNOPQRSTUVXWYZ123456789")

	h := make([]rune, lenght)

	for i := range h {
		rand.Seed(time.Now().UTC().UnixNano())
		h[i] = alfanum[rand.Intn(len(alfanum))]
	}

	rhash := retHash{
		Hash: string(h),
	}

	fmt.Println(rhash)

	return rhash
}

func Urls(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ToError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var rp reqPost

	err := json.NewDecoder(r.Body).Decode(&rp)
	if err != nil {
		utils.ToError(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(rp)
	hash := GenerateHash(7)

	payload, err := json.Marshal(hash)
	if err != nil {
		utils.ToError(w, err.Error(), http.StatusInternalServerError)
	}

	utils.ToJson(w, payload, http.
		StatusCreated)
}
