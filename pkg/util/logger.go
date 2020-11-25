package util

import "sync"

// Log Level
const (
	Error = iota
	Warning
	Info
	Debug
)

var GlobalLogger *Logger
var Level = Debug

type Logger struct {
	level int
	mu sync.Mutex
}
