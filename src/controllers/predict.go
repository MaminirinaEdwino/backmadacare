package controllers

import (
	"fmt"
	"net/http"
)

func Predicthandler(w http.ResponseWriter, r *http.Request ){
	fmt.Println("predict")
}