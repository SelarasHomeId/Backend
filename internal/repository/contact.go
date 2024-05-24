package repository

import (
	"selarashomeid/internal/abstraction"
	"selarashomeid/internal/model"

	"gorm.io/gorm"
)

type Contact interface {
	Create(ctx *abstraction.Context, data *model.ContactEntityModel) *gorm.DB
}

type contact struct {
	abstraction.Repository
}

func NewContact(db *gorm.DB) *contact {
	return &contact{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *contact) Create(ctx *abstraction.Context, data *model.ContactEntityModel) *gorm.DB {
	return r.CheckTrx(ctx).Create(data)
}
