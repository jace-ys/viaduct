package logging

import (
	"net/http"
)

// Implement wrapper around http.ResponseWriter that provides extra information about the response
type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (w *responseWriter) WriteHeader(s int) {
	w.status = s
}

func (w *responseWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b)
	w.size += size
	return size, err
}

// Check if response status code is 302 Found or 200 OK
func proxyStatus(code int) string {
	if code == 304 || code == 200 {
		return "\u2713"
	}
	return "\u2717"
}
