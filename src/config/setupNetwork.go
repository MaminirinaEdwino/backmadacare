package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/MaminirinaEdwino/gobayes"
)

func SetupMedicalNetwork(net *gobayes.Network) {
	if net == nil {
		log.Fatal("Le réseau est nil !")
	}

	// --- ÉTAGE 1 : LES OBSERVATIONS (Parents) ---
	// On définit les symptômes possibles
	net.AddNode("fievre", []string{"Non", "Oui"})
	net.AddNode("toux", []string{"Non", "Oui"})
	net.AddNode("toux_seche", []string{"Non", "Oui"})
	net.AddNode("perte_odorat", []string{"Non", "Oui"})
	net.AddNode("difficulte_respiratoire", []string{"Non", "Oui"})
	net.AddNode("convulsion", []string{"Non", "Oui"})
	net.AddNode("diarrhee_severe", []string{"Non", "Oui"})
	net.AddNode("paralysie", []string{"Non", "Oui"})
	net.AddNode("douleur_thoracique", []string{"Non", "Oui"})
	net.AddNode("raideur_nuque", []string{"Non", "Oui"})

	// Conditions environnementales ou démographiques
	net.AddNode("zone_endemique", []string{"Non", "Oui"})
	net.AddNode("enfant", []string{"Non", "Oui"})

	// --- ÉTAGE 2 : LE DIAGNOSTIC (Enfant) ---
	// Liste des maladies possibles (les états du nœud Maladie)
	net.AddNode("Maladie", []string{
		"Aucune", "Paludisme", "Paludisme_grave", "Grippe",
		"Covid-19", "Pneumonie", "Asthme", "Cholera",
		"AVC", "Infarctus", "Meningite",
	})

	// --- ÉTAGE 3 : LES LIENS (Edges) ---
	// On lie les symptômes vers la maladie
	net.AddEdge("fievre", "Maladie")
	net.AddEdge("zone_endemique", "Maladie")
	net.AddEdge("enfant", "Maladie")
	net.AddEdge("toux", "Maladie")
	net.AddEdge("perte_odorat", "Maladie")
	net.AddEdge("difficulte_respiratoire", "Maladie")
	net.AddEdge("convulsion", "Maladie")
	net.AddEdge("diarrhee_severe", "Maladie")
	net.AddEdge("paralysie", "Maladie")
	net.AddEdge("douleur_thoracique", "Maladie")
	net.AddEdge("raideur_nuque", "Maladie")
}

func SyncMedicalRules(net *gobayes.Network, rulesPath string) error {
	file, err := os.Open(rulesPath)
	if err != nil {
		return err
	}
	defer file.Close()

	var data struct {
		Rules []gobayes.ScoreRule `json:"maladie_rules"`
	}

	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return err
	}

	// On applique les probabilités au nœud cible "Maladie"
	maladieNode := net.Nodes["Maladie"]
	if maladieNode != nil {
		maladieNode.GenerateCPD(data.Rules)
		log.Printf("Table de Probabilité Conditionnelle (CPD) générée pour 'Maladie'")

		if len(maladieNode.CPD) == 0 {
			log.Fatal("ERREUR : La CPD est vide. Vérifie les labels dans rules.json")
		}
	}

	return nil
}

type dataUrgence struct {
	Rules map[string]string `json:"urgences"`
}

func SyncMedicalUrgenceRules(rulesPath string) dataUrgence {
	file, err := os.Open(rulesPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data dataUrgence
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		log.Fatal(err)
	}
	return data
}
