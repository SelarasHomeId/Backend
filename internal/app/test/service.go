package test

import (
	"selarashomeid/internal/abstraction"
	"selarashomeid/internal/dto"
	"selarashomeid/internal/factory"
	"selarashomeid/internal/repository"
	"selarashomeid/pkg/gomail"
	"selarashomeid/pkg/util/response"

	"gorm.io/gorm"
)

type Service interface {
	Test(*abstraction.Context) (*dto.TestResponse, error)
	TestGomail(*abstraction.Context, string) (*dto.TestResponse, error)
}

type service struct {
	Repository repository.Test
	Db         *gorm.DB
}

func NewService(f *factory.Factory) Service {
	repository := f.TestRepository
	db := f.Db
	return &service{
		repository,
		db,
	}
}

func (s *service) Test(ctx *abstraction.Context) (*dto.TestResponse, error) {
	result := dto.TestResponse{
		Message: "Success",
	}
	return &result, nil
}

func (s *service) TestGomail(ctx *abstraction.Context, recipient string) (*dto.TestResponse, error) {
	err := gomail.SendMail(recipient)
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}
	result := dto.TestResponse{
		Message: "Success",
	}
	return &result, nil
}
