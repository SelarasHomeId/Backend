package notification

import (
	"selarashomeid/internal/abstraction"
	"selarashomeid/internal/factory"
	"selarashomeid/internal/model"
	"selarashomeid/internal/repository"
	"selarashomeid/pkg/util/response"

	"gorm.io/gorm"
)

type Service interface {
	Find(ctx *abstraction.Context) ([]*model.NotificationEntityModel, error)
	CountUnread(ctx *abstraction.Context) (map[string]interface{}, error)
	SetRead(ctx *abstraction.Context, payload *model.SetNotificationRead) (map[string]interface{}, error)
}

type service struct {
	NotificationRepository repository.Notification

	DB *gorm.DB
}

func NewService(f *factory.Factory) Service {
	return &service{
		NotificationRepository: f.NotificationRepository,

		DB: f.Db,
	}
}

func (s *service) Find(ctx *abstraction.Context) ([]*model.NotificationEntityModel, error) {
	var (
		data []*model.NotificationEntityModel
		err  error
	)
	if data, err = s.NotificationRepository.GetAll(ctx); err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}
	return data, nil
}

func (s *service) CountUnread(ctx *abstraction.Context) (data map[string]interface{}, err error) {
	countUnread, err := s.NotificationRepository.CountUnread(ctx)
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}
	return map[string]interface{}{
		"count_unread": countUnread,
	}, nil
}

func (s *service) SetRead(ctx *abstraction.Context, payload *model.SetNotificationRead) (data map[string]interface{}, err error) {
	if err = s.NotificationRepository.SetRead(ctx, &model.NotificationEntityModel{
		Context: ctx,
		ID:      payload.ID,
	}).Error; err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}
	return map[string]interface{}{
		"message": "success",
	}, nil
}
