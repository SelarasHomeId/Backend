package access

import (
	"selarashomeid/internal/abstraction"
	"selarashomeid/internal/dto"
	"selarashomeid/internal/factory"
	"selarashomeid/internal/model"
	"selarashomeid/internal/repository"
	"selarashomeid/pkg/util/general"
	"selarashomeid/pkg/util/response"
	"selarashomeid/pkg/util/trxmanager"

	"gorm.io/gorm"
)

type Service interface {
	Create(ctx *abstraction.Context, payload *dto.AccessCreateRequest) (*dto.AccessCreateRequest, error)
}

type service struct {
	AccessRepository repository.Access

	DB *gorm.DB
}

func NewService(f *factory.Factory) Service {
	return &service{
		AccessRepository: f.AccessRepository,

		DB: f.Db,
	}
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.AccessCreateRequest) (data *dto.AccessCreateRequest, err error) {
	if err = trxmanager.New(s.DB).WithTrx(ctx, func(ctx *abstraction.Context) error {
		modelAccess := &model.AccessEntityModel{
			Context: ctx,
			AccessEntity: model.AccessEntity{
				Module:    *payload.Module,
				CreatedAt: *general.NowLocal(),
			},
		}
		if err = s.AccessRepository.Create(ctx, modelAccess).Error; err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &dto.AccessCreateRequest{
		Module: payload.Module,
		Option: payload.Option,
	}, nil
}
