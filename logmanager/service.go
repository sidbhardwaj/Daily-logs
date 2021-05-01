package logmanager

import (
	"context"

	"github.com/sidbhardwaj/Daily-logs/gen/models"
	"github.com/sidbhardwaj/Daily-logs/gen/restapi/operations/todo"
	log "github.com/sirupsen/logrus"
)

type Service interface {
	CreateTODO(ctx context.Context, params todo.CraeteTODOParams) (*models.Todo, error)
	ListTODO(ctx context.Context, params todo.ListTODOParams) (*models.TodoList, error)
}

type service struct {
	repo Repository
}

func New(r Repository) Service {
	return &service{
		repo: r,
	}
}

const (
	defaultUser = "1234"
)

func (s *service) CreateTODO(ctx context.Context, params todo.CraeteTODOParams) (*models.Todo, error) {
	log.Debug("CreateTODO service")
	return s.repo.CreateTODO(ctx, params, defaultUser)
}

func (s *service) ListTODO(ctx context.Context, params todo.ListTODOParams) (*models.TodoList, error) {
	log.Debug("ListTODO service")

	todos, totalSize, err := s.repo.ListTODO(ctx, params)
	if err != nil {
		return nil, err
	}

	var listdata = new(models.TodoList)
	listdata.Data = todos
	listdata.MetaData = &models.MetaData{
		Offset:    *params.Offset,
		PageSize:  *params.PageSize,
		TotalSize: totalSize,
	}
	return listdata, nil
}
