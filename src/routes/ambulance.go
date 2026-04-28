package routes

import (
	"net/http"

	"github.com/MaminirinaEdwino/backmadacare/src/controllers"
)

func RegisterRoutesAmbulance(mux *http.ServeMux) {
	mux.HandleFunc("GET /ambulances", controllers.GetAllAmbulance)
	mux.HandleFunc("GET /ambulances/{id}", controllers.GetOneAmbulance)
	mux.HandleFunc("POST /ambulances", controllers.CreateAmbulance)
	mux.HandleFunc("PUT /ambulances/{id}", controllers.UpdateAmbulance)
	mux.HandleFunc("DELETE /ambulances/{id}", controllers.DeleteAmbulance)

	// Endpoint spécifique pour la gestion opérationnelle
	mux.HandleFunc("GET /ambulances/disponibles", controllers.GetAvailableAmbulance)
}