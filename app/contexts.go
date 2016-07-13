//************************************************************************//
// API "UUID-Example": Application Contexts
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
	"golang.org/x/net/context"
)

// CreateUserContext provides the user create action context.
type CreateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	Payload *CreateUserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(ctx context.Context, service *goa.Service) (*CreateUserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := CreateUserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// createUserPayload is the user create action payload.
type createUserPayload struct {
	// email of user
	Email    *string `json:"email,omitempty" xml:"email,omitempty" form:"email,omitempty"`
	Password *string `json:"password,omitempty" xml:"password,omitempty" form:"password,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createUserPayload) Validate() (err error) {
	if payload.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}

	if payload.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *payload.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, *payload.Email, goa.FormatEmail, err2))
		}
	}
	return
}

// Publicize creates CreateUserPayload from createUserPayload
func (payload *createUserPayload) Publicize() *CreateUserPayload {
	var pub CreateUserPayload
	if payload.Email != nil {
		pub.Email = *payload.Email
	}
	if payload.Password != nil {
		pub.Password = *payload.Password
	}
	return &pub
}

// CreateUserPayload is the user create action payload.
type CreateUserPayload struct {
	// email of user
	Email    string `json:"email" xml:"email" form:"email"`
	Password string `json:"password" xml:"password" form:"password"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateUserPayload) Validate() (err error) {
	if payload.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}

	if err2 := goa.ValidateFormat(goa.FormatEmail, payload.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, payload.Email, goa.FormatEmail, err2))
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateUserContext) Created(r *User) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json")
	return ctx.Service.Send(ctx.Context, 201, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateUserContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// DeleteUserContext provides the user delete action context.
type DeleteUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	UserID  uuid.UUID
}

// NewDeleteUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller delete action.
func NewDeleteUserContext(ctx context.Context, service *goa.Service) (*DeleteUserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := DeleteUserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := uuid.FromString(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "uuid"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteUserContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowUserContext provides the user show action context.
type ShowUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	UserID  uuid.UUID
}

// NewShowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller show action.
func NewShowUserContext(ctx context.Context, service *goa.Service) (*ShowUserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ShowUserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := uuid.FromString(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "uuid"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OK(r *User) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateUserContext provides the user update action context.
type UpdateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	UserID  uuid.UUID
	Payload *UpdateUserPayload
}

// NewUpdateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller update action.
func NewUpdateUserContext(ctx context.Context, service *goa.Service) (*UpdateUserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := UpdateUserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := uuid.FromString(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "uuid"))
		}
	}
	return &rctx, err
}

// updateUserPayload is the user update action payload.
type updateUserPayload struct {
	// email of user
	Email       *string `json:"email,omitempty" xml:"email,omitempty" form:"email,omitempty"`
	NewPassword *string `json:"new_password,omitempty" xml:"new_password,omitempty" form:"new_password,omitempty"`
	OldPassword *string `json:"old_password,omitempty" xml:"old_password,omitempty" form:"old_password,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateUserPayload) Validate() (err error) {
	if payload.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *payload.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, *payload.Email, goa.FormatEmail, err2))
		}
	}
	return
}

// Publicize creates UpdateUserPayload from updateUserPayload
func (payload *updateUserPayload) Publicize() *UpdateUserPayload {
	var pub UpdateUserPayload
	if payload.Email != nil {
		pub.Email = payload.Email
	}
	if payload.NewPassword != nil {
		pub.NewPassword = payload.NewPassword
	}
	if payload.OldPassword != nil {
		pub.OldPassword = payload.OldPassword
	}
	return &pub
}

// UpdateUserPayload is the user update action payload.
type UpdateUserPayload struct {
	// email of user
	Email       *string `json:"email,omitempty" xml:"email,omitempty" form:"email,omitempty"`
	NewPassword *string `json:"new_password,omitempty" xml:"new_password,omitempty" form:"new_password,omitempty"`
	OldPassword *string `json:"old_password,omitempty" xml:"old_password,omitempty" form:"old_password,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateUserPayload) Validate() (err error) {
	if payload.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *payload.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`raw.email`, *payload.Email, goa.FormatEmail, err2))
		}
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateUserContext) OK(r *User) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdateUserContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *UpdateUserContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}
