package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/olahol/melody.v1"
)

func TestMain_a2e85e6415(t *testing.T) {
	// Test case 1: Test for successful request
	t.Run("successful request", func(t *testing.T) {
		r := gin.Default()
		m := melody.New()

		r.GET("/ws", func(c *gin.Context) {
			m.HandleRequest(c.Writer, c.Request)
		})

		req, err := http.NewRequest(http.MethodGet, "/ws", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusSwitchingProtocols, resp.Code)
		t.Log("Test for successful request passed")
	})

	// Test case 2: Test for request to non existing route
	t.Run("non existing route", func(t *testing.T) {
		r := gin.Default()
		m := melody.New()

		r.GET("/ws", func(c *gin.Context) {
			m.HandleRequest(c.Writer, c.Request)
		})

		req, err := http.NewRequest(http.MethodGet, "/non-existing-route", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
		t.Log("Test for non existing route passed")
	})
}
