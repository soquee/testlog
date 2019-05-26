package testlog_test

import (
	"code.soquee.net/testlog"
)

import (
	"testing"
)

func TestLog(t *testing.T) {
	logger := testlog.New(t)
	logger.Println("Logging should not cause a test failure")
}
