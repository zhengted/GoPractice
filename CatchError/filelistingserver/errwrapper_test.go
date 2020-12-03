package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNotPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("Unknown Error")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println("no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNotPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, ""},
}

// 测试代码		速度快 偏向单元测试
func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder() // fake request
		request := httptest.NewRequest(    // fake response
			http.MethodGet,
			"http://imooc.com", nil)
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}

}

// 测试整个服务器  速度慢 偏向整体测试

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response,
	expectedCode int, expectedMsg string, t *testing.T) {
	all, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(all), "\n")
	if resp.StatusCode != expectedCode ||
		body != expectedMsg {
		t.Errorf("expect (%d,%s)"+"got (%d,%s)",
			expectedCode, expectedMsg,
			resp.StatusCode, body)
	}
}
