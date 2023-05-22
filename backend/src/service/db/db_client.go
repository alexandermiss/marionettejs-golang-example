package db

import (
	"github.com/sirupsen/logrus"
    "linguanski/src/config"
    "linguanski/src/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type dbClient struct {
	Db *gorm.DB
}

func (p dbClient) GetConnection() *gorm.DB {
    return p.Db
}

func NewDBClientFromConfig(config config.Config) DBClient {
	return NewDBClientFromOptions(config)
}

func NewDBClientFromOptions(config config.Config) DBClient {
    var logger *logrus.Logger
    logger = config.Logger
	logger.Debug("Trying to connect ")
    dns := config.DbConfig.ServerUrl
	db, _ := gorm.Open(postgres.Open(dns))
	db.AutoMigrate(&models.Client{})
	return dbClient{
		Db: db,
	}
}

func NewDBClient(db *gorm.DB) DBClient {
	return dbClient{
		Db: db,
	}
}

