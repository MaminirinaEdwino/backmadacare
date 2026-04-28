package routes

import (
	"net/http"

	"github.com/MaminirinaEdwino/backmadacare/src/controllers"
)

func RegisterRoutesAdmin(mux *http.ServeMux) {
	mux.HandleFunc("GET /admins", controllers.GetAllAdmin)
	mux.HandleFunc("GET /admins/{id}", controllers.GetOneAdmin)
	mux.HandleFunc("POST /admins", controllers.CreateAdmin) // Inscription d'un admin
	mux.HandleFunc("PUT /admins/{id}", controllers.UpdateAdmin)
	mux.HandleFunc("DELETE /admins/{id}", controllers.DeleteAdmin)
}
