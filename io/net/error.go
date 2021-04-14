package net

import (
	"github.com/zhang1career/golab/log"
	"net/http"
	"os"
)

type userErrorInterface interface {
	Code() int
	error
	Message() string
}

func WrapError(a actionInterface) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {                                              // panic
			if r := recover(); r != nil {
				log.Error("Panic:%v", r)
				code, message := http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)
				http.Error(writer, message, code)
			}
		}()
		err := a.Handle(writer, request)
		if err != nil {
			log.Warn("Error handling request.")
			code, message := http.StatusOK, http.StatusText(http.StatusOK)
			if userErr, ok := err.(userErrorInterface); ok {
				code, message = userErr.Code(), userErr.Message()   // user error
			} else {
				code, message = Code(err), Error(err)               // system error
			}
			http.Error(writer, message, code)
		}
	}
}

func Code(err error) (code int) {
	switch {
	case os.IsPermission(err):
		code = http.StatusForbidden
	case os.IsNotExist(err):
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}
	return code
}

func Error(err error) string {
	return http.StatusText(Code(err))
}
