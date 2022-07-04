package ping_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"davidalen.dev/finances/ping"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestIdea(t *testing.T) {
	data := "{\"key\": \"value\"}"
	var j map[string]interface{}

	json.Unmarshal([]byte(data), &j)

	fmt.Println("value is:", j["key"])
}

func TestPing(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	ping.PingRoute(r)

	req, _ := http.NewRequest("GET", "/ping", nil)

	r.ServeHTTP(w, req)

	var body map[string]interface{}

	json.Unmarshal(w.Body.Bytes(), &body)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", body["message"])
}

func TestPingWithMessage(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	ping.PingRoute(r)

	req, _ := http.NewRequest("GET", "/ping?message=alen", nil)

	r.ServeHTTP(w, req)

	var body map[string]interface{}

	json.Unmarshal(w.Body.Bytes(), &body)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "alen", body["message"])
}
