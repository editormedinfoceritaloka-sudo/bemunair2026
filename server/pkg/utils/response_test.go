package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestEnvelopeSuccessAndError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/ok", func(c *gin.Context) { OK(c, "ok", gin.H{"id": 1}) })
	r.GET("/err", func(c *gin.Context) { Error(c, http.StatusForbidden, Forbidden, "forbidden") })

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/ok", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	var ok Response
	_ = json.Unmarshal(w.Body.Bytes(), &ok)
	if !ok.Status || ok.Message != "ok" {
		t.Fatalf("unexpected ok envelope: %+v", ok)
	}

	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/err", nil))
	var er Response
	_ = json.Unmarshal(w.Body.Bytes(), &er)
	if er.Status || er.Error != Forbidden {
		t.Fatalf("unexpected error envelope: %+v", er)
	}
}