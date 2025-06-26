package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

// TestParseLevel_ValidLevels check if levels are serializing correctly
func TestParseLevel_ValidLevels(t *testing.T) {
	testCases := []struct {
		input    string
		expected zapcore.Level
	}{
		{"debug", zapcore.DebugLevel},
		{"info", zapcore.InfoLevel},
		{"warn", zapcore.WarnLevel},
		{"error", zapcore.ErrorLevel},
	}

	for _, tc := range testCases {
		actual := ParseLevel(tc.input)
		assert.Equal(t, tc.expected, actual, "unexpected log level for input %q", tc.input)
	}
}

// TestParseLevel_InvalidLevelDefaultsToInfo test invalid level behaviour
func TestParseLevel_InvalidLevelDefaultsToInfo(t *testing.T) {
	invalid := "invalid-level"
	result := ParseLevel(invalid)
	assert.Equal(t, zapcore.InfoLevel, result)
}
