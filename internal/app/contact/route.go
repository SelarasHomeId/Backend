package contact

import "github.com/labstack/echo/v4"

func (h *handler) Route(v *echo.Group) {
	v.GET("", h.Create)
}
