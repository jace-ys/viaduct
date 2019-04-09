package logging

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/jace-ys/viaduct/pkg/config"
	"github.com/jace-ys/viaduct/pkg/middleware"
)

type Options struct {
	// Prefix is the keyword in front of log messages. Default: blank (with no brackets)
	Prefix string
	// DisableBrackets if set to true, will remove the square brackets around the prefix. Default: false
	DisableBrackets bool
	// Out is the destination to which the logged data will be written to. Default: os.Stdout
	Out io.Writer
	// Flags defines the logging properties. See http://golang.org/pkg/log/#pkg-constants.
	// To disable all flags, set this to `-1`. Default: log.LstdFlags (2006/01/02 15:04:05)
	Flags int
	// Format of log messages. Default: | {{.Service}} | {{.Status}} | {{.Duration}} | {{.Hostname}} | {{.Method}} {{.RequestURI}}
	Format string
}

type Logger struct {
	*log.Logger
	options Options
}

// logEntry struct to be passed to template
type logEntry struct {
	Service    string
	Status     int
	Duration   time.Duration
	Hostname   string
	Method     string
	RequestURI string
}

// New returns a new Logger instance with declared options
func NewLogger(opts ...Options) *Logger {
	// Return default Logger if
	var o Options
	if len(opts) == 0 {
		o = Options{}
	} else {
		o = opts[0]
	}

	// Determine prefix.
	prefix := o.Prefix
	if len(prefix) > 0 && o.DisableBrackets == false {
		prefix = "[" + prefix + "]"
	}

	// Determine output writer
	var output io.Writer
	if o.Out != nil {
		output = o.Out
	} else {
		// Default is stdout.
		output = os.Stdout
	}

	// Determine output flags.
	flags := log.LstdFlags
	if o.Flags == -1 {
		flags = 0
	} else if o.Flags != 0 {
		flags = o.Flags
	}

	// Set default format for logging
	if o.Format == "" {
		o.Format = "| {{.Service}} | {{.Status}} | {{.Duration}} | {{.Hostname}} | {{.Method}} {{.RequestURI}}"
	}

	return &Logger{
		Logger:  log.New(output, prefix+" ", flags),
		options: o,
	}
}

func CreateMiddleware(logger *Logger, registry *config.ServiceRegistry) middleware.Middleware {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}
		next(rw, r)

		entry := logEntry{
			Service:    getServiceName(registry, r.Host),
			Status:     rw.status, // Fix status logging
			Duration:   time.Since(start),
			Hostname:   r.Host,
			Method:     r.Method,
			RequestURI: r.URL.RequestURI(),
		}

		buffer := &bytes.Buffer{}
		t := template.Must(template.New("log").Parse(logger.options.Format))
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
