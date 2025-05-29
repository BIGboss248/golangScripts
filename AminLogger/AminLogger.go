/*
Setup logger with zerolog and ansi colors for terminal output.
*/
package AminLogger

// import packages
import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ANSI color codes for terminal output
const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
	// Text Colors
	FgBlack   = "\033[30m"
	FgRed	 = "\033[31m"
	FgGreen   = "\033[32m"
	FgYellow  = "\033[33m"
	FgBlue	= "\033[34m"
	FgMagenta = "\033[35m"
	FgCyan	= "\033[36m"
	FgWhite   = "\033[37m"
	// Background Colors
	BgBlack   = "\033[40m"
	BgRed	 = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue	= "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan	= "\033[46m"
	BgWhite   = "\033[47m"
)

//tip structs are a collection of variables that can be used to group related data together
//tip structs can't be constants
type AnsiColors struct {
   Reset string
Bold string
// Forground Colors
FgBlack string
FgRed string
FgGreen string
FgYellow string
FgBlue string
FgMagenta string
FgCyan string
FgWhite string
// Background Colors
BgBlack string
BgRed string
BgGreen string
BgYellow string
BgBlue string
BgMagenta string
BgCyan string
BgWhite string
}
// Create a variable of the struct type
var Ansi = AnsiColors{
   Reset : "\033[0m",
Bold  : "\033[1m",
// Forground Colors
FgBlack   : "\033[30m",
FgRed	 : "\033[31m",
FgGreen   : "\033[32m",
FgYellow  : "\033[33m",
FgBlue	: "\033[34m",
FgMagenta : "\033[35m",
FgCyan	: "\033[36m",
FgWhite   : "\033[37m",
// Background Colors,
BgBlack   : "\033[40m",
BgRed	 : "\033[41m",
BgGreen   : "\033[42m",
BgYellow  : "\033[43m",
BgBlue	: "\033[44m",
BgMagenta : "\033[45m",
BgCyan	: "\033[46m",
BgWhite   : "\033[47m",
}

/*
SetupLogger initializes zerolog to write to both console and a file.
*/
func SetupLogger(logFilePath string, logLevel zerolog.Level) (zerolog.Logger, error) {
	// Open or create the log file
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return zerolog.Logger{}, err
	}

	// Console writer with human-friendly formatting
	consoleWriter := zerolog.ConsoleWriter{
		Out:		os.Stdout,
		TimeFormat: time.RFC3339,
	}
	
	// Set global log level
	zerolog.SetGlobalLevel(logLevel)


	// Combine both writers
	multi := zerolog.MultiLevelWriter(consoleWriter, file)

	// Set global time format
	zerolog.TimeFieldFormat = time.RFC3339

	// Create the logger
	logger := zerolog.New(multi).With().Caller().Timestamp().Logger()

	// Set as the global logger
	log.Logger = logger

	return logger, nil
}

// The function that will be executed
func main() {
logger, err := SetupLogger("app.log", zerolog.InfoLevel)
startTime := time.Now() // Record start time
if err != nil {
panic(err)
}
logger.Info().Str("FunctionName:","main").Msg("Main function started")
defer func() {
logger.Info().Str("FunctionName:","main").TimeDiff("Duration (ms)", time.Now(), startTime).Msg("Main function ended.")
}()
	// Code here

}
