package appconfig

import (
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Config struct {
	Application AppConfig `mapstruct:"application"`
	Database    DBConfig  `mapstruct:"database"`
}

type AppConfig struct {
	Name string `mapstruct:"name"`
	Host string `mapstruct:"host"`
	Port string `mapstruct:"port"`
}

type DBConfig struct {
	Host     string `mapstruct:"host"`
	User     string `mapstruct:"user"`
	Password string `mapstruct:"password"`
	Name     string `mapstruct:"name"`
	Port     string `mapstruct:"port"`
	Ssl      string `mapstruct:"ssl"`
	Tz       string `mapstruct:"tz"`
}

func (d DBConfig) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("host", d.Host),
		slog.String("user", d.User),
		slog.String("name", d.Name),
		slog.String("port", d.Port),
		slog.String("ssl", d.Ssl),
		slog.String("tz", d.Tz),
	)
}

func (c DBConfig) Dsn() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		c.Host, c.User, c.Password, c.Name, c.Port, c.Ssl, c.Tz)
}

func GetDatabaseConfig(dbConfig DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbConfig.Dsn()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "tb_",
		},
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db, nil
}

func GetAppConfiguration() (*Config, error) {
	app := Config{}

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	var err = config.LoadFiles("./config/config.yaml")
	if err != nil {
		return nil, err
	}

	err = config.BindStruct("", &app)
	if err != nil {
		log.Fatalf("failed to bind config file to struct. Err: %v", err)
	}

	return &app, nil
}
