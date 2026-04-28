package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MaminirinaEdwino/backmadacare/src/config"
	"github.com/MaminirinaEdwino/backmadacare/src/controllers"
	"github.com/MaminirinaEdwino/backmadacare/src/routes"
	"github.com/MaminirinaEdwino/gobayes"
)



func main() {
	fmt.Println("back Mada care AI system")
	controllers.Network = gobayes.NewNetwork()
	config.SetupMedicalNetwork(controllers.Network)
	err := config.SyncMedicalRules(controllers.Network, "src/config/data/rules.json")
	if err != nil {
		log.Fatal("erreur : ", err)
	}
	mux := http.NewServeMux()
	routes.PredictRegisterRoutes(mux)
	routes.RegisterRoutesEtablissement(mux)
	routes.RegisterRoutesAdmin(mux)
	routes.RegisterRoutesPersonnel(mux)
	routes.RegisterRoutesPatient(mux)
	routes.RegisterRoutesAmbulance(mux)
	log.Println("serveur: localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}