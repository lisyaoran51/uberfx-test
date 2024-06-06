package db

import "github.com/lisyaoran51/uberfx-test/internal/entity"

type database struct{}

func NewDatabase() entity.DatabaseInterface {
	return &database{}
}
func (d *database) Save() {}
