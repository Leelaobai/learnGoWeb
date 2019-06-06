package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	h := request.Header
	fmt.Fprintln(writer, h)
}

func body(w http.ResponseWriter, r *http.Request) {
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location","http://google.com")
	w.WriteHeader(302)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/body", body)
	http.HandleFunc("/process", process)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)

	http.ListenAndServe(":8080", nil)
}
