package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gothello/go-encurtador/models"
	"github.com/gothello/go-encurtador/utils"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func GenerateHash(lenght int) models.RetHash {
	alfanum := []rune("abcdefghijklmnopqrstuvxwyzABCDEFGHIJKLMNOPQRSTUVXWYZ123456789")

	h := make([]rune, lenght)

	for i := range h {
		rand.Seed(time.Now().UTC().UnixNano())
		h[i] = alfanum[rand.Intn(len(alfanum))]
	}

	rhash := models.RetHash{
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

	var rp models.ReqPost

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
