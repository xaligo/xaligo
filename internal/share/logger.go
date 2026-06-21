package share

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

// LoggerConfig controls logger behavior.
type LoggerConfig struct {
	Component    string `json:"component" yaml:"component"`
	Service      string `json:"service" yaml:"service"`
	Level        string `json:"level" yaml:"level"`
	Structured   bool   `json:"structured" yaml:"structured"`
	EnableCaller bool   `json:"enable_caller" yaml:"enable_caller"`
	Output       string `json:"output" yaml:"output"`
}

// Logger defines the shared logging interface.
type Logger interface {
	DEBUG(mcode MCode, optionalMessage string, fields ...map[string]any)
	INFO(mcode MCode, optionalMessage string, fields ...map[string]any)
	WARN(mcode MCode, optionalMessage string, fields ...map[string]any)
	ERROR(mcode MCode, optionalMessage string, fields ...map[string]any)
	FATAL(mcode MCode, optionalMessage string, fields ...map[string]any)
}

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func (rcvr LogLevel) String() string {
	switch rcvr {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// LogEntry is the structured representation written by Logger.
type LogEntry struct {
	Timestamp string         `json:"timestamp"`
	Level     string         `json:"level"`
	Code      string         `json:"code"`
	Component string         `json:"component,omitempty"`
	Service   string         `json:"service,omitempty"`
	Message   string         `json:"message"`
	Fields    map[string]any `json:"fields,omitempty"`
	File      string         `json:"file,omitempty"`
	Function  string         `json:"function,omitempty"`
	Line      int            `json:"line,omitempty"`
	Error     string         `json:"error,omitempty"`
}

type logger struct {
	config LoggerConfig
	level  LogLevel
	output io.Writer
}

// NewLogger creates a shared logger from config.
func NewLogger(config LoggerConfig) Logger {
	return newLogger(config, nil)
}

func newLogger(config LoggerConfig, output io.Writer) Logger {
	logger := &logger{
		config: config,
		level:  parseLogLevel(config.Level),
		output: output,
	}
	if logger.output == nil {
		logger.output = openLogOutput(config.Output)
	}
	return logger
}

func parseLogLevel(level string) LogLevel {
	switch strings.ToUpper(strings.TrimSpace(level)) {
	case "DEBUG":
		return LevelDebug
	case "WARN", "WARNING":
		return LevelWarn
	case "ERROR":
		return LevelError
	case "FATAL":
		return LevelFatal
	default:
		return LevelInfo
	}
}

func openLogOutput(output string) io.Writer {
	switch strings.TrimSpace(output) {
	case "", "stdout":
		return os.Stdout
	case "stderr":
		return os.Stderr
	default:
		file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "logger output fallback: %s: %v\n", output, err)
			return os.Stdout
		}
		return file
	}
}

func optionalMessage(mcode MCode, optional string) string {
	if optional == "" {
		return mcode.Message
	}
	if mcode.Message == "" {
		return optional
	}
	return fmt.Sprintf("%s: %s", mcode.Message, optional)
}

func (rcvr *logger) DEBUG(mcode MCode, optionalMessage string, fields ...map[string]any) {
	rcvr.log(LevelDebug, mcode, optionalMessage, firstFields(fields))
}

func (rcvr *logger) INFO(mcode MCode, optionalMessage string, fields ...map[string]any) {
	rcvr.log(LevelInfo, mcode, optionalMessage, firstFields(fields))
}

func (rcvr *logger) WARN(mcode MCode, optionalMessage string, fields ...map[string]any) {
	rcvr.log(LevelWarn, mcode, optionalMessage, firstFields(fields))
}

func (rcvr *logger) ERROR(mcode MCode, optionalMessage string, fields ...map[string]any) {
	rcvr.log(LevelError, mcode, optionalMessage, firstFields(fields))
}

func (rcvr *logger) FATAL(mcode MCode, optionalMessage string, fields ...map[string]any) {
	rcvr.log(LevelFatal, mcode, optionalMessage, firstFields(fields))
	os.Exit(1)
}

func firstFields(fields []map[string]any) map[string]any {
	if len(fields) == 0 {
		return nil
	}
	return fields[0]
}

func (rcvr *logger) log(level LogLevel, mcode MCode, message string, fields map[string]any) {
	if level < rcvr.level {
		return
	}
	entry := LogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
		Level:     level.String(),
		Code:      mcode.Code,
		Component: rcvr.config.Component,
		Service:   rcvr.config.Service,
		Message:   optionalMessage(mcode, message),
		Fields:    cloneFields(fields),
	}
	if rcvr.config.EnableCaller || level == LevelDebug {
		entry.setCaller(3)
	}
	entry.extractError()
	rcvr.write(entry)
}

func cloneFields(fields map[string]any) map[string]any {
	if len(fields) == 0 {
		return nil
	}
	cloned := make(map[string]any, len(fields))
	for key, value := range fields {
		cloned[key] = value
	}
	return cloned
}

func (rcvr *LogEntry) setCaller(skip int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	rcvr.File = file
	rcvr.Line = line
	if fn := runtime.FuncForPC(pc); fn != nil {
		rcvr.Function = fn.Name()
	}
}

func (rcvr *LogEntry) extractError() {
	if rcvr.Fields == nil {
		return
	}
	switch value := rcvr.Fields["error"].(type) {
	case nil:
		return
	case error:
		rcvr.Error = value.Error()
	case string:
		rcvr.Error = value
	default:
		rcvr.Error = fmt.Sprint(value)
	}
	delete(rcvr.Fields, "error")
	if len(rcvr.Fields) == 0 {
		rcvr.Fields = nil
	}
}

func (rcvr *logger) write(entry LogEntry) {
	if rcvr.config.Structured {
		data, err := json.Marshal(entry)
		if err == nil {
			fmt.Fprintln(rcvr.output, string(data))
			return
		}
	}
	fmt.Fprintf(rcvr.output, "[%s] [%s] [%s] %s", entry.Timestamp, entry.Level, MCode{Code: entry.Code}.PaddedCode(), entry.Message)
	if len(entry.Fields) > 0 && entry.Level == LevelDebug.String() {
		if data, err := json.Marshal(entry.Fields); err == nil {
			fmt.Fprintf(rcvr.output, " %s", string(data))
		}
	}
	if entry.Error != "" {
		fmt.Fprintf(rcvr.output, " error=%q", entry.Error)
	}
	fmt.Fprintln(rcvr.output)
}
