package todos

import (
	"github.com/Phazon85/restapp-demo/services/todos"
)

type Handler struct {
	service *todos.Service
}

func New(repo *todos.Service) *Handler {
	return &Handler{
		service: repo,
	}
}
