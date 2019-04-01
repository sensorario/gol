package gol

import "testing"

func TestConstructor(t *testing.T) {
	l := NewLogger("foo")
	path := l.Path()
	if path != "/var/log/foo/logger.log" {
		t.Errorf("Oops! Not expected path!")
	}
}

func TestCustomLoggerPath(t *testing.T) {
	l := NewCustomLogger("foo/bar")
	path := l.Path()
	if path != "foo/bar/logger.log" {
		t.Errorf("Oops! Not expected path!")
	}
}
