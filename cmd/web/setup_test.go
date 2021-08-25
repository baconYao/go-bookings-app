package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// 在執行測試前，做些事情...

	// 執行測試
	os.Exit(m.Run())
}

// myHanlder 實作一個 httpHandler
type myHanlder struct{}

// ServeHTTP 實作一個 httpHandler 需要的 method，有了此 method，就是 httpHandler interface
func (mh *myHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
