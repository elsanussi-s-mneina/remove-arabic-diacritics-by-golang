# Go lang application to remove Arabic diacritics.
3 percent implemented.

# Diacritics currently handled
(Note unicode ascii names listed here:)
- ARABIC FATHATAN


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