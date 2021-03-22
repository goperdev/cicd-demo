package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!test 11111 22222 3333"))
	})
	http.ListenAndServe(":8000", nil)
}
