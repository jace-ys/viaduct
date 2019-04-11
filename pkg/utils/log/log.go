package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

var logger *LogWithLevels

type LogWithLevels struct {
	Info    *log.Logger
	Warn    *log.Logger
	Error   *log.Logger
	Request *log.Logger
}

func WithLevels(opts ...Options) {
	logger = &LogWithLevels{
		Info:    NewLogger("INFO", opts...),
		Warn:    NewLogger("WARNING", opts...),
		Error:   NewLogger("ERROR", opts...),
		Request: NewLogger("REQUEST", opts...),
	}
}

// Returns a new Logger instance, with coloured level added to prefix
func NewLogger(level string, opts ...Options) *log.Logger {
	// Return default Logger if no options declared
	var o Options
	if len(opts) == 0 {
		o = Options{}
	} else {
		o = opts[0]
	}

	// Determine prefix
	prefix := o.Prefix
	if len(prefix) > 0 {
		prefix = "[" + prefix + "] "
	}

	// Determine output writer
	var output io.Writer
	if o.Out != nil {
		output = o.Out
	} else {
		// Default is stdout.
		output = os.Stdout
	}

	// Determine output flags
	flags := log.LstdFlags
	if o.Flags == -1 {
		flags = 0
	} else if o.Flags != 0 {
		flags = o.Flags
	}

	prefix = fmt.Sprintf("%s%s: ", prefix, level)
	return log.New(output, prefix, flags)
}

func Info() *log.Logger {
	return logger.Info
}

func Warn() *log.Logger {
	return logger.Warn
}

func Error() *log.Logger {
	return logger.Error
}

func Request() *log.Logger {
	return logger.Request
}
