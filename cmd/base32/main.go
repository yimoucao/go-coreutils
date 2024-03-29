package main

import (
	"encoding/base32"
	"flag"
	"io"
	"os"

	"github.com/yimoucao/go-coreutils/pkg/exit"
)

var (
	inputFile  string
	outputFile string
)

func init() {
	flag.StringVar(&inputFile, "i", "", "input_file")
	flag.StringVar(&outputFile, "o", "", "output_file")
}

func getInput() io.ReadCloser {
	if inputFile == "" {
		inputFile = flag.Arg(0)
		if inputFile == "" {
			return os.Stdin
		}
	}
	f, err := os.Open(inputFile)
	if err != nil {
		exit.Error(err)
	}
	return f
}

func getOutput() io.WriteCloser {
	if outputFile == "" {
		return os.Stdout
	}
	f, err := os.Create(outputFile)
	if err != nil {
		exit.Error(err)
	}
	return f
}

func main() {
	flag.Parse()

	input := getInput()   // TODO: where to close?
	output := getOutput() // TODO: where to close?

	wc := base32.NewEncoder(base32.StdEncoding, output)
	_, err := io.Copy(wc, input)
	if err != nil {
		exit.Error(err)
	}
	wc.Close()
	output.Write([]byte{'\n'})
}
