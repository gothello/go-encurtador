package controllers

import (

	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gothello/go-encurtador/utils"
)

type reqPost struct {
	UserID string `json:"user_id"`
	OriginUrl string `json:"original_url"`
}

type reqGet struct {
	Hash string `json:"hash"`
}

func GetHash(userID string, originalUrl string) (reqGet, error) {
	retHash := reqGet{
		Hash: "abcdefghi",
	}

	fmt.Println(retHash)

	return retHash, nil
}

func Urls(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ToError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var rp reqPost

	json.NewDecoder(r.Body).Decode(&rp)

	hash, err := GetHash(rp.UserID, rp.OriginUrl)
	if err != nil {
		utils.ToError(w, err.Error(), http.StatusInternalServerError)
	}

	payload, err := json.Marshal(hash)
	if err != nil {
		utils.ToError(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println(string(payload))

	utils.ToJson(w, payload, http.StatusCreated)
}