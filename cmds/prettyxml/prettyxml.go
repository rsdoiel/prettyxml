//
// prettyxml.go - a quick hack to pretty print simple XML
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const Version = "0.0.1"

var (
	showHelp    bool
	showVersion bool
	//showLicense bool
)

type Element struct {
	Attr     []xml.Attr
	XMLName  xml.Name
	Children []Element `xml:",any"`
	Text     string    `xml:",chardata"`
}

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help information")
	flag.BoolVar(&showVersion, "v", false, "display version information")
	//flag.BoolVar(&showLicense, "l", false, "display license information")
}

func main() {
	appname := os.Args[0]
	flag.Parse()
	args := flag.Args()

	if showHelp == true {
		fmt.Fprintf(os.Stdout, `
 USAGE: %s [OPTIONS] [IN_FILENAME] [OUT_FILENAME]
`, appname)

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("    -%s  (defaults to %s) %s\n", f.Name, f.DefValue, f.Usage)
		})
		fmt.Printf("\n\n Version %s\n", Version)
		os.Exit(0)
	}

	if showVersion == true {
		fmt.Printf(" Version %s\n", Version)
		os.Exit(0)
	}

	var err error
	in := os.Stdin
	out := os.Stdout
	if len(args) > 0 {
		in, err = os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
		defer in.Close()
	}
	if len(args) > 1 {
		out, err = os.Create(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
		defer out.Close()
	}

	src, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	root := Element{}
	_ = xml.Unmarshal(src, &root)
	buf, _ := xml.MarshalIndent(root, "", "   ")
	fmt.Fprintln(out, strings.Replace(string(buf), "&#xA;", "", -1))
}
