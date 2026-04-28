package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MaminirinaEdwino/backmadacare/src/config"
	"github.com/MaminirinaEdwino/backmadacare/src/models"
)



func RegisterRoutesAmbulance(mux *http.ServeMux) {
	mux.HandleFunc("GET /ambulances", GetAllAmbulance)
	mux.HandleFunc("GET /ambulances/{id}", GetOneAmbulance)
	mux.HandleFunc("POST /ambulances", CreateAmbulance)
	mux.HandleFunc("PUT /ambulances/{id}", UpdateAmbulance)
	mux.HandleFunc("DELETE /ambulances/{id}", DeleteAmbulance)

	// Endpoint spécifique pour la gestion opérationnelle
	mux.HandleFunc("GET /ambulances/disponibles", GetAvailableAmbulance)
}

func CreateAmbulance(w http.ResponseWriter, r *http.Request) {
	var a models.Ambulance
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&a).Error; err != nil {
		http.Error(w, "Erreur lors de l'enregistrement de l'ambulance", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(a)
}

func GetAllAmbulance(w http.ResponseWriter, r *http.Request) {
	var ambulances []models.Ambulance
	// On charge les informations du chauffeur (Personnel)
	config.DB.Preload("Chauffeur").Find(&ambulances)
	json.NewEncoder(w).Encode(ambulances)
}

func GetOneAmbulance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var a models.Ambulance
	if err := config.DB.Preload("Chauffeur").First(&a, id).Error; err != nil {
		http.Error(w, "Ambulance non trouvée", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(a)
}

func GetAvailableAmbulance(w http.ResponseWriter, r *http.Request) {
	var ambulances []models.Ambulance
	
	// On filtre par statut (ex: 'disponible' ou 'libre')
	err := config.DB.Preload("Chauffeur").Where("status = ?", "disponible").Find(&ambulances).Error
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ambulances)
}

func UpdateAmbulance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var a models.Ambulance
	if err := config.DB.First(&a, id).Error; err != nil {
		http.Error(w, "Ambulance non trouvée", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&a)
	config.DB.Save(&a)
	json.NewEncoder(w).Encode(a)
}

func DeleteAmbulance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	config.DB.Delete(&models.Ambulance{}, id)
	w.WriteHeader(http.StatusNoContent)
}