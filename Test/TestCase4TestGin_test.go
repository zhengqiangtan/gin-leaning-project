package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

//=== RUN   TestPingRoute
//[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
//
//[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//- using env:	export GIN_MODE=release
//- using code:	gin.SetMode(gin.ReleaseMode)
//
//[GIN-debug] GET    /ping                     --> gin-leaning-project/Test.setupRouter.func1 (3 handlers)
//[GIN] 2022/10/08 - 15:05:34 | 200 |            0s |                 | GET      "/ping"
//--- PASS: TestPingRoute (0.01s)
//PASS
