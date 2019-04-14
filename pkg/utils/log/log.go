package log

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/logrusorgru/aurora"
)

var logger *LogWithLevels

type LogWithLevels struct {
	Debug   *log.Logger
	Warn    *log.Logger
	Error   *log.Logger
	Request *log.Logger
}

type Logger struct {
	*log.Logger
	Options Options
	Aurora  aurora.Value
}

// Initialize instance of LogWithLevels to be used whenever this package is used
func WithLevels(o ...Options) {
	// Return default Loggers if no options declared
	var opts Options
	if len(o) == 0 {
		opts = Options{}
	} else {
		opts = o[0]
	}

	au := aurora.NewAurora(!opts.DisableColors)

	// Create Loggers with prefixes of different logging levels
	logger = &LogWithLevels{
		Debug:   NewLogger(au.Green(" DEBUG "), opts),
		Warn:    NewLogger(au.Brown("WARNING"), opts),
		Error:   NewLogger(au.Red(" ERROR "), opts),
		Request: NewLogger(au.Cyan("REQUEST"), opts),
	}
}

// Returns a new Logger instance with specified options
func NewLogger(level aurora.Value, o Options) *log.Logger {
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

	prefix = fmt.Sprintf("%s%s : ", prefix, level)
	return log.New(output, prefix, flags)
}

// Helper functions to access a specific Logger
func Debug() *log.Logger {
	return logger.Debug
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
