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
		t.Error("Failed to remove FATHA on BA")
	}

	v = TextFilter("سَلام")
	if v != "سلام" {
		t.Error("Failed to remove FATHA on word SALAAM")
	}
}

func TestTextFilterOnWordForBall(t *testing.T) {
	ball_nominative := "كُرَةٌ"
	ball_accusative := "كُرَةً"
	ball_genitive := "كُرَةٍ"

	v := TextFilter(ball_nominative)
	if v != "كرة" {
		t.Error("Failed to remove DAMMA, FATHA, FATHATAN")
	}

	v = TextFilter(ball_accusative)
	if v != "كرة" {
		t.Error("Failed to remove DAMMA, FATHA, FATHATAN")
	}

	v = TextFilter(ball_genitive)
	if v != "كرة" {
		t.Error("Failed to remove DAMMA, FATHA, KASRATAN")
	}

}

func TestTextFilterOnWordForGoodness(t *testing.T) {
	nominative := "بِرٌّ"
	accusative := "بِرٌّ"
	genitive := "بِرٌّ"

	v := TextFilter(nominative)
	if v != "بر" {
		t.Error("Failed to remove KASRA, SHADDA, DAMMATAN")
	}

	v = TextFilter(accusative)
	if v != "بر" {
		t.Error("Failed to remove KASRA, SHADDA, FATHATAN")
	}

	v = TextFilter(genitive)
	if v != "بر" {
		t.Error("Failed to remove KASRA, SHADDA, KASRATAN")
	}

}

func TestTextFilterOnWordForSchool(t *testing.T) {
	nominative := "مَدْرَسَةٌ"

	v := TextFilter(nominative)
	if v != "مدرسة" {
		t.Error("Failed to remove FATHA, SUKUN, FATHA, DAMMATAN from Arabic word for School")
	}
}
