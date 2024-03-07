package tests

import (
	"testing"
	"time"
)

func TestTest(t *testing.T) {
	t.Log("integration test")
}

func TestTerminate(t *testing.T) {
	time.Sleep(3 * time.Second)
}
