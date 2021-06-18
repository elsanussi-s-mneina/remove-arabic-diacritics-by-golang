// Here we are going to go through each character in 
// the Unicode chart for Arabic
// located at https://www.unicode.org/charts/PDF/U0600.pdf
// as of June 18, 2021.

package main

import (
	"testing"
)

func preserveTest(t *testing.T, original, testName string) {
	result := RemoveDiacritics(original)
	if result != original {
		t.Error("Failed to preserve (" + testName + ")!")
	}
}

func TestPreserveArabicNumberSign(t *testing.T) {
	preserveTest(t, "\u0600", "ARABIC NUMBER SIGN")
}

func TestPreserveArabicSignSanah(t *testing.T) {
	preserveTest(t, "\u0601", "ARABIC SIGN SANAH")
}

func TestPreserveArabicFootnoteMarker(t *testing.T) {
	preserveTest(t, "\u0602", "ARABIC FOOTNOTE MARKER")
}

