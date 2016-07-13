package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("WheelStorageGroup", func() {
	Description("This is the global storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the Postgres relational store")

		Model("User", func() {
			RendersTo(User)
			Description("Wheel User")
			Field("id", gorma.UUID, func() {
				PrimaryKey()
			})
			Field("email", gorma.String)
			Field("password_hash", gorma.String)
		})

	})
})
