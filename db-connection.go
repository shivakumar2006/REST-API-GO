package main

import "gorm.io/gorm"

var Database *gorm.DB
var urlDSN = "root:admin@tcp(127.0.0.1:3306)/newdb"

func DataMigration() {

}