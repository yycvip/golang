package main

import "net/http"

func main() {
	//fmt.Println("ceshi")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})
	http.ListenAndServe("localhost:80", nil)
}
