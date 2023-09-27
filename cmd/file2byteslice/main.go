// fileToByteSlice is a dead simple tool to embed a file to Go.
package main

import (
	"flag"
	"io"
	"os"

	"github.com/yngfoxx/fileToByteSlice"
)

var (
	inputFilename  = flag.String("input", "", "input filename")
	outputFilename = flag.String("output", "", "output filename")
	packageName    = flag.String("package", "main", "package name")
	varName        = flag.String("var", "_", "variable name")
	compress       = flag.Bool("compress", false, "use gzip compression")
	buildTags      = flag.String("buildtags", "", "build tags")
)

func run() error {
	var out io.Writer
	if *outputFilename != "" {
		f, err := os.Create(*outputFilename)
		if err != nil {
			return err
		}
		defer f.Close()
		out = f
	} else {
		out = os.Stdout
	}

	var in io.Reader
	if *inputFilename != "" {
		f, err := os.Open(*inputFilename)
		if err != nil {
			return err
		}
		defer f.Close()
		in = f
	} else {
		in = os.Stdin
	}

	if err := fileToByteSlice.Write(out, in, *compress, *buildTags, *packageName, *varName); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		panic(err)
	}
}
