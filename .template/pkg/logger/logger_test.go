package logger

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_configure(t *testing.T) {
	assert.Equal(t, log.Flags(), 3)
	Initialise()
	assert.Equal(t, log.Flags(), 7)
}

func Test_loggerOutput(t *testing.T) {
	tests := []struct {
		name     string
		levelFn  func(string, ...interface{})
		message  string
		args     []interface{}
		expected string
	}{
		{
			name:     "test trace output",
			levelFn:  Trace,
			message:  "test trace output %d",
			args:     []interface{}{123},
			expected: "[TRACE] test trace output 123\n",
		},
		{
			name:     "test debug output",
			levelFn:  Debug,
			message:  "test debug output %s",
			args:     []interface{}{"abc"},
			expected: "[DEBUG] test debug output abc\n",
		},
		{
			name:     "test info output",
			levelFn:  Info,
			message:  "test info output %t",
			args:     []interface{}{true},
			expected: "[INFO] test info output true\n",
		},
		{
			name:     "test warn output",
			levelFn:  Warn,
			message:  "test warn output %s",
			args:     []interface{}{"abc"},
			expected: "[WARN] test warn output abc\n",
		},
		{
			name:     "test error output",
			levelFn:  Error,
			message:  "test error output %s",
			args:     []interface{}{"abc"},
			expected: "[ERROR] test error output abc\n",
		},
		{
			name:     "test fatal output",
			levelFn:  Fatal,
			message:  "test fatal output %.2f",
			args:     []interface{}{1234.5},
			expected: "[FATAL] test fatal output 1234.50\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer := bytes.NewBufferString("")
			log.SetOutput(buffer)
			tt.levelFn(tt.message, tt.args...)
			assert.True(t, strings.HasSuffix(buffer.String(), tt.expected))
		})
	}
}
