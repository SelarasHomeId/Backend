package affiliate

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
	Create(ctx *abstraction.Context, payload *dto.AffiliateCreateRequest) (map[string]interface{}, error)
}

type service struct {
	AffiliateRepository repository.Affiliate

	DB *gorm.DB
}

func NewService(f *factory.Factory) Service {
	return &service{
		AffiliateRepository: f.AffiliateRepository,

		DB: f.Db,
	}
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.AffiliateCreateRequest) (data map[string]interface{}, err error) {
	if err = trxmanager.New(s.DB).WithTrx(ctx, func(ctx *abstraction.Context) error {
		if err = s.AffiliateRepository.Create(ctx, &model.AffiliateEntityModel{
			Context: ctx,
			AffiliateEntity: model.AffiliateEntity{
				Name:      payload.Name,
				Email:     payload.Email,
				Phone:     payload.Phone,
				Instagram: payload.Instagram,
				Tiktok:    payload.Tiktok,
				CreatedAt: *general.DateTodayLocal(),
			},
		}).Error; err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	data = map[string]interface{}{
		"message": "success",
	}
	return data, nil
}
