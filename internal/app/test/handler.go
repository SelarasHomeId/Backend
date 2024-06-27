package test

import (
	"selarashomeid/internal/abstraction"
	"selarashomeid/internal/dto"
	"selarashomeid/internal/factory"
	res "selarashomeid/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

var err error

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

// @Summary      Test
// @Description  Test
// @Tags         Test
// @Accept       json
// @Produce      json
// @Success      200      {object}  dto.TestResponse
// @Failure      400      {object}  res.errorResponse
// @Failure      401      {object}  res.errorResponse
// @Failure      404      {object}  res.errorResponse
// @Failure      500      {object}  res.errorResponse
// @Router       /api/v1/test [get]
func (h *handler) Test(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.TestResponse)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	data, err := h.service.Test(cc)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

func (h *handler) TestGomail(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.TestGomailRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	data, err := h.service.TestGomail(cc, payload.Recipient)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
