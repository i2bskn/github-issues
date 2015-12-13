package main

import (
	"strings"
	"testing"
)

func TestNewFormat__defaultFormat(t *testing.T) {
	format := NewFormat(defaultFormat)
	expected := strings.Replace(defaultFormat, "\\t", "\t", -1)
	if format.base != expected {
		t.Fatalf("Expected %v, but %v", expected, format.base)
	}
}
