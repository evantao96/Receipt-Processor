package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
	"src/controllers"
	"github.com/gin-gonic/gin"
)

func TestProcessReceipt(t *testing.T) {
	jsonBody := []byte(`{"retailer": "Target"}`)
 	bodyReader := bytes.NewReader(jsonBody)
    r := gin.Default()
    r.POST("/receipts/process", controllers.ProcessReceipt)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/receipts/process", bodyReader)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"status":"The receipt is invalid"}`, w.Body.String())
}