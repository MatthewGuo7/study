package httpTest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler2(t *testing.T) {
	reqData := struct {
		Info string `json:"info"`
	}{Info: "P123451"}

	reqBody, _ := json.Marshal(reqData)
	fmt.Println("input:", string(reqBody))
	req := httptest.NewRequest(
		http.MethodPost,
		"/health-check",
		bytes.NewReader(reqBody),
	)

	req.Header.Set("userid", "wdt")
	req.Header.Set("commpay", "brk")

	rr := httptest.NewRecorder()
	HealthCheckHandler(rr, req)

	result := rr.Result()

	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println(string(body))

	if result.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 : %+v",result.StatusCode)
	}

	fmt.Printf("resp = %+v", string(body))
}

func TestHealthCheckHandler(t *testing.T) {
	//创建一个请求
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("start")
	// 我们创建一个 ResponseRecorder (which satisfies http.ResponseWriter)来记录响应
	rr := httptest.NewRecorder()

	//直接使用HealthCheckHandler，传入参数rr,req
	HealthCheckHandler(rr, req)

	// 检测返回的状态码
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// 检测返回的数据
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	fmt.Println(rr.Body.String())
}

