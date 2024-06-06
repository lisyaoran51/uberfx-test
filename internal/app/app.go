package app

import "entity"

type app struct {
	db entity.DatabaseInterface
}

func NewApp(db entity.DatabaseInterface) *app {
	return &app{db: db}
}
func (a *app) Run() {
	a.db.Save()
}
