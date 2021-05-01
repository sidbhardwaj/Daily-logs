package logmanager

import (
	"context"
	"database/sql"

	"github.com/sidbhardwaj/Daily-logs/gen/models"
	"github.com/sidbhardwaj/Daily-logs/gen/restapi/operations/todo"
)

var (
	posDB *sql.DB
)

type Repository interface {
	CreateTODO(ctx context.Context, params todo.CraeteTODOParams) (models.Todo, error)
}

type repository struct{}

func NewRepository(db *sql.DB) Repository {
	posDB = db
	return &repository{}
}

func getDB() *sql.DB {
	return posDB
}
