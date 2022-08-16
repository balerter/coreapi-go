package coreapi

type ModuleLog struct {
	rf requestFunc
}

// Error sends a message to the log with the Error level.
func (l ModuleLog) Error(message string) error {
	_, err := l.rf("log/error", "text/plain", []byte(message))
	return err
}

// Warn sends a message to the log with the Warn level.
func (l ModuleLog) Warn(message string) error {
	_, err := l.rf("log/warn", "text/plain", []byte(message))
	return err
}

// Info sends a message to the log with the Info level.
func (l ModuleLog) Info(message string) error {
	_, err := l.rf("log/info", "text/plain", []byte(message))
	return err
}

// Debug sends a message to the log with the Debug level.
func (l ModuleLog) Debug(message string) error {
	_, err := l.rf("log/debug", "text/plain", []byte(message))
	return err
}
