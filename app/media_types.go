//************************************************************************//
// API "UUID-Example": Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/potatogopher/goa-uuid-ex/design
// --notest=true
// --out=$(GOPATH)/src/github.com/potatogopher/goa-uuid-ex
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	uuid "github.com/satori/go.uuid"
	"time"
)

// User media type.
//
// Identifier: application/vnd.user+json
type User struct {
	// Creation time of the user
	CreatedAt time.Time `json:"created_at" xml:"created_at" form:"created_at"`
	// email of user
	Email string `json:"email" xml:"email" form:"email"`
	// API href of the user
	Href string `json:"href" xml:"href" form:"href"`
	// ID of user
	ID uuid.UUID `json:"id" xml:"id" form:"id"`
	// Last time user was updated
	UpdatedAt time.Time `json:"updated_at" xml:"updated_at" form:"updated_at"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}

	if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2))
	}
	return
}

// UserLink media type.
//
// Identifier: application/vnd.user+json
type UserLink struct {
	// API href of the user
	Href string `json:"href" xml:"href" form:"href"`
	// ID of user
	ID uuid.UUID `json:"id" xml:"id" form:"id"`
}

// Validate validates the UserLink media type instance.
func (mt *UserLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}

	return
}
