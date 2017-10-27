package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

//Logger comments
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method+strings.Repeat(" ", 6-len(r.Method)),
			r.RequestURI+strings.Repeat(" ", 30-len(r.RequestURI)),
			name+strings.Repeat(" ", 20-len(name)),
			time.Since(start),
		)
	})
}
