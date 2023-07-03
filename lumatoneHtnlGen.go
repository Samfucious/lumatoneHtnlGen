package main

import (
	"fmt"
	"htnlgen/pkg/htnl"
	"htnlgen/pkg/lumatone"
	"io/ioutil"
	"os"
	"strings"
)

type Args struct {
	inputFile string
}

func main() {
	lt := lumatone.NewLumatone()

	args := parseArgs()
	json, err := ioutil.ReadFile(args.inputFile)
	if err != nil {
		panic(err)
	}

	layouts, err := htnl.UnmarshalJSON(json)
	for _, l := range *layouts {
		l.SetPitchesForSection(lt.GetSection(l.Board))
		l.SetChannelForSection(lt.GetSection(l.Board))
		l.SetColorsForSection(lt.GetSection(l.Board))
	}

	var builder strings.Builder
	lt.WriteLtnt(&builder)
	fmt.Println(builder.String())
}

func parseArgs() *Args {
	args := &Args{}

	for i := 0; i < len(os.Args); i++ {
		j := i + 1

		if os.Args[i] == "-f" {
			if j < len(os.Args) {
				args.inputFile = os.Args[j]
			} else {
				panic("not enough arguments")
			}
		}
	}

	return args
}
