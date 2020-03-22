package main

import (
	"log"
	"net/http"

	"epc-test/routes/product"
)

const (
	PORT = "8080"
)

func main() {

	http.HandleFunc(product.PATH, product.CalculateProductRoute)

	log.Println("** Service Started on Port " + PORT + " **")

	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}
}
