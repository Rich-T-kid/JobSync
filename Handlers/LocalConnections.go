package Handlers

import "net/http"

func LocalConnectionHome(w http.ResponseWriter, r *http.Request) {
	info := []byte("hello world local connections ")
	w.Write(info)
}
