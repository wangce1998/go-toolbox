package study

import (
	"testing"
)

func TestLogWrite(t *testing.T) {
	log := NewLogManager(File{})
	log.Write("hello")
}
