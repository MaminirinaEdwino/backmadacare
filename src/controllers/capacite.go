package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MaminirinaEdwino/backmadacare/src/config"
	"github.com/MaminirinaEdwino/backmadacare/src/models"
)

func CreateCapacite(w http.ResponseWriter, r *http.Request) {
	var c models.Capacite
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&c).Error; err != nil {
		http.Error(w, "Erreur lors de la création de la capacité", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func GetAllCapacite(w http.ResponseWriter, r *http.Request) {
	var capacites []models.Capacite
	// On preload l'établissement pour la clarté
	config.DB.Preload("Etablissement").Find(&capacites)
	json.NewEncoder(w).Encode(capacites)
}

func GetOneCapacite(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var c models.Capacite
	if err := config.DB.Preload("Etablissement").First(&c, id).Error; err != nil {
		http.Error(w, "Capacité non trouvée", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(c)
}

func GetByEtablissementCapacite(w http.ResponseWriter, r *http.Request) {
	etabID := r.PathValue("id")
	var capacites []models.Capacite
	
	config.DB.Where("etablissement_id = ?", etabID).Find(&capacites)
	json.NewEncoder(w).Encode(capacites)
}

func UpdateCapacite(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var c models.Capacite
	if err := config.DB.First(&c, id).Error; err != nil {
		http.Error(w, "Capacité non trouvée", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&c)
	config.DB.Save(&c)
	json.NewEncoder(w).Encode(c)
}

func DeleteCapacite(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	config.DB.Delete(&models.Capacite{}, id)
	w.WriteHeader(http.StatusNoContent)
}