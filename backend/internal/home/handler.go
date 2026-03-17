package home

import "net/http"

var HomeFunc = func(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("HOMEPAGE"))
	if err != nil {
		panic(err)
	}
}
