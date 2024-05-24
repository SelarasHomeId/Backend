package repository

import (
	"selarashomeid/internal/abstraction"
	"selarashomeid/internal/model"

	"gorm.io/gorm"
)

type Affiliate interface {
	Create(ctx *abstraction.Context, data *model.AffiliateEntityModel) *gorm.DB
}

type affiliate struct {
	abstraction.Repository
}

func NewAffiliate(db *gorm.DB) *affiliate {
	return &affiliate{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *affiliate) Create(ctx *abstraction.Context, data *model.AffiliateEntityModel) *gorm.DB {
	return r.CheckTrx(ctx).Create(data)
}
