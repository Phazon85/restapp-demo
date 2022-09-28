package todos_test

import (
	"net/http"
	"testing"

	"github.com/Phazon85/restapp-demo/handlers/handlerstest"
	. "github.com/Phazon85/restapp-demo/handlers/todos"
	"github.com/gin-gonic/gin"
)

func TestHandler_Get(t *testing.T) {
	t.Parallel()
	type fields struct {
		handler *Handler
	}
	type args struct {
		c *gin.Context
	}
	type want struct {
		status  int
		header  http.Header
		payload any
	}
	tests := []*struct {
		name   string
		fields *fields
		args   *args
		want   *want
	}{
		{
			name: "Handler: Get returns StatusOK",
			fields: &fields{
				handler: New(&MockService{}),
			},
			args: &args{
				c: handlerstest.NewGinContext(t, nil, nil),
			},
			want: &want{
				status:  http.StatusOK,
				header:  http.Header{"Content-Type": []string{"application/json; charset=utf-8"}},
				payload: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.fields.handler.Get(tt.args.c)
			status, header, payload := handlerstest.ProcessResults(t, tt.args.c, tt.want.payload)
			handlerstest.AssertIntsAreEqual(t, tt.want.status, status)
			handlerstest.AssertObjectsAreEqual(t, tt.want.header, header)
			handlerstest.AssertObjectsAreEqual(t, tt.want.payload, payload)

		})
	}
}
