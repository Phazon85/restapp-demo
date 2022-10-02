package todos_test

// import (
// 	"testing"

// 	"github.com/Phazon85/restapp-demo/handlers/handlerstest"
// 	. "github.com/Phazon85/restapp-demo/handlers/todos"
// 	"github.com/Phazon85/restapp-demo/services/todos"
// )

// type MockService struct {
// 	getResponse []*todos.Entry
// 	err         error
// }

// func (m *MockService) Get() ([]*todos.Entry, error) {
// 	return m.getResponse, m.err
// }

// func (m *MockService) Delete(id string) error {
// 	return m.err
// }

// func (m *MockService) Post(entry *todos.Entry) error {
// 	return m.err
// }

// func (m *MockService) Put(entry *todos.Entry) error {
// 	return m.err
// }

// func TestNew(t *testing.T) {
// 	t.Parallel()
// 	type args struct {
// 		repo Service
// 	}
// 	type want struct {
// 		handler *Handler
// 	}
// 	tests := []struct {
// 		name string
// 		args *args
// 		want *Handler
// 	}{
// 		{
// 			name: "Handler: New() returns expected handler",
// 			args: &args{
// 				repo: &MockService{},
// 			},
// 			want: New(&MockService{}),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			handler := New(tt.args.repo)
// 			handlerstest.AssertObjectsAreEqual(t, tt.want, handler)
// 		})
// 	}
// }
