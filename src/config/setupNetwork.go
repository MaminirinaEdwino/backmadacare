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
	net.AddNode("bokotra manaintaina ao amin'ny ventriny na ankihibe", []string{"Non", "Oui"})
	net.AddNode("tazo tampoka > 39°C", []string{"Non", "Oui"})
	net.AddNode("tazo telo andro na efatra andro", []string{"Non", "Oui"})
	net.AddNode("mangovitra mafy", []string{"Non", "Oui"})
	net.AddNode("fandroana be", []string{"Non", "Oui"})
	net.AddNode("rà ao amin'ny pivia (hématurie terminale)", []string{"Non", "Oui"})
	net.AddNode("fanaintainana rehefa misitrana", []string{"Non", "Oui"})
	net.AddNode("tsindoka (épilepsie)", []string{"Non", "Oui"})
	net.AddNode("aretin-doha maharitra", []string{"Non", "Oui"})
	net.AddNode("fihazonana ao anatiny ny atidoha (hypertension intracrânienne)", []string{"Non", "Oui"})
	net.AddNode("tazo maharitra 39-40°C", []string{"Non", "Oui"})
	net.AddNode("pouls tsy mirindra amin'ny tazo (pouls dissocié)", []string{"Non", "Oui"})
	net.AddNode("tuphos (fahasahiranana saina sy fadiranovana tanteraka)", []string{"Non", "Oui"})

	// --- ÉTAGE 2 : LE DIAGNOSTIC (Enfant) ---
	// Liste des maladies possibles (les états du nœud Maladie)
	net.AddNode("Maladie", []string{
		"Pesta bubonique (Pesta)", "Tazomoka (Paludisme/Malaria)", "Bilharziose (Schistosomiase)", "Neurocysticercose", "Tazo Tifoida (Fièvre Typhoïde)",
	})

	// --- ÉTAGE 3 : LES LIENS (Edges) ---
	// On lie les symptômes vers la maladie
	net.AddEdge("bokotra manaintaina ao amin'ny ventriny na ankihibe", "Maladie")
	net.AddEdge("tazo tampoka > 39°C", "Maladie")
	net.AddEdge("tazo telo andro na efatra andro", "Maladie")
	net.AddEdge("mangovitra mafy", "Maladie")
	net.AddEdge("fandroana be", "Maladie")
	net.AddEdge("rà ao amin'ny pivia (hématurie terminale)", "Maladie")
	net.AddEdge("fanaintainana rehefa misitrana", "Maladie")
	net.AddEdge("tsindoka (épilepsie)", "Maladie")
	net.AddEdge("aretin-doha maharitra", "Maladie")
	net.AddEdge("fihazonana ao anatiny ny atidoha (hypertension intracrânienne)", "Maladie")
	net.AddEdge("tazo maharitra 39-40°C", "Maladie")
	net.AddEdge("pouls tsy mirindra amin'ny tazo (pouls dissocié)", "Maladie")
	net.AddEdge("tuphos (fahasahiranana saina sy fadiranovana tanteraka)", "Maladie")

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
