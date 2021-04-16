package main

import (
	"bytes"
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

func decodeString(r io.Reader, w io.Writer) {
	sourceBytes := make([]byte, 1024)
	n, err := r.Read(sourceBytes)
	if err != nil {
		panic(err)
	}

	sourceBytes = sourceBytes[0:n]
	w.Write(base58.Decode(string(sourceBytes)))
}

// getSuppliedFlags gets all flags that were actually entered on the command line
func getSuppliedFlags() map[string]*flag.Flag {
	visitedFlags := map[string]*flag.Flag{}
	flag.Visit(func(flag *flag.Flag) {
		visitedFlags[flag.Name] = flag
	})
	return visitedFlags
}

// TODO: add parameters to allow specifying length, how many, etc.
func main() {
	encFlag := flag.Bool("e", true, "Encode")
	decFlag := flag.Bool("d", false, "Decode")
	flag.Parse()

	suppliedFlags := getSuppliedFlags()
	if *encFlag && *decFlag && suppliedFlags["d"] != nil && suppliedFlags["e"] != nil {
		panic("Cannot specify both -d and -e")
	}

	if *decFlag && suppliedFlags["e"] == nil {
		// They haven't explicitly specified -e, so set the encFlag to false
		*encFlag = false
	}

	var source *bytes.Buffer
	// stdin
	var b []byte = make([]byte, 1024)
	n, err := os.Stdin.Read(b)
	if err != nil {
		panic(err)
	}
	b = b[0:n]
	source = bytes.NewBuffer(b)

	switch {
	case *encFlag:
		encodeBytes(source, os.Stdout)

	case *decFlag:
		decodeString(source, os.Stdout)
	}

	if isTerminal() {
		// Pretty-print by adding a newline to the output
		println()
	}
}
