package main

import (
	"testing"
)

func TestRemoveDiacriticsOnLatinText(t *testing.T) {
	result := RemoveDiacritics("hello")
	if result != "hello" {
		t.Error("Failed the test!")
	}
}

func TestRemoveDiacriticsOnFathatan(t *testing.T) {
	result := RemoveDiacritics("مً")
	if result != "م" {
		t.Error("Failed to remove FATHATAN on MEEM")
	}

	result = RemoveDiacritics("نً")
	if result != "ن" {
		t.Error("Failed to remove FATHATAN on NOON")
	}

}

func TestRemoveDiacriticsOnFatha(t *testing.T) {
	result := RemoveDiacritics("بَ")
	if result != "ب" {
		t.Error("Failed to remove FATHA on BA")
	}

	result = RemoveDiacritics("سَلام")
	if result != "سلام" {
		t.Error("Failed to remove FATHA on word SALAAM")
	}
}

func TestRemoveDiacritics_RemoveDamma(t *testing.T) {
	word := "كُرة"

	result := RemoveDiacritics(word)
	if result != "كرة" {
		t.Error("Failed to remove DAMMA")
	}
}

func TestRemoveDiacritics_RemoveFatha(t *testing.T) {
	word := "كرَة"

	result := RemoveDiacritics(word)
	if result != "كرة" {
		t.Error("Failed to remove FATHA")
	}
}

func TestRemoveDiacritics_RemoveDammatan(t *testing.T) {
	word := "كرةٌ"

	result := RemoveDiacritics(word)
	if result != "كرة" {
		t.Error("Failed to remove DAMMATAN")
	}
}

func TestRemoveDiacritics_RemoveKasratan(t *testing.T) {
	word := "كرةٍ"

	result := RemoveDiacritics(word)
	if result != "كرة" {
		t.Error("Failed to remove KASRATAN")
	}
}

func TestRemoveDiacritics_RemoveFathatan(t *testing.T) {
	word := "كرةً"

	result := RemoveDiacritics(word)
	if result != "كرة" {
		t.Error("Failed to remove FATHATAN")
	}
}
func TestRemoveDiacriticsOnWordForBall(t *testing.T) {
	ball_nominative := "كُرَةٌ"
	ball_accusative := "كُرَةً"
	ball_genitive := "كُرَةٍ"

	result := RemoveDiacritics(ball_nominative)
	if result != "كرة" {
		t.Error("Failed to remove DAMMA, FATHA, FATHATAN")
	}

	result = RemoveDiacritics(ball_accusative)
	if result != "كرة" {
		t.Error("Failed to remove DAMMA, FATHA, FATHATAN")
	}

	result = RemoveDiacritics(ball_genitive)
	if result != "كرة" {
		t.Error("Failed to remove DAMMA, FATHA, KASRATAN")
	}
}

func TestRemoveDiacritic_RemoveShadda(t *testing.T) {
	word := "برّ"
	result := RemoveDiacritics(word)
	if result != "بر" {
		t.Error("Failed to remove SHADDA from Arabic word for good")
	}
}

func TestRemoveDiacriticsOnWordForGoodness(t *testing.T) {
	nominative := "بِرٌّ"
	result := RemoveDiacritics(nominative)
	if result != "بر" {
		t.Error("Failed to remove KASRA, SHADDA, DAMMATAN from Arabic word for good")
	}

	accusative := "بِرٌّ"
	result = RemoveDiacritics(accusative)
	if result != "بر" {
		t.Error("Failed to remove KASRA, SHADDA, FATHATAN from Arabic word for good")
	}

	genitive := "بِرٌّ"
	result = RemoveDiacritics(genitive)
	if result != "بر" {
		t.Error("Failed to remove KASRA, SHADDA, KASRATAN from Arabic word for good")
	}
}

func TestRemoveDiacriticsOnWordForSchool(t *testing.T) {
	nominative := "مَدْرَسَةٌ"

	result := RemoveDiacritics(nominative)
	if result != "مدرسة" {
		t.Error("Failed to remove FATHA, SUKUN, FATHA, DAMMATAN from Arabic word for School")
	}
}

func TestRemoveDiacritics_RemoveSuperscriptAlef(t *testing.T) {
	word := "هٰذا"

	result := RemoveDiacritics(word)
	if result != "هذا" {
		t.Error("Failed to remove ARABIC LETTER SUPERSCRIPT ALEF from Arabic word for this (m.sg.)")
	}
}

func TestRemoveDiacritics_RemoveSubscriptAlef(t *testing.T) {
	word := "ٖهذه"

	result := RemoveDiacritics(word)
	if result != "هذه" {
		t.Error("Failed to remove ARABIC SUBSCRIPT ALEF from Arabic word for this (f.sg.)")
	}
}
