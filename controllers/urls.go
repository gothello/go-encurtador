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

var (

	hashes = map[string]bool{}

	seeders = []rune("abcdefghijklmnopqrstuvxwyzABCDEFGHIJKLMNOPQRSTUVXWYZ123456789")

)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func DeleteHash {
}

func GenerateHash(lenght int) models.RetHash {
	var rhash models.RetHash
	h := make([]rune, lenght)

	for {

		for i := range h {
			rand.Seed(time.Now().UTC().UnixNano())
			h[i] = seeders[rand.Intn(len(seeders))]
		}


		if _, ok := hashes[string(h)]; !ok {

			hashes[string(h)] = true

			rhash = models.RetHash{
				Hash: string(h),
			}

			break
		}
	}

	fmt.Printf("%+v\n", rhash)

	return rhash
}

func Urls(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ToError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	hash := GenerateHash(7)

	payload, err := json.Marshal(hash)
	if err != nil {
		utils.ToError(w, err.Error(), http.StatusInternalServerError)
	}

	utils.ToJson(w, payload, http.
		StatusCreated)
}
