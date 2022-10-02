package todos_test

import (
	"database/sql"
	"reflect"
	"testing"

	. "github.com/Phazon85/restapp-demo/services/todos"
	"go.uber.org/zap"
)

type MockService struct {
	GetTodoRes     []*Entry
	GetTodoByIDRes *Entry
	err            error
}

func (m *MockService) GetTodo() ([]*Entry, error) {
	return m.GetTodoRes, m.err
}

func (m *MockService) GetTodoByID(id string) (*Entry, error) {
	return m.GetTodoByIDRes, m.err
}

func (m *MockService) DeleteTodo(id string) error {
	return m.err
}
func (m *MockService) PostTodo(entry *Entry) error {
	return m.err
}
func (m *MockService) PutTodo(entry *Entry) error {
	return m.err
}

func TestNew(t *testing.T) {
	t.Parallel()
	type args struct {
		logger *zap.Logger
		db     *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Parallel()
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.logger, tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
