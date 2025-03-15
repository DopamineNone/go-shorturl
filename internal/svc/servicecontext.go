package svc

import (
	"go-shorturl/internal/config"
	"go-shorturl/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	ShortUrlModel model.ShortUrlMapModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	return &ServiceContext{
		Config: c,
		ShortUrlModel: model.NewShortUrlMapModel(conn),
	}
}
