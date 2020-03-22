package product

import (
	"encoding/json"
	"log"
	"net/http"

	"epc-test/services/product"
)

const (
	PATH = "/product"
)

func CalculateProductRoute(w http.ResponseWriter, r *http.Request) {
	var reqBody *CalculateProductDTO

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err, code := reqBody.Check(); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	offer, err := product.Calculate(&reqBody.Product, reqBody.Conditions)

	if err != nil {
		log.Println("Unexpected error while calculating offer", reqBody.Product, reqBody.Conditions)
		http.Error(w, "Can't calculate offer for given product and condition", http.StatusNotFound)
		return
	}

	if offer == nil || len(offer.Components) == 0 {
		http.Error(w, "Can't calculate offer for given product and condition", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(offer); err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
}
