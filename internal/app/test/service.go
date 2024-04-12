package test

import (
	"selarashomeid/internal/abstraction"
	"selarashomeid/internal/dto"
	"selarashomeid/internal/factory"
	"selarashomeid/internal/repository"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Service interface {
	Test(*abstraction.Context) (*dto.TestResponse, error)
}

type service struct {
	Repository repository.Test
	Db         *gorm.DB
	DbRedis    *redis.Client
}

func NewService(f *factory.Factory) Service {
	repository := f.TestRepository
	db := f.Db
	redis := f.DbRedis
	return &service{
		repository,
		db,
		redis,
	}
}

func (s *service) Test(ctx *abstraction.Context) (*dto.TestResponse, error) {
	result := dto.TestResponse{
		Message: "Success",
	}
	return &result, nil
}
