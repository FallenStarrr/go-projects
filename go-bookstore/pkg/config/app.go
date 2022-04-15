package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  )

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=10.4.110.3 user=postuser password=Jv#/TWw3 dbname=file_service port=5432 sslmode=disable"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}