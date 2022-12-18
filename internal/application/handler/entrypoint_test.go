package handler_test

import (
	"gcloudrungo/internal/application/handler"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEntrypoint(t *testing.T) {
	r := httptest.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()

	handler.Entrypoint(w, r)

	resp := w.Result()
	_, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("read of response body should not return error, but returned %v", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %v, but returned %v", http.StatusOK, resp.StatusCode)
	}
}
