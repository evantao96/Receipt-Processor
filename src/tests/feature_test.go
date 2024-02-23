package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"src/controllers"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"fmt"
)

type processReceiptTest struct {
	arg string
	expectedCode int
	expectedBody string
}

type getPointsTest struct {
	arg string 
	expectedCode int 
	expectedBody string
}

var processReceiptTests = []processReceiptTest {
	processReceiptTest{"json/test1.json", 200, `{"id":"^\\S+$"}`},
	processReceiptTest{"json/test2.json", 200, `{"id":"^\\S+$"}`},
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
	getPointsTest{"", 404, "No receipt found for that id"},
	getPointsTest{"abc", 404, "No receipt found for that id"},
}

var id1 = ""
var id2 = ""

func TestProcessReceipt(t *testing.T) {
    r := gin.Default()
    r.POST("/receipts/process", controllers.ProcessReceipt)
    for _,test := range processReceiptTests{
    	w := httptest.NewRecorder()
    	file, _ := ioutil.ReadFile(test.arg)
 		reader := bytes.NewReader(file)
 		req, _ := http.NewRequest("POST", "/receipts/process", reader)
 		r.ServeHTTP(w, req)
 		if outputCode := w.Code; outputCode != test.expectedCode {
 			t.Errorf("Output code %d not equal to expected code %d", outputCode, test.expectedCode)
 		}
 		if outputBody := w.Body.String(); outputBody != test.expectedBody {
			t.Errorf(`Output body "%s" not equal to expected body "%s"`, outputBody, test.expectedBody)
 		}
		if test.arg == "json/test1.json" {
			outputBody := w.Body.String(); id1 = outputBody 
			fmt.Println(id1)
 		}
    }
}

func TestGetPoints(t *testing.T) {
    r := gin.Default()
	r.GET("/receipts/:id/points", controllers.GetPoints)
    for _,test := range getPointsTests{
    	w := httptest.NewRecorder()
    	file, _ := ioutil.ReadFile(test.arg)
 		reader := bytes.NewReader(file)
 		req, _ := http.NewRequest("GET", "/receipts/:id/points", reader)
 		r.ServeHTTP(w, req)
 		if outputCode := w.Code; outputCode != test.expectedCode {
 			t.Errorf("Output code %d not equal to expected code %d", outputCode, test.expectedCode)
 		}
 		if outputBody := w.Body.String(); outputBody != test.expectedBody {
			t.Errorf(`Output body "%s" not equal to expected body "%s"`, outputBody, test.expectedBody)
 		}
    }
}


