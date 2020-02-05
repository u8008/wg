package log

import (
	"strings"
	"testing"

	"github.com/u8008/wg/test"
)

// Global logger
var l *Logger

// Mocked io.Writer
type tWr struct {
	s string
}

func (tl *tWr) Write(b []byte) (n int, e error) {
	tl.s = string(b)
	return
}

// Global io.Writer
var tw tWr

func TestNewLog(t *testing.T) {
	// Invalid type
	_, e := NewLog(1)
	test.Ok(t, e != nil, "NewLog() must fail")

	// Create logger
	l, e = NewLog(&tw)
	test.Ok(t, e == nil, "NewLog() failed for: %v", e)
}

func TestInfo(t *testing.T) {
	s := "info"
	l.Info(s)
	test.Ok(t, strings.Contains(tw.s, s), "'%v' must contains '%v'", tw.s, s)
}

func TestDebug(t *testing.T) {
	s := "debug"
	l.Debug(s)
	test.Ok(t, !strings.Contains(tw.s, s), "'%v' must not contains '%v'", tw.s, s)
}

func TestFreeLog(t *testing.T) {
	e := FreeLog(l)
	test.Ok(t, e == nil, "FreeLog() failed for: %v", e)
}
