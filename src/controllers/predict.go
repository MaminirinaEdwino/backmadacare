package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MaminirinaEdwino/backmadacare/src/config"
	"github.com/MaminirinaEdwino/backmadacare/src/models"
	"github.com/MaminirinaEdwino/gobayes"
)

var Network *gobayes.Network

func FindPossibleMaladie(prediction map[string]float64) string {
	var prob float64
	prob = 0
	maladie := ""
	for key, value := range prediction {
		if value > float64(prob) {
			prob = value
			maladie = key
		}
	}
	return maladie
}

func Predicthandler(w http.ResponseWriter, r *http.Request) {
	dataUrgence := config.SyncMedicalUrgenceRules("src/config/data/urgence.json")
	var req models.RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requête JSON invalide", http.StatusBadRequest)
		return
	}

	resultFactor := Network.Query("Maladie", req.Evidence)

	targetNode, exists := Network.Nodes["Maladie"]
	if !exists {
		http.Error(w, "Nœud cible introuvable dans le réseau", http.StatusNotFound)
		return
	}
	predictions := make(map[string]float64)
	for i, stateName := range targetNode.States {
		predictions[stateName] = resultFactor.Values[i]
	}
	maladie := FindPossibleMaladie(predictions)
	resp := models.ResponseBody{
		Maladie: maladie,
	}

	

    // Utilisation de GORM pour filtrer
    result := config.DB.Where("region = ?", req.Region).Find(&resp.Etablissement)
    
    if result.Error != nil {
        http.Error(w, "Erreur lors de la recherche", http.StatusInternalServerError)
        return
    }

	resp.Urgence = dataUrgence.Rules[maladie]
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("ERREUR ENCODAGE JSON: %v", err)
	}
}
