package main

import (
	"errors"
	"net/http"
	"os"
	"testing"
)

func throwNotFound(http.ResponseWriter, *http.Request) error {
	return os.ErrNotExist
}

func throwNoPermission(http.ResponseWriter, *http.Request) error {
	return os.ErrPermission
}

func throwPanic(http.ResponseWriter, *http.Request) error {
	panic("panic!")
}

func throwUserError(http.ResponseWriter, *http.Request) error {
	return userError("user error")
}

func throwUnknown(http.ResponseWriter, *http.Request) error {
	return errors.New("unknown error")
}

func throwNoError(http.ResponseWriter, *http.Request) error {
	return nil
}

var cases = []struct {
	handler appHandler
	code    int
	message string
}{}

func TestErrorWrapperInServer(t *testing.T) {

}
