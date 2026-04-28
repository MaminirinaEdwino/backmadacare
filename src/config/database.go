package config

import (
	"fmt"
	"log"

	"github.com/MaminirinaEdwino/backmadacare/src/models"
	"gorm.io/driver/postgres" // Ou mysql
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Charge normalement ces infos depuis un fichier .env
	dsn := "host=localhost user=postgres password=root dbname=madacare port=5432 sslmode=disable"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données :", err)
	}

	fmt.Println("Connexion à la base de données réussie !")

	// Migration automatique des tables (AutoMigrate crée les tables si elles n'existent pas)
	err = db.AutoMigrate(
		&models.Etablissement{},
		&models.Admin{},
		&models.Personnel{},
		&models.Patient{},
		&models.Ambulance{},
		&models.Capacite{},
	)
	
	if err != nil {
		log.Fatal("Erreur lors de la migration :", err)
	}

	DB = db
}