//
// prettyxml.go - a quick hack to pretty print simple XML
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
//
// Copyright (c) 2016, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// * Neither the name of prettyxml nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
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
