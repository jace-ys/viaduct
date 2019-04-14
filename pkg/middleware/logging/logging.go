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
var defaultFormat = "| {{.ApiName}} | {{.Hostname}} | {{.Method}} {{.RequestURI}} | {{.Status}} {{.Duration}}"

// logEntry struct to be passed to template
type requestEntry struct {
	ApiName    string
	Status     string
	Duration   time.Duration
	Hostname   string
	Method     string
	RequestURI string
}

func CreateMiddleware(logger *log.Logger, registry *config.ApiRegistry) middleware.Middleware {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}
		next(rw, r)

		apiContext := getApiContext(r, registry)

		entry := requestEntry{
			ApiName:    apiContext.Name,
			Status:     proxyStatus(rw.status),
			Duration:   time.Since(start),
			Hostname:   apiContext.Host,
			Method:     apiContext.Method,
			RequestURI: apiContext.RequestURI,
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
