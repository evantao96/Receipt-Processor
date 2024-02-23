package tests

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"src/controllers"
	"testing"
)

// Struct for expected HTTP code and body based on the receipt file
type processReceiptTest struct {
	arg string
	expectedCode int
	expectedBody string
}

// Struct for expected HTTP code and body based on receipt ID
type getPointsTest struct {
	arg string 
	expectedCode int 
	expectedBody string
}



// Tests for the expected HTTP code and body based on the receipt file
var processReceiptTests = []processReceiptTest {
	processReceiptTest{"json/test1.json", 200, `{"id":`},
	processReceiptTest{"json/test2.json", 200, `{"id":`},
	processReceiptTest{"json/test3.json", 400, "The receipt is invalid"},
	processReceiptTest{"json/test4.json", 400, "The receipt is invalid"},
	processReceiptTest{"json/test5.json", 400, "The receipt is invalid"},
	processReceiptTest{"json/test6.json", 400, "The receipt is invalid"},
	processReceiptTest{"json/test7.json", 400, "The receipt is invalid"},
	processReceiptTest{"json/test8.json", 400, "The receipt is invalid"},
	processReceiptTest{"json/test9.json", 400, "The receipt is invalid"},
	processReceiptTest{"json/test10.json", 400, "The receipt is invalid"},
	processReceiptTest{"json/test11.json", 400, "The receipt is invalid"},
	processReceiptTest{"json/test12.json", 400, "The receipt is invalid"},
	processReceiptTest{"json/test13.json", 400, "The receipt is invalid"},
	processReceiptTest{"json/test14.json", 400, "The receipt is invalid"},
	processReceiptTest{"json/test15.json", 400, "The receipt is invalid"},
	processReceiptTest{"json/test16.json", 400, "The receipt is invalid"},
	processReceiptTest{"json/test17.json", 400, "The receipt is invalid"},
}

var getPointsTests = []getPointsTest {
	getPointsTest{"12345", 200, `{"points":`},
	getPointsTest{"54321", 200, `{"points":`},
	getPointsTest{"abc", 404, `No receipt found for that id`},

}

// Iterates through tests and prints the results
func TestProcessReceipt(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
    r := gin.Default()
    r.POST("/receipts/process", controllers.ProcessReceipt)
    for _,test := range processReceiptTests{
    	w := httptest.NewRecorder()
    	file, _ := ioutil.ReadFile(test.arg)
 		reader := bytes.NewReader(file)
 		req, _ := http.NewRequest("POST", "/receipts/process", reader)
 		r.ServeHTTP(w, req)
 		outputCode := w.Code
 		outputBody := w.Body.String()
 		assert.Equal(t, outputCode, test.expectedCode)
 		assert.Contains(t, outputBody, test.expectedBody)
    }
}

func TestGetPoints(t *testing.T) {
    r := gin.Default()
    r.GET("/receipts/:id/points", controllers.GetPoints)
    for _,test := range getPointsTests{
    	w := httptest.NewRecorder()
 		req, _ := http.NewRequest("GET", "/receipts/" + test.arg + "/points", nil)
 		r.ServeHTTP(w, req)
 		outputCode := w.Code
 		outputBody := w.Body.String()
 		assert.Equal(t, outputCode, test.expectedCode)
 		assert.Contains(t, outputBody, test.expectedBody)
    }
}


