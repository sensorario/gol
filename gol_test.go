package gol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	l := NewLogger("foo")
	path := l.Path()
	assert.Equal(t, "/var/log/foo/logger.log", path, "Oops! Not expected path!")
}

func TestCustomLoggerPath(t *testing.T) {
	l := NewCustomLogger("foo/bar")
	path := l.Path()
	assert.Equal(t, "foo/bar/logger.log", path, "Oops! Not expected path!")
}
