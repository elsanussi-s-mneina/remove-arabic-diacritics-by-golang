// Here we are going to go through each character in
// the Unicode chart for Arabic
// located at https://www.unicode.org/charts/PDF/U0600.pdf
// as of June 18, 2021.

package main

import (
	"fmt"
	"testing"
)

func preserveCharTest(t *testing.T, original rune, testName string) {
	result := RemoveDiacritics(string(original))
	if result != string(original) {
		unicodeCharPoint := fmt.Sprintf("%#U", original)
		t.Error("Failed to preserve " + unicodeCharPoint + " " + testName)
	}
}

func TestPreserveArabicNumberSign(t *testing.T) {
	preserveCharTest(t, '\u0600', "ARABIC NUMBER SIGN")
}

func TestPreserveArabicSignSanah(t *testing.T) {
	preserveCharTest(t, '\u0601', "ARABIC SIGN SANAH")
}

func TestPreserveArabicFootnoteMarker(t *testing.T) {
	preserveCharTest(t, '\u0602', "ARABIC FOOTNOTE MARKER")
}

func TestPreserveArabicSignSafha(t *testing.T) {
	preserveCharTest(t, '\u0603', "ARABIC SIGN SAFHA")
}

func TestPreserveArabicSignSamvat(t *testing.T) {
	preserveCharTest(t, '\u0604', "ARABIC SIGN SAMVAT")
}

func TestPreserveArabicNumberMarkAbove(t *testing.T) {
	preserveCharTest(t, '\u0605', "ARABIC NUMBER MARK ABOVE")
}
func TestPreserveArabicIndicCubeRoot(t *testing.T) {
	preserveCharTest(t, '\u0606', "ARABIC-INDIC CUBE ROOT")
}
func TestPreserveArabicIndicFourthRoot(t *testing.T) {
	preserveCharTest(t, '\u0607', "ARABIC-INDIC FOURTH ROOT")
}
func TestPreserveArabicRay(t *testing.T) {
	preserveCharTest(t, '\u0608', "ARABIC RAY")
}
func TestPreserveArabicIndicPerMilleSign(t *testing.T) {
	preserveCharTest(t, '\u0609', "ARABIC-INDIC PER MILLE SIGN")
}
func TestPreserveArabicArabicIndicPerTenThousandSign(t *testing.T) {
	preserveCharTest(t, '\u060A', "ARABIC-INDIC PER TEN THOUSAND SIGN")
}
func TestPreserveAfghaniSign(t *testing.T) {
	preserveCharTest(t, '\u060B', "Afghani Sign")
}
func TestPreserveArabicComma(t *testing.T) {
	preserveCharTest(t, '\u060C', "ARABIC COMMA")
}
func TestPreserveArabicDateSeparator(t *testing.T) {
	preserveCharTest(t, '\u060D', "ARABIC DATE SEPARATOR")
}
func TestPreserveArabicPoeticVerseSign(t *testing.T) {
	preserveCharTest(t, '\u060E', "ARABIC POETIC VERSE SIGN")
}
func TestPreserveArabicSignMisra(t *testing.T) {
	preserveCharTest(t, '\u060F', "ARABIC SIGN MISRA")
}
func TestPreserveArabicSignSallallahouAlayheWassallam(t *testing.T) {
	preserveCharTest(t, '\u0610', "ARABIC SIGN SALLALLAHOU ALAYHE WASSALLAM")
}
func TestPreserveArabicSignAlayheAssallam(t *testing.T) {
	preserveCharTest(t, '\u0611', "ARABIC SIGN ALAYHE ASSALLAM")
}
func TestPreserveArabicSignRahmatullahAlayhe(t *testing.T) {
	preserveCharTest(t, '\u0612', "ARABIC SIGN RAHMATULLAH ALAYHE")
}
func TestPreserveArabicSignRadiAllahouAnhu(t *testing.T) {
	preserveCharTest(t, '\u0613', "ARABIC SIGN RADI ALLAHOU ANHU")
}
func TestPreserveArabicSignTakhallus(t *testing.T) {
	preserveCharTest(t, '\u0614', "ARABIC SIGN TAKHALLUS")
}
func TestPreserveArabicSmallHighTah(t *testing.T) {
	preserveCharTest(t, '\u0615', "ARABIC SMALL HIGH TAH")
}
func TestPreserveArabicSmallHighLigatureAlefWithLamWithYeh(t *testing.T) {
	preserveCharTest(t, '\u0616', "ARABIC SMALL HIGH LIGATURE ALEF WITH LAM WITH YEH")
}
func TestPreserveArabicSmallHighZain(t *testing.T) {
	preserveCharTest(t, '\u0617', "ARABIC SMALL HIGH ZAIN")
}
func TestPreserveArabicSmallFatha(t *testing.T) {
	preserveCharTest(t, '\u0618', "ARABIC SMALL FATHA")
}
func TestPreserveArabicSmallDamma(t *testing.T) {
	preserveCharTest(t, '\u0619', "ARABIC SMALL DAMMA")
}
func TestPreserveArabicSmallKasra(t *testing.T) {
	preserveCharTest(t, '\u061A', "ARABIC SMALL KASRA")
}
func TestPreserveArabicSemicolon(t *testing.T) {
	preserveCharTest(t, '\u061B', "ARABIC SEMICOLON")
}
func TestPreserveArabicLetterMark(t *testing.T) {
	preserveCharTest(t, '\u061C', "ARABIC LETTER MARK")
}
func TestPreserveArabicTripleDotPunctuationMark(t *testing.T) {
	preserveCharTest(t, '\u061E', "ARABIC TRIPLE DOT PUNCTUATION MARK")
}
func TestPreserveArabicQuestionMark(t *testing.T) {
	preserveCharTest(t, '\u061F', "ARABIC QUESTION MARK")
}
func TestPreserveArabicLetterKashmiriYeh(t *testing.T) {
	preserveCharTest(t, '\u0620', "ARABIC LETTER KASHMIRI YEH")
}
func TestPreserveArabicLetterHamza(t *testing.T) {
	preserveCharTest(t, '\u0621', "ARABIC LETTER HAMZA")
}
func TestPreserveArabicLetterAlefWithMaddaAbove(t *testing.T) {
	preserveCharTest(t, '\u0622', "ARABIC LETTER ALEF WITH MADDA ABOVE")
}
func TestPreserveArabicLetterAlefWithHamzaAbove(t *testing.T) {
	preserveCharTest(t, '\u0623', "ARABIC LETTER ALEF WITH HAMZA ABOVE")
}
func TestPreserveArabicLetterWawWithHamzaAbove(t *testing.T) {
	preserveCharTest(t, '\u0624', "ARABIC LETTER WAW WITH HAMZA ABOVE")
}
func TestPreserveArabicLetterAlefWithHamzaBelow(t *testing.T) {
	preserveCharTest(t, '\u0625', "ARABIC LETTER ALEF WITH HAMZA BELOW")
}

func TestPreserveArabicLetterYehWithHamzaAbove(t *testing.T) {
	preserveCharTest(t, '\u0626', "ARABIC LETTER YEH WITH HAMZA ABOVE")
}

func TestPreserveArabicLetterAlef(t *testing.T) {
	preserveCharTest(t, '\u0627', "ARABIC LETTER ALEF")
}

/*  The following is a template.

func TestPreserve(t* testing.T) {
	preserveCharTest(t, '\u06', "")
}
The template ends here.
*/
