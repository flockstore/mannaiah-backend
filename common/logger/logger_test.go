package logger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNew_DebugLevel_PrintsDebug verifies that a logger at "debug" level
// prints both debug and info messages to the provided writer.
func TestNew_DebugLevel_PrintsDebug(t *testing.T) {
	var buf bytes.Buffer
	log := New("debug", &buf)

	log.Debug("debug message")
	log.Info("info message")

	output := buf.String()
	assert.Contains(t, output, "debug message")
	assert.Contains(t, output, "info message")
}

// TestNew_InfoLevel_SkipsDebug verifies that a logger at "info" level
// skips debug messages and prints only info and higher.
func TestNew_InfoLevel_SkipsDebug(t *testing.T) {
	var buf bytes.Buffer
	log := New("info", &buf)

	log.Debug("skip this")
	log.Info("include this")

	output := buf.String()
	assert.NotContains(t, output, "skip this")
	assert.Contains(t, output, "include this")
}
