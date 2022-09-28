package handlerstest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type responseRecorder struct {
	*httptest.ResponseRecorder
	conn net.Conn
}

func ProcessResults(t *testing.T, c *gin.Context, payloadTarget any) (status int, header http.Header, payload any) {
	t.Helper()
	status = c.Writer.Status()
	header = c.Writer.Header()
	if payloadTarget == nil {
		return status, header, nil
	}
	_, body, err := c.Writer.Hijack()
	if err != nil {
		t.Fatal(fmt.Errorf("gin.responsewriter_hijack: %w", err))
	}
	payload = processPayload(t, json.NewDecoder(body), payloadTarget)

	return status, header, payload
}

func processPayload(t *testing.T, decoder *json.Decoder, target any) any {
	t.Helper()
	decoder.DisallowUnknownFields()

	if targetType := reflect.TypeOf(target); targetType.Kind() == reflect.Pointer {
		payload := reflect.New(targetType.Elem()).Interface()
		err := decoder.Decode(payload)
		if err != nil {
			t.Fatal(fmt.Errorf("json.decoder_decode[0]: %w", err))
		}

		return payload
	}

	payload := reflect.New(reflect.TypeOf(target)).Interface()
	err := decoder.Decode(&payload)
	if err != nil {
		t.Fatal(fmt.Errorf("json.decoder_decode[1]: %w", err))
	}

	return reflect.ValueOf(payload).Elem().Interface()
}

func NewGinContext(t *testing.T, header http.Header, payload any) *gin.Context {
	t.Helper()
	c, _ := gin.CreateTestContext(newRecorder())
	c.Request = &http.Request{}
	c.Request.Header = header
	c.Request.URL = &url.URL{}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(fmt.Errorf("json.Marshal: %w", err))
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(payloadBytes))

	return c
}

func newRecorder() *responseRecorder {
	return &responseRecorder{
		ResponseRecorder: httptest.NewRecorder(),
	}
}

func AssertIntsAreEqual(t *testing.T, want, got int) bool {
	t.Helper()

	return assert.Equal(t, want, got)
}

func AssertObjectsAreEqual(t *testing.T, want, got any) bool {
	t.Helper()

	return assert.True(t, assert.ObjectsAreEqual(want, got), fmt.Sprintf("want: %+v; got: %+v", want, got))
}

func AssertEqualValues(t *testing.T, want, got any) bool {
	t.Helper()

	return assert.EqualValues(t, want, got, fmt.Sprintf("want: %+v; got: %+v", want, got))
}
