package global

import (
	"github.com/go-eagle/eagle/infrastructure/storage/orm"
	"gorm.io/gorm"
)

var (
	WriteDB *gorm.DB
	ReadDB  *gorm.DB
)

func InitDB(c *orm.Config) {
	WriteDB = orm.NewMySQL(c)
	ReadDB = orm.NewMySQL(c)
}
