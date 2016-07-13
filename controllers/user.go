package controllers

import (
	"encoding/hex"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"

	"github.com/elithrar/simple-scrypt"
	"github.com/potatogopher/goa-uuid-ex/app"
	"github.com/potatogopher/goa-uuid-ex/models"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	u := models.User{}

	hash, err := scrypt.GenerateFromPassword([]byte(ctx.Payload.Password), scrypt.DefaultParams)
	if err != nil {
		return ctx.InternalServerError()
	}

	u.Email = ctx.Payload.Email
	u.PasswordHash = hex.EncodeToString(hash)

	user, err := userDB.Add(ctx, &u)
	if err != nil {
		return ErrDatabaseError(err)
	}

	respUser := user.UserToUser()
	respUser.Href = app.UserHref(user.ID)

	ctx.ResponseData.Header().Set("Location", app.UserHref(user.ID))

	return ctx.Created(respUser)
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	var obj models.User
	db := userDB.Db.Delete(&obj, "id=?", ctx.UserID)
	if db.Error != nil {
		return ErrDatabaseError(db.Error)
	}

	if db.RowsAffected == 0 {
		return ctx.NotFound()
	}

	return ctx.NoContent()
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	user, err := userDB.OneUser(ctx.Context, ctx.UserID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ErrDatabaseError(err)
	}

	user.Href = app.UserHref(user.ID)

	return ctx.OK(user)
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	user, err := userDB.Get(ctx.Context, ctx.UserID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ErrDatabaseError(err)
	}

	user.Email = *ctx.Payload.Email

	passwordHash, err := hex.DecodeString(user.PasswordHash)
	if err != nil {
		return ctx.InternalServerError()
	}

	err = scrypt.CompareHashAndPassword(passwordHash, []byte(*ctx.Payload.OldPassword))
	if err != nil {
		return ctx.BadRequest()
	}

	hash, err := scrypt.GenerateFromPassword([]byte(*ctx.Payload.NewPassword), scrypt.DefaultParams)
	if err != nil {
		return ctx.InternalServerError()
	}

	user.PasswordHash = hex.EncodeToString(hash)

	err = userDB.Update(ctx, &user)
	if err != nil {
		return ErrDatabaseError(err)
	}

	respUser := user.UserToUser()
	respUser.Href = app.UserHref(user.ID)

	return ctx.OK(respUser)
}
