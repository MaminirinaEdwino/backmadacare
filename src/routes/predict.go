package routes

import (
	"net/http"

	"github.com/MaminirinaEdwino/backmadacare/src/controllers"
)

func PredictRegisterRoutes(mux *http.ServeMux) {
    mux.HandleFunc("POST /predict", controllers.Predicthandler)
}