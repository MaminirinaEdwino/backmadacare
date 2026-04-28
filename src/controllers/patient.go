package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MaminirinaEdwino/backmadacare/src/config"
	"github.com/MaminirinaEdwino/backmadacare/src/models"
)

func CreatePatient(w http.ResponseWriter, r *http.Request) {
	var p models.Patient
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	// Par défaut, on met la date d'admission à "maintenant" si non fournie
	if p.DateAdmission.IsZero() {
		p.DateAdmission = time.Now()
	}

	if err := config.DB.Create(&p).Error; err != nil {
		http.Error(w, "Erreur lors de l'admission du patient", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func GetAllPatient(w http.ResponseWriter, r *http.Request) {
	var patients []models.Patient
	// On affiche l'établissement pour chaque patient
	config.DB.Preload("Etablissement").Find(&patients)
	json.NewEncoder(w).Encode(patients)
}

func GetOnePatient(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var p models.Patient
	if err := config.DB.Preload("Etablissement").First(&p, id).Error; err != nil {
		http.Error(w, "Patient non trouvé", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var p models.Patient
	if err := config.DB.First(&p, id).Error; err != nil {
		http.Error(w, "Patient non trouvé", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&p)
	config.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}

func MarkDischarged(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var p models.Patient
	if err := config.DB.First(&p, id).Error; err != nil {
		http.Error(w, "Patient non trouvé", http.StatusNotFound)
		return
	}

	now := time.Now()
	p.DateSortie = &now
	p.Status = "Sorti"
	
	config.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}

func DeletePatient(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	config.DB.Delete(&models.Patient{}, id)
	w.WriteHeader(http.StatusNoContent)
}

func GetPendingPatient(w http.ResponseWriter, r *http.Request) {
	var patients []models.Patient
	
	// On filtre par le statut exact et on charge l'établissement lié
	result := config.DB.Preload("Etablissement").Where("status = ?", "en_attente").Find(&patients)
	
	if result.Error != nil {
		http.Error(w, "Erreur lors de la récupération des patients", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patients)
}