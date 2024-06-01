package repository

import (
	"selarashomeid/internal/abstraction"
	"selarashomeid/internal/dto"
	"selarashomeid/internal/model"

	"gorm.io/gorm"
)

type Affiliate interface {
	Create(ctx *abstraction.Context, data *model.AffiliateEntityModel) *gorm.DB
	Find(ctx *abstraction.Context, f *dto.AffiliateFilter, p *abstraction.Pagination) ([]*model.AffiliateEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id int) (data *model.AffiliateEntityModel, err error)
	DeleteByID(ctx *abstraction.Context, id int) *gorm.DB
	GetAll(ctx *abstraction.Context) ([]*model.AffiliateEntityModel, error)
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

func (r *affiliate) Find(ctx *abstraction.Context, f *dto.AffiliateFilter, p *abstraction.Pagination) ([]*model.AffiliateEntityModel, *abstraction.PaginationInfo, error) {
	var (
		data  []*model.AffiliateEntityModel
		count int64
		err   error

		info = &abstraction.PaginationInfo{Pagination: p}
	)

	if err = r.CheckTrx(ctx).Model(&model.AffiliateEntityModel{}).Scopes(func(tx *gorm.DB) *gorm.DB {
		if f != nil {
			f.Apply(tx, ctx)
		}
		return tx
	}).Count(&count).Error; err != nil {
		return nil, nil, err
	}

	if err = r.CheckTrx(ctx).Model(&model.AffiliateEntityModel{}).Scopes(func(tx *gorm.DB) *gorm.DB {
		if f != nil {
			f.Apply(tx, ctx)
		}
		if p != nil {
			if p.Page == nil || p.PageSize == nil {
				p.Init()
			}
			tx.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetOrderBy())
		}
		return tx
	}).Find(&data).Error; err != nil {
		return nil, nil, err
	}

	info.Count = count
	return data, info, nil
}

func (r *affiliate) FindByID(ctx *abstraction.Context, id int) (data *model.AffiliateEntityModel, err error) {
	err = r.CheckTrx(ctx).Where("id = ?", id).Take(&data).Error
	return
}

func (r *affiliate) DeleteByID(ctx *abstraction.Context, id int) *gorm.DB {
	return r.CheckTrx(ctx).Scopes(func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}).Delete(&model.AffiliateEntityModel{})
}

func (r *affiliate) GetAll(ctx *abstraction.Context) (data []*model.AffiliateEntityModel, err error) {
	err = r.CheckTrx(ctx).Find(&data).Error
	return
}
