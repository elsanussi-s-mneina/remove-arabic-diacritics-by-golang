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
		t.Error("Failed to remove FATHATAN on MEEM")
	}

	v = TextFilter("نً")
	if v != "ن" {
		t.Error("Failed to remove FATHATAN on NOON")
	}

}

func TestTextFilterOnFatha(t *testing.T) {
	v := TextFilter("بَ")
	if v != "ب" {
		t.Error("Failed to remove FATHA")
	}

	v = TextFilter("سَلام")
	if v != "سلام" {
		t.Error("Failed to remove FATHA on word SALAAM")
	}

}
