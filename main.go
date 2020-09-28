package main

import (
	"log"
	"net/http"
	"os"
	"product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	http.ListenAndServe(":9090", nil)
}