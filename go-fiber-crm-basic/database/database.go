package database

import (
	"gihub.com/jinzhu/gorm"
	_ "gihub.com/jinzhu/gorm/dialects/sqllite"
)

var (
	DBConn *gorm.DB
)