package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

const version = "0.0.0"

var (
	showHelp    bool
	showVersion bool
)

func formatXML(data []byte) ([]byte, error) {
	b := &bytes.Buffer{}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	encoder := xml.NewEncoder(b)
	encoder.Indent("", "  ")
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			encoder.Flush()
			return b.Bytes(), nil
		}
		if err != nil {
			return nil, err
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			return nil, err
		}
	}
}

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help information")
	flag.BoolVar(&showVersion, "v", false, "display version information")
}

func main() {
	flag.Parse()
	if len(os.Args) < 2 || showHelp == true {
		fmt.Fprintf(os.Stderr, `
 USAGE %s [OPTIONS] XML_FILENAME

 Pretty print XML documents to standard out.
 
 OPTIONS

`, path.Base(os.Args[0]))
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nVersion %s\n", version)
		if showHelp == true {
			os.Exit(0)
		}
		os.Exit(1)
	}
	if showVersion == true {
		fmt.Fprintf(os.Stderr, "Version %s\n", version)
		os.Exit(0)
	}

	args := flag.Args()
	for _, fname := range args {
		buf, err := ioutil.ReadFile(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error %s, %s\n", fname, err)
			os.Exit(1)
		}
		out, err := formatXML(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error %s, %s\n", fname, err)
		}
		fmt.Printf(`%s`, out)
	}
}
