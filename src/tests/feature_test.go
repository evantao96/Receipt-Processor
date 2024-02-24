package tests

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"src/models"
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
	processReceiptTest{"json/test1.json", 200, `{"id":"\S+"}`},
	processReceiptTest{"json/test2.json", 200, `{"id":"\S+"}`},
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

// Tests for the expected HTTP code and body based on the receipt ID
var getPointsTests = []getPointsTest {
	getPointsTest{"12345", 200, `{"points":\d+}`},
	getPointsTest{"54321", 200, `{"points":\d+}`},
	getPointsTest{"abc", 404, `No receipt found for that id`},

}

func TestProcessReceipt(t *testing.T) {

	// Creates a mock empty receipt map
	var mockReceipts = make(map[string]models.Receipt)

	// Starts a new server and handler
	gin.SetMode(gin.ReleaseMode)
    r := gin.Default()
    h := controllers.ReceiptHandler{Receipts: mockReceipts}
	r.POST("/receipts/process", h.ProcessReceipt)

	// Iterates through tests and sends file data as POST input
    for _,test := range processReceiptTests{
    	w := httptest.NewRecorder()
    	file, _ := ioutil.ReadFile(test.arg)
 		reader := bytes.NewReader(file)
 		req, _ := http.NewRequest("POST", "/receipts/process", reader)
 		r.ServeHTTP(w, req)

 		// Checks if HTTP code and body match the expected results
 		outputCode := w.Code
 		outputBody := w.Body.String()
 		assert.Equal(t, test.expectedCode, outputCode)
 		assert.Regexp(t, test.expectedBody, outputBody)
    }
}

func TestGetPoints(t *testing.T) {

	// Creates a mock receipt map with 2 IDs
	var mockReceipts = map[string]models.Receipt{
						"12345": models.Receipt{
								  "M&M Corner Market",
								  "2022-03-20",
								  "14:33",
								  []models.Item{models.Item{
								      "Gatorade",
								      "2.25",
								    }},
								  "2.25",
								 },
						"54321": models.Receipt{
								  "Target",
								  "2022-01-01",
								  "13:01",
								  []models.Item{models.Item{
								      "Mountain Dew 12PK",
								      "6.49",
								    }},
								  "6.49",
								 },
	}

	// Starts a new server and handler
    r := gin.Default()
    h := controllers.ReceiptHandler{Receipts: mockReceipts}
	r.GET("/receipts/:id/points", h.GetPoints)

	// Iterates through tests and sends GET request 
    for _,test := range getPointsTests{
    	w := httptest.NewRecorder()
 		req, _ := http.NewRequest("GET", "/receipts/" + test.arg + "/points", nil)
 		r.ServeHTTP(w, req)

 		// Checks if HTTP code and body match the expected results
 		outputCode := w.Code
 		outputBody := w.Body.String()
 		assert.Equal(t, test.expectedCode, outputCode)
 		assert.Regexp(t, test.expectedBody, outputBody)
    }
}


