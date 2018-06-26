package models

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection
var GormDB *gorm.DB

func init() {
	var err error
	env := envy.Get("GO_ENV", "development")
	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}

	deets := DB.Dialect.Details()
	GormDB, err = gorm.Open(deets.Dialect, DB.URL())
	if err != nil {
		log.Fatal(err)
	}
	GormDB = GormDB.LogMode(true)
}
