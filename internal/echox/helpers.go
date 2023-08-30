package echox

import (
	"github.com/boichique/avito-test-task/internal/apperrors"
	"github.com/labstack/echo/v4"
)

func Bind[T any](c echo.Context) (*T, error) {
	req := new(T)
	if err := c.Bind(req); err != nil {
		return nil, apperrors.BadRequestHidden(err, "invalid request")
	}

	return req, nil
}
