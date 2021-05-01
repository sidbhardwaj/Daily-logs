package logmanager

var (
	repo Repository
)

type Service interface {
}

type service struct{}

func New(r Repository) Service {
	repo = r
	return &service{}
}
