package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

var DefaultWriter io.Writer = os.Stdout

// Info logs an [INFO] statement to StdOut.
func Info(logData string, reqID *string) {
	_, _ = fmt.Fprintf(DefaultWriter, "[INFO] %s %s\n", logData, supplyMetadata(reqID))
}

// Debug logs an [DEBUG] statement to StdOut.
func Debug(logData string, reqID *string) {
	_, _ = fmt.Fprintf(DefaultWriter, "[DEBUG] %s %s\n", logData, supplyMetadata(reqID))
}

// Error logs an [ERROR] statement to StdOut.
func Error(logData string, reqID *string) {
	_, _ = fmt.Fprintf(DefaultWriter, "[ERROR] %s %s\n", logData, supplyMetadata(reqID))
}

// Format takes a struct/map and returns a pretty formatted JSON string representation.
func Format(m map[string]any) (string, error) {
	jsonBytes, err := json.Marshal(m)
	return string(jsonBytes), err
}

// supplyMetadata returns a string of log metadata
func supplyMetadata(reqID *string) string {
	mdMap := map[string]any{"service": "erc20-contract-service", "timestamp": time.Now().Format("2006-01-02 15:04:05.000000")}
	if reqID != nil {
		mdMap["requestId"] = reqID
	}
	md, _ := Format(mdMap)
	return md
}
