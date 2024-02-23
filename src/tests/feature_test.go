package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
	"src/controllers"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type processTest struct {
	arg string
	expectedCode int
	expectedMsg string
}

var processTests = []processTest{
	processTest{"json/test1.json", 200, `{"id":"The receipt is invalid"}`},
	processTest{"json/test2.json", 200, `{"id":"The receipt is invalid"}`},
	processTest{"json/test3.json", 400, "The receipt is invalid"},
	processTest{"json/test4.json", 400, "The receipt is invalid"},
	processTest{"json/test5.json", 400, "The receipt is invalid"},
	processTest{"json/test6.json", 400, "The receipt is invalid"},
	processTest{"json/test7.json", 400, "The receipt is invalid"},
	processTest{"json/test8.json", 400, "The receipt is invalid"},
	processTest{"json/test9.json", 400, "The receipt is invalid"},
	processTest{"json/test10.json", 400, "The receipt is invalid"},
	processTest{"json/test11.json", 400, "The receipt is invalid"},
	processTest{"json/test12.json", 400, "The receipt is invalid"},
	processTest{"json/test13.json", 400, "The receipt is invalid"},
	processTest{"json/test14.json", 400, "The receipt is invalid"},
	processTest{"json/test15.json", 400, "The receipt is invalid"},
	processTest{"json/test16.json", 400, "The receipt is invalid"},
	processTest{"json/test17.json", 400, "The receipt is invalid"},
}

func TestProcessReceipt(t *testing.T) {
    r := gin.Default()
    r.POST("/receipts/process", controllers.ProcessReceipt)
    for _,test := range processTests{
    	w := httptest.NewRecorder()
    	file, _ := ioutil.ReadFile(test.arg)
 		reader := bytes.NewReader(file)
 		req, _ := http.NewRequest("POST", "/receipts/process", reader)
 		r.ServeHTTP(w, req)
		assert.Equal(t, test.expectedCode, w.Code)
		assert.Equal(t, test.expectedMsg, w.Body.String())
    }
}


