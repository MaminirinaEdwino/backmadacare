package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MaminirinaEdwino/backmadacare/src/config"
	"github.com/MaminirinaEdwino/backmadacare/src/models"
	"golang.org/x/crypto/bcrypt" // Pense à faire : go get golang.org/x/crypto/bcrypt
)



// Create : Inscription avec hachage du mot de passe
func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var admin models.Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}

	// Hachage du mot de passe avant sauvegarde
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Mdp), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erreur lors du traitement du mot de passe", http.StatusInternalServerError)
		return
	}
	admin.Mdp = string(hashedPassword)

	if err := config.DB.Create(&admin).Error; err != nil {
		http.Error(w, "Erreur lors de la création (Username ou Email peut-être déjà pris)", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(admin)
}

func GetAllAdmin(w http.ResponseWriter, r *http.Request) {
	var admins []models.Admin
	// On utilise Preload pour charger aussi les infos de l'établissement lié
	config.DB.Preload("Etablissement").Find(&admins)
	json.NewEncoder(w).Encode(admins)
}

func GetOneAdmin(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var admin models.Admin
	if err := config.DB.Preload("Etablissement").First(&admin, id).Error; err != nil {
		http.Error(w, "Administrateur non trouvé", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(admin)
}

func UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var admin models.Admin
	if err := config.DB.First(&admin, id).Error; err != nil {
		http.Error(w, "Admin non trouvé", http.StatusNotFound)
		return
	}

	var updateData models.Admin
	json.NewDecoder(r.Body).Decode(&updateData)

	// Si le mot de passe est modifié, on le re-hache
	if updateData.Mdp != "" {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(updateData.Mdp), bcrypt.DefaultCost)
		updateData.Mdp = string(hashed)
	}

	config.DB.Model(&admin).Updates(updateData)
	json.NewEncoder(w).Encode(admin)
}

func DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	config.DB.Delete(&models.Admin{}, id)
	w.WriteHeader(http.StatusNoContent)
}