package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MaminirinaEdwino/backmadacare/src/models"
	"github.com/MaminirinaEdwino/gobayes"
)

var Network *gobayes.Network

func Predicthandler(w http.ResponseWriter, r *http.Request) {
	var req models.RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requête JSON invalide", http.StatusBadRequest)
		return
	}

	resultFactor := Network.Query(req.Target, req.Evidence)

	targetNode, exists := Network.Nodes[req.Target]
	if !exists {
		http.Error(w, "Nœud cible introuvable dans le réseau", http.StatusNotFound)
		return
	}
	predictions := make(map[string]float64)
	for i, stateName := range targetNode.States {
		predictions[stateName] = resultFactor.Values[i]
	}
	resp := models.ResponseBody{
		Target:      req.Target,
		Predictions: predictions,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("ERREUR ENCODAGE JSON: %v", err)
	}
}
