package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VanshAwasthi/go-moviesstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegistermoviesStoreRoutes(r)

	fmt.Println("Server starting on port 9000")
	err := http.ListenAndServe(":9000", r)
	if err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}
}
