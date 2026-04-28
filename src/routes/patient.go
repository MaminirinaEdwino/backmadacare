package routes

import (
	"net/http"

	"github.com/MaminirinaEdwino/backmadacare/src/controllers"
)

func RegisterRoutesPatient(mux *http.ServeMux) {
	mux.HandleFunc("GET /patients", controllers.GetAllPatient)
	mux.HandleFunc("GET /patients/{id}", controllers.GetOnePatient)
	mux.HandleFunc("POST /patients", controllers.CreatePatient)
	mux.HandleFunc("PUT /patients/{id}", controllers.UpdatePatient)
	mux.HandleFunc("DELETE /patients/{id}", controllers.DeletePatient)
	
	// Route pour libérer un patient (date de sortie)
	mux.HandleFunc("PUT /patients/{id}/sortie", controllers.MarkDischarged)
	mux.HandleFunc("GET /patients/en-attente", controllers.GetPendingPatient)
}