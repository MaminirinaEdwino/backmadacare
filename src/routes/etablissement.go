package routes

import (
	"net/http"

	"github.com/MaminirinaEdwino/backmadacare/src/controllers"
)

func  RegisterRoutesEtablissement(mux *http.ServeMux) {
	mux.HandleFunc("GET /etablissements", controllers.GetAllEtablissement)
	mux.HandleFunc("GET /etablissements/{id}", controllers.GetOneEtablissement)
	mux.HandleFunc("POST /etablissements", controllers.CreateEtablissement)
	mux.HandleFunc("PUT /etablissements/{id}", controllers.UpdateEtablissement)
	mux.HandleFunc("DELETE /etablissements/{id}", controllers.DeleteEtablissement)
}