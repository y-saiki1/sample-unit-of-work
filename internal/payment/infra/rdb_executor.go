package infra

import (
	"upsidr-coding-test/internal/rdb"

	"gorm.io/gorm"
)

type ExecFunc func(tx *gorm.DB) error

type ExecutorRDB struct {
	funcList []ExecFunc
}

func NewExecutorRDB() ExecutorRDB {
	return ExecutorRDB{funcList: make([]ExecFunc, 0, 5)}
}

func (e *ExecutorRDB) Append(fns ...ExecFunc) {
	e.funcList = append(e.funcList, fns...)
}

func (e *ExecutorRDB) Exec() error {
	return rdb.DB.Transaction(func(tx *gorm.DB) error {
		for _, fn := range e.funcList {
			if err := fn(tx); err != nil {
				return err
			}
		}
		return nil
	})
}
