package handlers

import (
	"log"
	"strconv"
	"net/http"
	"product-api/data"
	"github.com/gorilla/mux"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
//func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
//	// handle the request for a list of products
//	if r.Method == http.MethodGet {
//		p.getProducts(rw, r)
//		return
//	}
//
//	if r.Method == http.MethodPost {
//		p.addProduct(rw, r)
//		return
//	}
//
//	if r.Method == http.MethodPut {
//		reg := regexp.MustCompile( `/([0-9]+)`)
//		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
//
//		if len(g) != 1 {
//			http.Error(rw, "Invalid URI", http.StatusBadRequest)
//			return
//		}
//		if len(g[0]) != 2 {
//			http.Error(rw, "Invalid URI", http.StatusBadRequest)
//		}
//		idString := g[0][1]
//		id, err := strconv.Atoi(idString)
//		if err != nil {
//			http.Error(rw,"Invalid URI", http.StatusBadRequest)
//			return
//		}
//		p.updateProduct(id, rw, r)
//		return
//	}
//
//	// catch all
//	// if no method is satisfied return an error
//	rw.WriteHeader(http.StatusMethodNotAllowed)
//}

// getProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Println("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id to string.", http.StatusBadRequest)
	}

	p.l.Println("Handle Update Products", id)

	prod := &data.Product{}
	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	err = data.UpdateProduct(prod, id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found.", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found.", http.StatusBadRequest)
		return
	}
}

type KeyProduct struct {

}

//func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
//	return http.HandlerFunc(rw http.ResponseWriter, r *http.Request) {
//		prod := &data.Product{}
//
//		err := prod.FromJSON(rw, "Unable to unmarshal json", http.StatusBadRequest)
//		return
//	}
//	ctx := r.Context().WithValue(KeyProduct{}, prod)
//	req := r.WithContext(ctx)
//
//	next.ServeHTTP(rw, req)
//}