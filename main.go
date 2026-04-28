package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MaminirinaEdwino/backmadacare/src/routes"
)



func main() {
	fmt.Println("back Mada care AI system")

	mux := http.NewServeMux()
	routes.PredictRegisterRoutes(mux)

	log.Println("serveur: localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}