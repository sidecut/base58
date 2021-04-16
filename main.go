package main

import (
	"crypto/rand"
	"flag"
	"os"

	"github.com/btcsuite/btcutil/base58"
)

// From https://rosettacode.org/wiki/Check_output_device_is_a_terminal#Go
func isTerminal() bool {
	fileInfo, err := os.Stdout.Stat()
	if err != nil {
		panic(err)
	}

	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

// TODO: add parameters to allow specifying length, how many, etc.
func main() {
	encFlag := flag.Bool("e", true, "Encode")
	decFlag := flag.Bool("d", false, "Decode")
	numberFlag := flag.Int("n", 0, "Number of random bytes to generate.  Implies -e")
	flag.Parse()

	if *decFlag {
		*encFlag = false
	}
	if *numberFlag > 1024 {
		panic("-n is limited to 1024 bytes")
	}
	if *numberFlag > 0 {
		*encFlag = true

		if *decFlag {
			panic("Cannot specify -d and -e.  (-e is implied by -n)")
		}

		randBytes := make([]byte, *numberFlag)
		_, err := rand.Read(randBytes)
		if err != nil {
			panic(err)
		}

		encrypted := base58.Encode(randBytes)
		print(encrypted)

		if isTerminal() {
			// Pretty-print by adding a newline to the output
			println()
		}

	} else {
		panic("Piping stdin is not currently supported")
	}
}
