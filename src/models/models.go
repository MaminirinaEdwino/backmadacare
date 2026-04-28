package models

import (
	"time"

	"gorm.io/gorm"
)

type RequestBody struct {
	Evidence map[string]int `json:"evidence"`
	Region   string         `json:"region"`
}

type ResponseBody struct {
	Maladie       string          `json:"maladie"`
	Urgence       string          `json:"urgence"`
	Etablissement []Etablissement `json:"etablissement"`
}

type Etablissement struct {
	gorm.Model
	Nom       string `json:"nom"`
	Region    string `json:"region"`
	Contact   string `json:"contact"`
	Categorie string `json:"categorie"`
}

type Admin struct {
	gorm.Model
	Nom             string        `json:"nom"`
	Prenom          string        `json:"prenom"`
	Username        string        `gorm:"unique;not null" json:"username"`
	Mdp             string        `json:"mdp"` // "-" pour ne pas exposer le mot de passe en JSON
	Email           string        `gorm:"unique" json:"email"`
	EtablissementID uint          `json:"etablissement_id"`
	Etablissement   Etablissement `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
}

type Personnel struct {
	gorm.Model
	Nom             string        `json:"nom"`
	Prenom          string        `json:"prenom"`
	Contact         string        `json:"contact"`
	Poste           string        `json:"poste"`
	Age             int           `json:"age"`
	EtablissementID uint          `json:"etablissement_id"`
	Etablissement   Etablissement `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
}

type Patient struct {
	gorm.Model
	Nom             string        `json:"nom"`
	Prenom          string        `json:"prenom"`
	Maladies        string        `json:"maladies"` // Stocké en texte ou JSON
	EtablissementID uint          `json:"etablissement_id"`
	Etablissement   Etablissement `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	DateAdmission   time.Time     `json:"date_admission"`
	DateSortie      *time.Time    `json:"date_sortie"` // Pointeur pour gérer le NULL
	Status          string        `json:"status"`
}

type Ambulance struct {
	gorm.Model
	Refference  string    `gorm:"unique" json:"refference"`
	ChauffeurID uint      `json:"chauffeur_id"`
	Chauffeur   Personnel `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	Status      string    `json:"status"`
}

type Capacite struct {
	gorm.Model
	Maladies        string        `json:"maladies"` // Ex: "Paludisme", "COVID-19"
	Espaces         int           `json:"espaces"`  // Ex: "15 lits", "5 chambres"
	EtablissementID uint          `json:"etablissement_id"`
	Etablissement   Etablissement `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
}
