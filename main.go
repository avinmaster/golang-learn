package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	} else if r.URL.Path == "/contact" || r.URL.Path == "/contact/" {
		fmt.Fprint(w, "<a href=\"mailto:support@test.com\">support@test.com</a>")
	} else {
		fmt.Fprint(w, "<h1>404!</h1>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}