package models

import (
	"log"

	"gorm.io/gorm"
)

type EmailModel struct {
	Db *gorm.DB
}

type Email struct {
	EmailId uint   `gorm:"primaryKey;autoIncrement:true"`
	UserId  uint   `gorm:"primaryKey;autoIncrement:false"`
	Email   string `gorm:"unique;not null"`
}

func (e *EmailModel) CreateEmail(email *Email) {
	result := e.Db.Select("UserId", "Email").Create(email)
	if result.Error != nil {
		log.Fatalf("Failed to insert email into db: %v", result.Error)
	}
}

func (e *EmailModel) FindEmailById(email *Email) {
	result := e.Db.Find(&email)
	if result.Error != nil {
		log.Fatalf("Cannot find emails with id (%v, %v). Error: %v", email.EmailId, email.UserId, result.Error)
	}
}

func (e *EmailModel) FindAllEmails() ([]Email, error) {
	var email []Email
	result := e.Db.Find(&email)
	if result.Error != nil {
		log.Fatalf("Cannot find emails. Error: %v", result.Error)
	}
	return email, nil
}

func (e *EmailModel) DeleteEmail(email *Email) {
	result := e.Db.Delete(email)
	if result.Error != nil {
		log.Fatalf("Cannot delete email with id (%v, %v). Error: %v", email.EmailId, email.UserId, result.Error)
	}
}
