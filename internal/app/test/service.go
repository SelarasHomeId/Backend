package test

import (
	"selarashomeid/internal/abstraction"
	"selarashomeid/internal/dto"
	"selarashomeid/internal/factory"
	"selarashomeid/internal/repository"

	"gorm.io/gorm"
)

type Service interface {
	Test(*abstraction.Context) (*dto.TestResponse, error)
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
