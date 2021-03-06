package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/georgemac/truce"
	"github.com/georgemac/truce/pkg/generate"
)

func main() {
	if len(os.Args) < 3 {
		usage := fmt.Sprintf("Usage: %s", os.Args[0])
		fmt.Printf("%s <command>\n", usage)
		fmt.Printf("%s val[idate] <specification>\n", pad(len(usage)))
		fmt.Printf("%s gen[erate] <specification>\n", pad(len(usage)))
		os.Exit(2)
	}

	targetRaw, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		panic(err)
	}

	var spec truce.Specification

	if err = truce.Unmarshal(targetRaw, &spec); err != nil {
		panic(err)
	}

	switch os.Args[1] {
	case "validate", "val":
		break
	case "generate", "gen":
		for n, versions := range spec.Outputs {
			for v, output := range versions {
				versions, ok := spec.Specifications[n]
				if !ok {
					fmt.Printf("\"%s v%s\" specification not found\n", n, v)
					os.Exit(1)
				}

				api, ok := versions[v]
				if !ok {
					fmt.Printf("\"%s v%s\" specification not found\n", n, v)
					os.Exit(1)
				}

				g, err := generate.New(api)
				if err != nil {
					panic(err)
				}

				if http := output.HTTP; http != nil {
					fmt.Printf("Generating API \"%s v%s\" HTTP Types\n", n, v)

					if target := http.Types; target != nil {
						writeTo(target.Path, g.WriteTypesTo,
							generate.WithPackageName(target.Pkg))
					}

					if target := http.Server; target != nil {
						writeTo(target.Path, g.WriteServerTo,
							generate.WithPackageName(target.Pkg),
							generate.WithServerName(target.Type))
					}

					if target := http.Client; target != nil {
						writeTo(target.Path, g.WriteClientTo,
							generate.WithPackageName(target.Pkg),
							generate.WithClientName(target.Type))
					}
				}
			}
		}
	default:
		fmt.Printf("unexpected sub-command: %q\n", flag.Arg(0))
		os.Exit(2)
	}
}

func writeTo(path string, w func(io.Writer, ...generate.Option) error, opts ...generate.Option) {
	var f io.WriteCloser = nopWriteCloser{os.Stdout}
	if path != "<stdout>" {
		var err error
		f, err = os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	defer f.Close()

	if err := w(f, opts...); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type nopWriteCloser struct {
	io.Writer
}

func (n nopWriteCloser) Close() error { return nil }

func pad(n int) (v string) {
	for i := 0; i < n; i++ {
		v += " "
	}
	return
}
