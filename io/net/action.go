package net

import (
	"net/http"
)

type actionInterface interface {
	Handle(http.ResponseWriter, *http.Request) (err error)
}
