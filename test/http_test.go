package test

import (
	"kawan-usaha-api/server"
	"kawan-usaha-api/server/lib"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	lib.EnvLoaderTest()
	router := server.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"data\":null,\"message\":\"Welcome to Kawan Usaha API!\",\"success\":true}", w.Body.String())
}
