package logger

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogrusLogger_LogError(t *testing.T) {
	t.Parallel()

	l := NewLogrusLogger()

	l.LogError("message", nil)
	l.LogError("message", errors.New("error"))
}

func TestLogrusLogger_LogInfo(t *testing.T) {
	t.Parallel()

	l := NewLogrusLogger()

	l.LogInfo("message")
}

func TestNewLogrusLogger(t *testing.T) {
	t.Parallel()

	l := NewLogrusLogger()

	assert.Implements(t, (*ILogger)(nil), l)
	assert.IsType(t, (*LogrusLogger)(nil), l)
}
