package main

import (
	"testing"
)

func TestTextFilter(t *testing.T) {
	v := TextFilter("hello")
	if v != "hello" {
		t.Error("Failed the test!")
	}
}
