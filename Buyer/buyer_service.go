package buyer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//create a buyer struct

type Buyer struct {
	// buyer struct
	// buyer has an id, price, quantity, time
	// buyer id is unique

	Id       int
	Price    float64
	Quantity int
	Time     time.Time
}

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Order struct {
	ID     int `json:"id"`
	Amount int `json:"amount"`
}

func SearchProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sellerURL := "http://seller-service:8080/products" // Assumes seller service is running in a container with the name "seller-service"
	resp, err := http.Get(sellerURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var products []Product
	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Filtering products based on query
	var filteredProducts []Product
	for _, product := range products {
		if product.Name == query {
			filteredProducts = append(filteredProducts, product)
		}
	}

	json.NewEncoder(w).Encode(filteredProducts)
}

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("product_id")
	amountStr := r.URL.Query().Get("amount")
	if productID == "" || amountStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sellerURL := fmt.Sprintf("http://seller-service:8080/product?id=%s", url.QueryEscape(productID))
	resp, err := http.Get(sellerURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var product Product
	err = json.NewDecoder(resp.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if product.Quantity < amount {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	order := Order{
		ID:     productIDInt,
		Amount: amount,
	}

	orderBytes, err := json.Marshal(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	orderURL := "http://order-service:8080/order"

	resp, err = http.Post(orderURL, "application/json", bytes.NewBuffer(orderBytes))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

