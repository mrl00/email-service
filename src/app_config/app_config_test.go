package appconfig_test

import (
	"testing"

	appconfig "github.com/mrl00/email-service/src/app_config"
)

func TestDsn(t *testing.T) {
	dbConfig := appconfig.DBConfig{
		Host:     "localhost",
		User:     "postgres",
		Password: "123465",
		Name:     "email_db",
		Port:     "5432",
		Ssl:      "disable",
		Tz:       "America/Sao_Paulo",
	}

	dsn := "host=localhost user=postgres password=123465 dbname=email_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	if dbConfig.Dsn() != dsn {
		t.Error("invalid dns")
	}
}
