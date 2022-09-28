package todos

import (
	"github.com/Phazon85/restapp-demo/services/todos"
)

type Service interface {
	Get() ([]*todos.Entry, error)
	Delete(id string) error
	Post(entry *todos.Entry) error
	Put(entry *todos.Entry) error
}

type Handler struct {
	service Service
}

func New(repo Service) *Handler {
	return &Handler{
		service: repo,
	}
}
