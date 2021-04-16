package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"io"
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

func encodeBytes(r *bytes.Buffer, w io.Writer) {
	sourceBytes := r.Bytes()
	encoded := base58.Encode(sourceBytes)
	w.Write([]byte(encoded))
}

// TODO: add parameters to allow specifying length, how many, etc.
func main() {
	encFlag := flag.Bool("e", true, "Encode")
	decFlag := flag.Bool("d", false, "Decode")
	numberFlag := flag.Int("n", 0, "Number of random bytes to generate.  Implies -e."+
		"If absent, ")
	flag.Parse()

	if *decFlag {
		*encFlag = false
	}
	if *numberFlag > 1024 {
		panic("-n is limited to 1024 bytes")
	}
	var source *bytes.Buffer
	if *numberFlag <= 0 {
		// stdin
		var b []byte = make([]byte, 1024)
		n, err := os.Stdin.Read(b)
		if err != nil {
			panic(err)
		}
		b = b[0:n]
		source = bytes.NewBuffer(b)
	} else {
		randBytes := make([]byte, *numberFlag)
		_, err := rand.Read(randBytes)
		if err != nil {
			panic(err)
		}
		source = bytes.NewBuffer(randBytes)
	}
	*encFlag = true

	if *decFlag {
		panic("Cannot specify -d and -e.  (-e is implied by -n)")
	}

	encodeBytes(source, os.Stdout)

	if isTerminal() {
		// Pretty-print by adding a newline to the output
		println()
	}
}
