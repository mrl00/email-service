package models

import (
	"log"

	"gorm.io/gorm"
)

type Email struct {
	EmailId uint   `gorm:"primaryKey;autoIncrement:true"`
	UserId  uint   `gorm:"primaryKey;autoIncrement:false"`
	Email   string `gorm:"unique;not null"`
}

func CreateEmail(db *gorm.DB, email *Email) {
	result := db.Select("UserId", "Email").Create(email)
	if result.Error != nil {
		log.Fatalf("Failed to insert email into db: %v", result.Error)
	}
}

func FindEmailById(db *gorm.DB, email *Email) {
	result := db.Find(&email)
	if result.Error != nil {
		log.Fatalf("Cannot find emails with id (%v, %v). Error: %v", email.EmailId, email.UserId, result.Error)
	}
}

func FindAllEmails(db *gorm.DB) ([]Email, error) {
	var email []Email
	result := db.Find(&email)
	if result.Error != nil {
		log.Fatalf("Cannot find emails. Error: %v", result.Error)
	}
	return email, nil
}

func DeleteEmail(db *gorm.DB, email *Email) {
	result := db.Delete(email)
	if result.Error != nil {
		log.Fatalf("Cannot delete email with id (%v, %v). Error: %v", email.EmailId, email.UserId, result.Error)
	}
}
