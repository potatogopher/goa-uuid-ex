//************************************************************************//
// User Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/potatogopher/goa-uuid-ex/design
// --notool=true
// --out=$(GOPATH)/src/github.com/potatogopher/goa-uuid-ex
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

// A User
type User struct {
	// Creation time of the user
	CreatedAt time.Time `json:"created_at" xml:"created_at" form:"created_at"`
	// When user was deleted
	DeletedAt *time.Time `json:"deleted_at,omitempty" xml:"deleted_at,omitempty" form:"deleted_at,omitempty"`
	// email of user
	Email string `json:"email" xml:"email" form:"email"`
	// API href of the user
	Href string `json:"href" xml:"href" form:"href"`
	// ID of user
	ID uuid.UUID `json:"id" xml:"id" form:"id"`
	// Last time user was updated
	UpdatedAt time.Time `json:"updated_at" xml:"updated_at" form:"updated_at"`
}

// DecodeUser decodes the User instance encoded in resp body.
func (c *Client) DecodeUser(resp *http.Response) (*User, error) {
	var decoded User
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
