package global

import (
	"context"
	"fmt"
	"github.com/go-eagle/eagle/infrastructure/logger"
	"gorm.io/gorm"
)

type SysContext interface {
	context.Context
	WriteDB() *gorm.DB
	ReadDB() *gorm.DB
	Transaction(func() error) error
}

type dbCtx struct {
	tx     *gorm.DB
	dbType string
}

func (db *dbCtx) WriteDB() *gorm.DB {
	if db.tx != nil {
		return db.tx
	}
	return WriteDB
}

func (db *dbCtx) ReadDB() *gorm.DB {
	if db.tx != nil {
		return db.tx
	}
	return ReadDB
}

func (db *dbCtx) Transaction(fc func() error) (err error) {
	if db.tx != nil {
		return fc()
	}

	panicked := true
	db.tx = db.WriteDB().Begin()
	defer func() {
		if panicked || err != nil {
			db.tx.Rollback()
		}
		db.tx = nil
	}()

	err = fc()
	if err == nil {
		err1 := db.tx.Commit().Error
		if err1 != nil {
			errMsg := fmt.Sprintf("DB commiy failed, err:%s", err.Error())
			logger.Errorf(errMsg)
			err = fmt.Errorf(errMsg)
		}
	}
	panicked = false
	return err
}

type CommonContext struct {
	context.Context
	*dbCtx
}

func NewCommonContext(c context.Context) *CommonContext {
	if c == nil {
		c = context.Background()
	}
	return &CommonContext{
		Context: c,
		dbCtx:   new(dbCtx),
	}
}
