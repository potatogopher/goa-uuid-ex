//************************************************************************//
// API "UUID-Example": Model Helpers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/potatogopher/goa-uuid-ex/design
// --out=$(GOPATH)/src/github.com/potatogopher/goa-uuid-ex
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/potatogopher/goa-uuid-ex/app"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"time"
)

// MediaType Retrieval Functions

// ListUser returns an array of view: default.
func (m *UserDB) ListUser(ctx context.Context) []*app.User {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuser"}, time.Now())

	var native []*User
	var objs []*app.User
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToUser())
	}

	return objs
}

// UserToUser loads a User and builds the default view of media type User.
func (m *User) UserToUser() *app.User {
	user := &app.User{}
	user.CreatedAt = m.CreatedAt
	user.Email = m.Email
	user.ID = m.ID
	user.UpdatedAt = m.UpdatedAt

	return user
}

// OneUser loads a User and builds the default view of media type User.
func (m *UserDB) OneUser(ctx context.Context, id uuid.UUID) (*app.User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuser"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToUser()
	return &view, err
}

// MediaType Retrieval Functions

// ListUserLink returns an array of view: link.
func (m *UserDB) ListUserLink(ctx context.Context) []*app.UserLink {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuserlink"}, time.Now())

	var native []*User
	var objs []*app.UserLink
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToUserLink())
	}

	return objs
}

// UserToUserLink loads a User and builds the link view of media type User.
func (m *User) UserToUserLink() *app.UserLink {
	user := &app.UserLink{}
	user.ID = m.ID

	return user
}

// OneUserLink loads a User and builds the link view of media type User.
func (m *UserDB) OneUserLink(ctx context.Context, id uuid.UUID) (*app.UserLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuserlink"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToUserLink()
	return &view, err
}
