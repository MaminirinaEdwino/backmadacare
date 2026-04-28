package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MaminirinaEdwino/backmadacare/src/config"
	"github.com/MaminirinaEdwino/backmadacare/src/models"
)



// Routes expose les points d'entrée CRUD


func  GetAllEtablissement(w http.ResponseWriter, r *http.Request) {
	var items []models.Etablissement
	config.DB.Find(&items)
	json.NewEncoder(w).Encode(items)
}

func  GetOneEtablissement(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var item models.Etablissement
	if err := config.DB.First(&item, id).Error; err != nil {
		http.Error(w, "Établissement non trouvé", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func CreateEtablissement(w http.ResponseWriter, r *http.Request) {
	var item models.Etablissement
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	config.DB.Create(&item)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func UpdateEtablissement(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var item models.Etablissement
	if err := config.DB.First(&item, id).Error; err != nil {
		http.Error(w, "Non trouvé", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&item)
	config.DB.Save(&item)
	json.NewEncoder(w).Encode(item)
}

func DeleteEtablissement(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	config.DB.Delete(&models.Etablissement{}, id)
	w.WriteHeader(http.StatusNoContent)
}