package apiserver

import (
	"github.com/stretchr/testify/assert"
	"github.com/yerassylbolatov/http-rest-api/internal/app/store/teststore"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s := newServer(teststore.New())
	s.ServeHttp(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}