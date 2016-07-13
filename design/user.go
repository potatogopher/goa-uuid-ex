package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// User is the user resource media type
var User = MediaType("application/vnd.user+json", func() {
	Description("A User")
	Attributes(func() {
		Attribute("id", UUID, "ID of user", func() {
			Example("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
		})
		Attribute("href", String, "API href of the user", func() {
			Example("/users/6ba7b810-9dad-11d1-80b4-00c04fd430c8")
		})
		Attribute("email", String, "email of user", func() {
			Format("email")
			Example("foo@example")
		})
		Attribute("created_at", DateTime, "Creation time of the user")
		Attribute("updated_at", DateTime, "Last time user was updated")
		Attribute("deleted_at", DateTime, "When user was deleted")

		Required("id", "href", "email", "created_at", "updated_at")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("email")
		Attribute("created_at")
		Attribute("updated_at")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})

var _ = Resource("user", func() {
	DefaultMedia(User)
	BasePath("/users")

	Action("show", func() {
		Description("Retrieve user with given id.")

		Routing(
			GET("/:userID"),
		)
		Params(func() {
			Param("userID", UUID, "User ID")
		})
		Response(OK, User)
		Response(NotFound)
	})

	Action("create", func() {
		Description("Create new user.")

		Routing(
			POST(""),
		)
		Payload(func() {
			Member("email")
			Member("password")
			Required("email")
			Required("password")
		})
		Response(Created, User, "/users/[0-9-a-f]+")
		Response(InternalServerError)
	})

	Action("update", func() {
		Description("Update users's attributes for given id.")

		Routing(
			PUT("/:userID"),
		)
		Params(func() {
			Param("userID", UUID, "User ID")
		})
		Payload(func() {
			Member("email")
			Member("old_password")
			Member("new_password")
		})
		Response(OK, User)
		Response(NotFound)
		Response(InternalServerError)
		Response(BadRequest)
	})

	Action("delete", func() {
		Description("Delete user with given id.")

		Routing(
			DELETE("/:userID"),
		)
		Params(func() {
			Param("userID", UUID, "User ID")
		})
		Response(NoContent)
		Response(NotFound)
	})
})
