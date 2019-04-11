package logging

import (
	"bytes"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/jace-ys/viaduct/pkg/config"
	"github.com/jace-ys/viaduct/pkg/middleware"
)

// Set default format for logging requests
var defaultFormat = "| {{.Service}} | {{.Status}} | {{.Duration}} | {{.Hostname}} | {{.Method}} {{.RequestURI}}"

// logEntry struct to be passed to template
type requestEntry struct {
	Service    string
	Status     int
	Duration   time.Duration
	Hostname   string
	Method     string
	RequestURI string
}

func CreateMiddleware(logger *log.Logger, registry *config.ServiceRegistry) middleware.Middleware {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}
		next(rw, r)

		entry := requestEntry{
			Service:    getServiceName(registry, r.Host),
			Status:     rw.status, // Fix status logging
			Duration:   time.Since(start),
			Hostname:   r.Host,
			Method:     r.Method,
			RequestURI: r.URL.RequestURI(),
		}

		buffer := &bytes.Buffer{}
		t := template.Must(template.New("log").Parse(defaultFormat))
		err := t.Execute(buffer, entry)
		if err != nil {
			logger.Fatal(err)
		}

		logger.Println(buffer.String())
	}
}

func getServiceName(registry *config.ServiceRegistry, requestedHost string) string {
	for _, service := range registry.Services {
		if service.UpstreamUrl.Host == requestedHost {
			return service.Name
		}
	}
	return requestedHost
}

// Implement wrapper around http.ResponseWriter that provides extra information about the response
type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (w *responseWriter) WriteHeader(s int) {
	w.status = s
	w.ResponseWriter.WriteHeader(s)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.WriteHeader(http.StatusOK)
	size, err := w.ResponseWriter.Write(b)
	w.size += size
	return size, err
}
