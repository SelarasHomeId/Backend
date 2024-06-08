package repository

import (
	"selarashomeid/internal/abstraction"
	"selarashomeid/internal/model"

	"gorm.io/gorm"
)

type Access interface {
	Create(ctx *abstraction.Context, data *model.AccessEntityModel) *gorm.DB
}

type access struct {
	abstraction.Repository
}

func NewAccess(db *gorm.DB) *access {
	return &access{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *access) Create(ctx *abstraction.Context, data *model.AccessEntityModel) *gorm.DB {
	return r.CheckTrx(ctx).Create(data)
}
