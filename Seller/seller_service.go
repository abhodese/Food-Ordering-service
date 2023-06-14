package seller

//create a seller service for order matching system  

import (
	"sort"
	"sync"
	"time"
	"net/http"
	"encoding/json"
	"fmt"
)

//create a seller struct

type Seller struct {
	// seller struct
	// seller has an id, price, quantity, time
	// seller id is unique

	Id       int
	Price    float64
	Quantity int
	Time     time.Time
}

func NewSeller(id int, price float64, quantity int, time time.Time) *Seller {
	// create a new seller
	// return a pointer to the seller

	return &Seller{
		Id:       id,
		Price:    price,
		Quantity: quantity,
		Time:     time,
	}
}

//create a product struct for order matchign service

type Product struct {
	// product struct
	// product has a product id and a seller service
	// product id is unique

	ID       int     `json:"id"`
	Name     string  `json:"name"`
	SellerService *SellerService
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

var products []Product



func NewProduct(id int) *Product {
	// create a new product

	return &Product{
		ID:       id,
		SellerService: NewSellerService(),
	}
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	products = append(products, product)
	w.WriteHeader(http.StatusCreated)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, product := range products {
		if fmt.Sprintf("%d", product.ID) == id {
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}



type SellerService struct {
	// seller service struct
	// seller service has a seller map and a mutex
	// seller map is a map of seller id and seller struct

	sellerMap map[int]*Seller
	mutex     sync.Mutex
}

func NewSellerService() *SellerService {
	// create a new seller service
	// return a pointer to the seller service
	return &SellerService{
		sellerMap: make(map[int]*Seller),
	}
}

func (s *SellerService) AddSeller(seller *Seller) {
	// add a seller to the seller service
	// seller service is protected by a mutex
	// seller service has a map of seller id and seller struct
	// seller id is the key of the map
	// seller struct is the value of the map
	// seller id is unique

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.sellerMap[seller.Id] = seller
}

func (s *SellerService) GetSeller(id int) *Seller {
	// get a seller from the seller service
	// seller service is protected by a mutex
	// seller service has a map of seller id and seller struct
	// seller id is the key of the map
	// seller struct is the value of the map
	// seller id is unique

	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.sellerMap[id]
}

func (s *SellerService) GetSellers() []*Seller {
	// get all sellers from the seller service
	// seller service is protected by a mutex
	// seller service has a map of seller id and seller struct
	// seller id is the key of the map
	// seller struct is the value of the map
	// seller id is unique

	s.mutex.Lock()
	defer s.mutex.Unlock()

	var sellers []*Seller
	for _, seller := range s.sellerMap {
		sellers = append(sellers, seller)
	}
	return sellers
}

func (s *SellerService) GetSellerIds() []int {
	// get all seller ids from the seller service
	// seller service is protected by a mutex
	// seller service has a map of seller id and seller struct
	// seller id is the key of the map
	// seller struct is the value of the map
	// seller id is unique

	s.mutex.Lock()
	defer s.mutex.Unlock()

	var sellerIds []int
	for sellerId := range s.sellerMap {
		sellerIds = append(sellerIds, sellerId)
	}
	return sellerIds
}


func (s *SellerService) GetSellerIdsSorted() []int {
	// get all seller ids from the seller service
	// seller service is protected by a mutex
	// seller service has a map of seller id and seller struct
	// seller id is the key of the map
	// seller struct is the value of the map
	// seller id is unique
	// seller ids are sorted

	sellerIds := s.GetSellerIds()
	sort.Ints(sellerIds)
	return sellerIds
}

func (s *SellerService) GetSellerIdsSortedByTime() []int {
	// get all seller ids from the seller service
	// seller service is protected by a mutex
	// seller service has a map of seller id and seller struct
	// seller id is the key of the map
	// seller struct is the value of the map
	// seller id is unique
	// seller ids are sorted by time

	sellerIds := s.GetSellerIds()
	sort.Slice(sellerIds, func(i, j int) bool {
		return s.sellerMap[sellerIds[i]].Time.Before(s.sellerMap[sellerIds[j]].Time)
	})
	return sellerIds
}
