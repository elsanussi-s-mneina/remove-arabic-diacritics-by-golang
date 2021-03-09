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

func TestTextFilterOnFathatan(t *testing.T) {
	v := TextFilter("مً")
	if v != "م" {
		t.Error("Failed to remove FATHATAN")
	}

	v = TextFilter("نً")
	if v != "ن" {
		t.Error("Failed to remove FATHATAN")
	}

}
