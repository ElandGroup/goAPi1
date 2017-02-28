package dao

import (
	. "goApi1/core"
)

type FruitDao struct{}

func (FruitDao) GetAll() (fruits []Fruit, err error) {
	err = Db.Find(&fruits)
	return
}
