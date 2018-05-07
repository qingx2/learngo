package main

import (
	"net/http"
	"os"

	"github.com/gpmgo/gopm/modules/log"
	"github.com/guopuke/learngo/errhandiling/filelistingserver/filelisting"
)

type appHandler func(write http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {

	return func(writer http.ResponseWriter, r *http.Request) {
		err := handler(writer, r)
		if err != nil {
			// Console 日志记录
			log.Warn("Error handling request: %s", err.Error())

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandlerFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
