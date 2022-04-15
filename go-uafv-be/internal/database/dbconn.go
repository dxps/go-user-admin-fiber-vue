package database

import (
	"fmt"

	"dxps.io/go-user-admin-fiber-vue/be/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (db *gorm.DB, deferDbCloseFn func() error, err error) {

	dsn := "host=localhost port=5438 sslmode=disable user=go-uafv password=go-uafv dbname=go-uafv"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	deferDbCloseFn = func() error {
		idb, err := db.DB()
		if err != nil {
			return err
		} else {
			idb.Close()
			return nil
		}
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		err = fmt.Errorf("Database automigrate err: %v", err)
	}

	return
}
