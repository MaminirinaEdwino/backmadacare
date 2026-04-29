package controllers

import (
	"encoding/json"
	"fmt"
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
	listeSympt := []string{
		"bokotra manaintaina ao amin'ny ventriny na ankihibe",
		"tazo tampoka \u003e 39°C",
		"fandroana be",
		"mangovitra mafy",
		"tazo telo andro na efatra andro",
		"fanaintainana rehefa misitrana",
		"rà ao amin'ny pivia (hématurie terminale)",
		"aretin-doha maharitra",
		"fihazonana ao anatiny ny atidoha (hypertension intracrânienne)",
		"tsindoka (épilepsie)",
		"pouls tsy mirindra amin'ny tazo (pouls dissocié)",
		"tazo maharitra 39-40°C", 
		"tuphos (fahasahiranana saina sy fadiranovana tanteraka)",
	}
	dataUrgence := config.SyncMedicalUrgenceRules("src/config/data/urgence2.json")
	var req models.RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requête JSON invalide", http.StatusBadRequest)
		return
	}
	var Evidence map[string]int = map[string]int{}

	for _, val := range listeSympt {
		_, exist := req.Evidence[val]
		if exist {
			Evidence[val] = 1
		}else{
			Evidence[val] = 0
		}
	}

	resultFactor := Network.Query("Maladie", Evidence)

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
	fmt.Println(predictions)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("ERREUR ENCODAGE JSON: %v", err)
	}
}
