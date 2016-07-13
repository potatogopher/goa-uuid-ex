package controllers

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"

	"github.com/potatogopher/goa-uuid-ex/app"
	"github.com/potatogopher/goa-uuid-ex/models"
)

var (
	// ErrDatabaseError is an error we can't handle ourselves.
	ErrDatabaseError = goa.NewErrorClass("db_error", 500)

	userDB *models.UserDB
)

// Mount mounts all of the necessary controllers for a service.
func Mount(service *goa.Service) {
	app.MountUserController(service, NewUserController(service))
}

// SetupDatabase configures and connects to the database.
func SetupDatabase(db *gorm.DB, logMode bool) {
	db.LogMode(logMode)

	userDB = models.NewUserDB(*db)
}
