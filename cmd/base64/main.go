package main

import (
	"encoding/base64"
	"flag"
	"io"
	"os"

	"github.com/yimoucao/go-coreutils/pkg/exit"
)

/*
base64(1)                 BSD General Commands Manual                base64(1)

NAME
     base64 -- Encode and decode using Base64 representation

SYNOPSIS
     base64 [-h | -D | -d] [-b count] [-i input_file] [-o output_file]

DESCRIPTION
     base64 encodes and decodes Base64 data, as specified in RFC 4648. With no options, base64 reads raw data from stdin and writes encoded data as a continuous block to stdout.

OPTIONS
     The following options are available:
     -b count
     --break=count        Insert line breaks every count characters. Default is 0, which generates an unbroken stream.

     -d
     -D
     --decode             Decode incoming Base64 stream into binary data.

     -h
     --help               Print usage summary and exit.

     -i input_file
     --input=input_file   Read input from input_file.  Default is stdin; passing - also represents stdin.

     -o output_file
     --output=output_file
                          Write output to output_file.  Default is stdout; passing - also represents stdout.

SEE ALSO
     openssl(1), wikipedia page <http://en.wikipedia.org/wiki/Base64>, RFC 4648 <http://tools.ietf.org/html/rfc4648>

Mac OS X 10.7                  February 8, 2011                  Mac OS X 10.7
*/

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

	wc := base64.NewEncoder(base64.StdEncoding, output)
	_, err := io.Copy(wc, input)
	if err != nil {
		exit.Error(err)
	}
	wc.Close()
	output.Write([]byte{'\n'})
}
