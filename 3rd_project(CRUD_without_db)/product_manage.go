package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	ID           string        `json:"id"`
	Address      string        `json:"address"`
	Name         string        `json:"title"`
	Manufacturer *Manufacturer `json:"manufacturer"`
}

type Manufacturer struct {
	Company_name    string `json:"company_name"`
	Company_address string `json:"company_address"`
}

// if the client side ask for the  all of the movies, just send the information stored in "products"slice as the form of json to client(via this func)
func getproducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// cancelling order via the id ( obtained from url ) //
func Cancel_order(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			deletedItem := item
			products = append(products[:index], products[index+1:]...)
			json.NewEncoder(w).Encode(deletedItem)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "product not found"})
}

// if the user wants a specific type of the product( via the id )then this func comes into play//
func get_Product(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "product not found"})
}

//if you want to create new product , then this func comes into the play , ( most imp ) //

func create_Product(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var products_new Product // (products_new is going to go for the Product struct so) //
	err := json.NewDecoder(r.Body).Decode(&products_new)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid JSON"})
		return
	}
	products_new.ID = strconv.Itoa(rand.Intn(999)) // decoding or extracting the json formatteddata given by the user ( asuumption like this was made because the methods is POST)

	products = append(products, products_new) //adding new product into the products slice as of now

	json.NewEncoder(w).Encode(products_new)
}

//if we want to update any kinds of the product then this func will come handy //

func update_Product(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			var products_new Product
			err := json.NewDecoder(r.Body).Decode(&products_new)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "invalid JSON"})
				return
			}
			products_new.ID = params["id"]
			products = append(products, products_new)
			json.NewEncoder(w).Encode(products_new)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "product not found"})
}

var products []Product

func main() {
	r := mux.NewRouter()
	products = append(products, Product{ID: "1", Address: "kharibot", Name: "CCA-Pheonix IEM", Manufacturer: &Manufacturer{Company_name: "CCA", Company_address: "Taiwan"}})

	products = append(products, Product{ID: "2", Address: "Banasthali", Name: "AULA_F75", Manufacturer: &Manufacturer{Company_name: "AULA", Company_address: "China"}})

	r.HandleFunc("/products", getproducts).Methods("GET")

	r.HandleFunc("/products/{id}", get_Product).Methods("GET")

	r.HandleFunc("/products", create_Product).Methods("POST")

	r.HandleFunc("/products/{id}", update_Product).Methods("PUT")

	r.HandleFunc("/products/{id}", Cancel_order).Methods("DELETE")

	fmt.Println("Starting the server from the port 5000")

	log.Fatal(http.ListenAndServe(":5000", r))

}
