# Go lang application to remove Arabic diacritics.
90 percent implemented. We just need to set up modules so that it can be imported properly in other projects.

# Diacritics currently handled
(Note unicode ascii names listed here:)
- ARABIC FATHATAN, KASRATAN, DAMMATAN
- ARABIC FATHA, KASRA, DAMMA
- ARABIC SHADDA
- ARABIC SUKUN
- ARABIC LETTER SUPERSCRIPT ALEF
- ARABIC SUBSCRIPT ALEF

# Development Requirements
- Go (See golang.org)

## How to run the program
### With input file.
Assuming you have a text file named "input.txt" in the same directory.
Run the following command:

`go run . < input.txt`

### Interactively
Run the following command:

`go run .`

Then type the text that you want to remove diacritics from and press enter.

To exit, hold the control key and press the C key at the same time.

## How to run the unit tests
Run the following command:

`go test`