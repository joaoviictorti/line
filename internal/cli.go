package internal

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func Banner() string {
	banner := `         
	 _     _  _      _____
	/ \   / \/ \  /|/  __/
	| |   | || |\ |||  \  
	| |_/\| || | \|||  /_ 
	\____/\_/\_/  \|\____\											  
	`
	return banner
}

func Argumentos() {

	parser := argparse.NewParser("line", Banner())

	s := parser.Flag(
		"s",
		"silent",
		&argparse.Options{Help: "Send to or file, to disable or stdout", Default: false},
	)

	i := parser.Flag(
		"i",
		"ignored",
		&argparse.Options{Help: "Ignored case sensitive", Default: false},
	)

	e := parser.Flag(
		"e",
		"empty",
		&argparse.Options{Help: "Remove empty lines", Default: false},
	)

	t := parser.Flag(
		"t",
		"trim",
		&argparse.Options{Help: "Remove spaces between lines", Default: false},
	)

	f := parser.String(
		"f",
		"file",
		&argparse.Options{Help: "Insert File", Default: ""},
	)

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		return
	}

	dados := Functions{
		Ignored: *i,
		Silent:  *s,
		File:    *f,
		Empty:   *e,
		Trim:    *t,
	}

	dados.Duplicate()
}
