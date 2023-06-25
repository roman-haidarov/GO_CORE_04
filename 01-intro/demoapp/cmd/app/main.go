package main

import (
	"go-core-04/01-intro/demoapp/pkg/reversestr"
	"net/http"
)

func main() {
	http.ListenAndServe(
		":8080",
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				s := r.URL.Query()["s"]
				if len(s) == 0 {
					http.Error(w, "bad query", http.StatusBadRequest)
				}

				rev := reversestr.Rev(s[0])

				w.Write([]byte(rev))
			},
		),
	)
}