package di

import (
	"database/sql"
)

type DI interface {
	GetDB() *sql.DB
}

type DependencyInjection struct {
	DB *sql.DB
}

func InitDependencyInjection(db *sql.DB) *DependencyInjection {
	return &DependencyInjection{
		DB: db,
	}
}

func (di *DependencyInjection) GetDB() *sql.DB {
	return di.DB
}
