package routes

import (
	"net/http"

	"github.com/MaminirinaEdwino/backmadacare/src/controllers"
)

func RegisterRoutesPersonnel(mux *http.ServeMux) {
	mux.HandleFunc("GET /personnel", controllers.GetAllPersonnel)
	mux.HandleFunc("GET /personnel/{id}", controllers.GetOnePersonnel)
	mux.HandleFunc("POST /personnel", controllers.CreatePersonnel)
	mux.HandleFunc("PUT /personnel/{id}", controllers.UpdatePersonnel)
	mux.HandleFunc("DELETE /personnel/{id}", controllers.DeletePersonnel)
	
	// Route spécifique pour filtrer par établissement
	mux.HandleFunc("GET /etablissements/{id}/personnel", controllers.GetByEtablissementPersonnel)
}