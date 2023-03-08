package main

import (
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := ioutil.ReadFile("./public/tyger.jpg")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}

func main() {
	handler := http.HandlerFunc(handler)
	http.Handle("/image/tyger.jpg", handler)
	http.ListenAndServe(":9090", nil)
}
