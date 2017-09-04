package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB is a database connection.
var DB *gorm.DB

// Connect initializes connection to the database.
func Connect(host string, port int, user string, password string, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)

	var err error
	DB, err = gorm.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	}
}
