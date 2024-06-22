package models_test

import (
	"testing"

	"github.com/mrl00/email-service/src/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func setupDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "tb_",
		},
	})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Email{})
	assert.NoError(t, err)

	return db
}

func TestCreateEmail(t *testing.T) {
	db := setupDB(t)

	email1 := models.Email{
		UserId: 2,
		Email:  "email1@mail.com",
	}
	models.CreateEmail(db, &email1)
	expectedEmail1 := models.Email{
		EmailId: 1,
		UserId:  2,
		Email:   "email1@mail.com",
	}
	assert.Equal(t, expectedEmail1, email1, "email1 not equals")

	email2 := models.Email{
		UserId: 2,
		Email:  "email2@mail.com",
	}
	models.CreateEmail(db, &email2)
	expectedEmail2 := models.Email{
		EmailId: 2,
		UserId:  2,
		Email:   "email2@mail.com",
	}
	assert.Equal(t, expectedEmail2, email2, "email2 not equals")
}

func TestFindEmailById(t *testing.T) {
	db := setupDB(t)

	addedEmail := models.Email{
		UserId: 2,
		Email:  "email1@mail.com",
	}
	findEmail := models.Email{
		EmailId: 1,
		UserId:  2,
	}
	expectedEmail1 := models.Email{
		EmailId: 1,
		UserId:  2,
		Email:   "email1@mail.com",
	}

	models.CreateEmail(db, &addedEmail)
	models.FindEmailById(db, &findEmail)

	assert.Equal(t, expectedEmail1, findEmail, "emails are not equals")
}

func TestFindAllEmails(t *testing.T) {
	db := setupDB(t)

	email1 := models.Email{
		UserId: 2,
		Email:  "email1@mail.com",
	}
	email2 := models.Email{
		UserId: 2,
		Email:  "email2@mail.com",
	}

	expectedEmails := []models.Email{
		{
			EmailId: 1,
			UserId:  2,
			Email:   "email1@mail.com",
		},
		{
			EmailId: 2,
			UserId:  2,
			Email:   "email2@mail.com",
		},
	}

	models.CreateEmail(db, &email1)
	models.CreateEmail(db, &email2)

	foundEmails, _ := models.FindAllEmails(db)

	assert.Equal(t, expectedEmails, foundEmails, "list of emails are no equals")
}
