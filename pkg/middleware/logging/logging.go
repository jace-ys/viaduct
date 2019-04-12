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
var defaultFormat = "| {{.Service}} | {{.Hostname}} | {{.Method}} {{.RequestURI}} | {{.Status}} {{.Duration}}"

// logEntry struct to be passed to template
type requestEntry struct {
	Service    string
	Status     string
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

		serviceContext := getServiceContext(r, registry)

		entry := requestEntry{
			Service:    serviceContext.Name,
			Status:     proxyStatus(rw.status),
			Duration:   time.Since(start),
			Hostname:   serviceContext.Host,
			Method:     serviceContext.Method,
			RequestURI: serviceContext.RequestURI,
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
