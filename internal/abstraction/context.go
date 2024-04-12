package abstraction

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Context struct {
	echo.Context
	Trx *TrxContext
}

type TrxContext struct {
	Db *gorm.DB
}
