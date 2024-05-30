package repository

import (
	"selarashomeid/internal/abstraction"
	"selarashomeid/internal/model"

	"gorm.io/gorm"
)

type Admin interface {
	FindByUsername(ctx *abstraction.Context, username string) (data *model.AdminEntityModel, err error)
}

type admin struct {
	abstraction.Repository
}

func NewAdmin(db *gorm.DB) *admin {
	return &admin{
		Repository: abstraction.Repository{
			Db: db,
		},
	}
}

func (r *admin) FindByUsername(ctx *abstraction.Context, username string) (data *model.AdminEntityModel, err error) {
	err = r.CheckTrx(ctx).Where("username = ?", username).Take(&data).Error
	return
}
