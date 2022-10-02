package todos

import (
	"github.com/Phazon85/restapp-demo/services/todos"
)

type Service interface {
	GetTodo() ([]*todos.Entry, error)
	GetTodoByID(id string) (*todos.Entry, error)
	DeleteTodo(id string) error
	PostTodo(entry *todos.Entry) error
	PutTodo(entry *todos.Entry) error
}

type Handler struct {
	service Service
}

func New(repo Service) *Handler {
	return &Handler{
		service: repo,
	}
}
