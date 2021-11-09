package model

import (
	"github.com/jinzhu/gorm"
	"lucky/db_server"
)

var db *gorm.DB = db_server.MySqlDb
