package log

import (
	"io"
)

type Options struct {
	// Keyword wrapped with square brackets prefixed to log messages. Default: ""
	Prefix string
	// Disable colored logging. Default: false
	DisableColors bool
	// Output destination to write logs to. Default: os.Stdout
	Out io.Writer
	// Flags define the logging properties. See http://golang.org/pkg/log/#pkg-constants.
	// To disable all flags, set to `-1`. Default: log.LstdFlags (2006/01/02 15:04:05)
	Flags int
}
