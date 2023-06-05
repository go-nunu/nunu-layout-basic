package dao

import (
	"github.com/go-nunu/nunu-layout-basic/pkg/log"
	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
	//rdb    *redis.Client
	logger *log.Logger
}

func NewDao(logger *log.Logger) *Dao {
	return &Dao{
		//db:     db,
		//rdb:    rdb,
		logger: logger,
	}
}
func NewDb() *gorm.DB {
	// TODO: init db
	//db, err := gorm.Open(mysql.Open(conf.GetString("data.mysql.user")), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	//return db
	return &gorm.DB{}
}
