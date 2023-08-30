package users

import (
	"github.com/boichique/avito-test-task/internal/echox"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

//	@Summary		Create a new user
//	@Description	Create a new user with the given user ID
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		int		true	"User ID"
//	@Success		200		{string}	string	"OK"
//	@Failure		400		{string}	string	"Bad Request"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/api/users/{userID} [post]
func (h *Handler) Create(c echo.Context) error {
	req, err := echox.Bind[CreateUserRequest](c)
	if err != nil {
		return err
	}

	user := &User{
		ID: req.UserID,
	}

	return h.service.Create(c.Request().Context(), user)
}

//	@Summary		Delete a user
//	@Description	Delete a user by user ID
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		int		true	"User ID"
//	@Success		200		{string}	string	"OK"
//	@Failure		400		{string}	string	"Bad Request"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/api/users/{userID} [delete]
func (h *Handler) Delete(c echo.Context) error {
	req, err := echox.Bind[DeleteUserRequest](c)
	if err != nil {
		return err
	}

	return h.service.Delete(c.Request().Context(), req.UserID)
}
