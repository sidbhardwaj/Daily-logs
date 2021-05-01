package logmanager

import (
	"context"

	"github.com/ido50/sqlz"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sidbhardwaj/Daily-logs/gen/models"
	"github.com/sidbhardwaj/Daily-logs/gen/restapi/operations/todo"
	log "github.com/sirupsen/logrus"
)

type Repository interface {
	CreateTODO(ctx context.Context, params todo.CraeteTODOParams, userID string) (*models.Todo, error)
	ListTODO(ctx context.Context, params todo.ListTODOParams) ([]*models.Todo, int64, error)
}

type repository struct {
	posDB *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{posDB: db}
}

const (
	todoTable = "daily_log.todo"
)

var todoCol = []string{
	"id as ID",
	"coalesce(title, '') as Title",
	"coalesce(description, '') as Description",
	"coalesce(status, '') as Status",
	"created_at as CreatedAT",
	"updated_at as UpdatedAT",
}

func (repo *repository) CreateTODO(ctx context.Context, params todo.CraeteTODOParams, userID string) (*models.Todo, error) {
	log.Debug("CreateTODO repo...")
	values := make(map[string]interface{})

	values["title"] = params.Todo.Title
	values["description"] = params.Todo.Description
	values["status"] = params.Todo.Status
	values["userid"] = userID

	query := sqlz.Newx(repo.posDB).
		InsertInto(todoTable).
		ValueMap(values).Returning(todoCol...)

	sqlRaw, b := query.ToSQL(true)
	log.Debugf("SQL: %s, DATA: %+v", sqlRaw, b)

	var result = TODOSQL{}

	err := query.GetRow(&result)
	if err != nil {
		return nil, errors.Wrap(err, "failed while inserting in todo table")
	}

	return result.toTODO(), nil
}

func (repo *repository) ListTODO(ctx context.Context, params todo.ListTODOParams) ([]*models.Todo, int64, error) {
	log.Debug("ListTODO repo ...")
	query := sqlz.Newx(repo.posDB).
		Select(todoCol...).
		From(todoTable).
		Offset(*params.Offset)

	sqlRaw, b := query.ToSQL(true)
	log.Debugf("SQL: %s, DATA: %+v", sqlRaw, b)

	totalCount, err := query.GetCount()
	if err != nil {
		return nil, 0, errors.Wrap(err, "unable to get todo list total size")
	}
	query = query.Limit(*params.PageSize)

	var results = []TODOSQL{}
	err = query.GetAll(&results)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed while getting todo from table")
	}
	var todos = []*models.Todo{}
	for _, res := range results {
		todos = append(todos, res.toTODO())
	}
	return todos, totalCount, nil
}
