package net

import (
	"errors"
	"fmt"
	"github.com/zhang1career/app/www/action"
	"github.com/zhang1career/module/error/front"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// Examples
func ExampleCode() {
	var code int
	
	code = Code(os.ErrPermission)
	fmt.Println(code)
	
	code = Code(os.ErrNotExist)
	fmt.Println(code)
	
	code = Code(errors.New("unknown error"))
	fmt.Println(code)
	
	code = Code(nil)
	fmt.Println(code)

	// output:
	// 403
	// 404
	// 500
	// 500
}



// Tests
type panicAction struct {
	action.Action
}
func (act *panicAction) Handle(writer http.ResponseWriter, request *http.Request) (err error) {
	panic(123)
}

type frontAction struct {
	action.Action
}
func (act *frontAction) Handle(writer http.ResponseWriter, request *http.Request) (err error) {
	return &front.Error{http.StatusOK, "ok", "success"}
}

type backAction struct {
	action.Action
}
func (act *backAction) Handle(writer http.ResponseWriter, request *http.Request) (err error) {
	return &front.Error{http.StatusNotFound, "not found", "resource not found"}
}

type systemAction struct {
	action.Action
}
func (act *systemAction) Handle(writer http.ResponseWriter, request *http.Request) (err error) {
	return errors.New("this is a system error")
}

type notFoundAction struct {
	action.Action
}
func (act *notFoundAction) Handle(writer http.ResponseWriter, request *http.Request) (err error) {
	return os.ErrNotExist
}

type noPermissionAction struct {
	action.Action
}
func (act *noPermissionAction) Handle(writer http.ResponseWriter, request *http.Request) (err error) {
	return os.ErrPermission
}

type unknownAction struct {
	action.Action
}
func (act *unknownAction) Handle(writer http.ResponseWriter, request *http.Request) (err error) {
	return errors.New("unknown error")
}

type okAction struct {
	action.Action
}
func (act *okAction) Handle(writer http.ResponseWriter, request *http.Request) (err error) {
	return nil
}


var tests = []struct {
	action  actionInterface
	code    int
	msg     string
}{
	{&panicAction{}, http.StatusInternalServerError, "Internal Server Error"},
	{&frontAction{}, http.StatusOK, "success"},
	{&backAction{}, http.StatusNotFound, "resource not found"},
	{&systemAction{}, http.StatusInternalServerError, "Internal Server Error"},
	{&notFoundAction{}, http.StatusNotFound, "Not Found"},
	{&noPermissionAction{}, http.StatusForbidden, "Forbidden"},
	{&unknownAction{}, http.StatusInternalServerError, "Internal Server Error"},
	{&okAction{}, http.StatusOK, ""},
}


func TestWrapError(t *testing.T) {
	for _, tt := range tests {
		dialog := WrapError(tt.action)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.test.com", nil)
		dialog(response, request)
		verifyResponse(t, response.Result(), tt.code, tt.msg)
	}
}

func TestWrapErrorInServer(t *testing.T) {
	for _, tt := range tests {
		dialog := WrapError(tt.action)
		server := httptest.NewServer(http.HandlerFunc(dialog))
		response, _ := http.Get(server.URL)
		verifyResponse(t, response, tt.code, tt.msg)
	}
}

func verifyResponse(t *testing.T, response *http.Response, expectCode int, expectMsg string) {
	actualCode := response.StatusCode
	all, _ := ioutil.ReadAll(response.Body)
	actualMsg := strings.Trim(string(all), "\n")
	
	if actualCode != expectCode || actualMsg != expectMsg {
		t.Errorf("expect (%d, %s), got (%d, %s)", expectCode, expectMsg, actualCode, actualMsg)
	}
}