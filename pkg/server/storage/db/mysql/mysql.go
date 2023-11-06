package mysql

import (
	"gorm.io/gorm"

	"app/pkg/conf"
)

var (
	DB *gorm.DB
)

func Init(cfg *conf.MysqlConfig) (err error) {
	if cfg == nil {
		return
	}
	return
}
