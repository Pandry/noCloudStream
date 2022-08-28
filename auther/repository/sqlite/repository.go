package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"auther/models"
)

type Database struct {
	l  models.Logger
	db *gorm.DB
}

func InitDatabase(l models.Logger, dbName string) (*Database, error) {
	d := &Database{l: l}
	var err error
	d.db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	d.db.AutoMigrate(&models.User{})
	return d, err
}

func (d *Database) Close() {

}
