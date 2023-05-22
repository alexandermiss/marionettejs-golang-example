package db

import (
    "gorm.io/gorm"
)

type DBClient interface {
    GetConnection() *gorm.DB
}
