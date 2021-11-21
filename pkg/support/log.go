package support

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	Reconfig()
}

func Reconfig() {
	zerolog.SetGlobalLevel(getLogLevel())
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.DurationFieldInteger = true
}

func Logger(process string, tags map[string]interface{}) zerolog.Logger {
	return buildBaseLogger(log.Logger).
		With().
		Str("process", process).
		Fields(tags).
		Logger()
}

func getLogLevel() zerolog.Level {
	level := strings.ToLower(GetEnv("REDIVIEW_LOG_LEVEL", "info"))

	switch level {
	case "debug":
		return zerolog.DebugLevel
	case "trace":
		return zerolog.TraceLevel
	default:
		return zerolog.InfoLevel
	}
}

func buildBaseLogger(l zerolog.Logger) zerolog.Logger {
	format := strings.ToLower(GetEnv("LOG_FORMAT", "text"))

	switch format {
	case "json":
		return l
	default:
		return l.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}
