package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MaminirinaEdwino/backmadacare/src/config"
	"github.com/MaminirinaEdwino/backmadacare/src/models"
)




func CreatePersonnel(w http.ResponseWriter, r *http.Request) {
	var p models.Personnel
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&p).Error; err != nil {
		http.Error(w, "Erreur lors de l'ajout du personnel", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func GetAllPersonnel(w http.ResponseWriter, r *http.Request) {
	var staff []models.Personnel
	// On charge l'établissement pour savoir où travaille chaque membre
	config.DB.Preload("Etablissement").Find(&staff)
	json.NewEncoder(w).Encode(staff)
}

func GetOnePersonnel(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var p models.Personnel
	if err := config.DB.Preload("Etablissement").First(&p, id).Error; err != nil {
		http.Error(w, "Membre du personnel non trouvé", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func GetByEtablissementPersonnel(w http.ResponseWriter, r *http.Request) {
	etabID := r.PathValue("id")
	var staff []models.Personnel
	
	config.DB.Where("etablissement_id = ?", etabID).Find(&staff)
	json.NewEncoder(w).Encode(staff)
}

func UpdatePersonnel(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var p models.Personnel
	if err := config.DB.First(&p, id).Error; err != nil {
		http.Error(w, "Personnel non trouvé", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&p)
	config.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}

func DeletePersonnel(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	config.DB.Delete(&models.Personnel{}, id)
	w.WriteHeader(http.StatusNoContent)
}