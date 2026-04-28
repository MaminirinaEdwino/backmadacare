package routes

import (
	"net/http"

	"github.com/MaminirinaEdwino/backmadacare/src/controllers"
)

func RegisterRoutesCapacite(mux *http.ServeMux) {
	mux.HandleFunc("GET /capacites", controllers.GetAllCapacite)
	mux.HandleFunc("GET /capacites/{id}", controllers.GetOneCapacite)
	mux.HandleFunc("POST /capacites", controllers.CreateCapacite)
	mux.HandleFunc("PUT /capacites/{id}", controllers.UpdateCapacite)
	mux.HandleFunc("DELETE /capacites/{id}", controllers.DeleteCapacite)
	
	// Route pour voir les capacités d'un établissement spécifique
	mux.HandleFunc("GET /etablissements/{id}/capacites", controllers.GetByEtablissementCapacite)
}