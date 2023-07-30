package internal

import (
	"flag"
	"fmt"
)

func Banner() {
	banner := `         
	 _     _  _      _____
	/ \   / \/ \  /|/  __/
	| |   | || |\ |||  \  
	| |_/\| || | \|||  /_ 
	\____/\_/\_/  \|\____\											  
	`
	fmt.Println(banner)
}

func Argumentos() {
	var s bool
	var i bool
	var e bool
	var t bool

	flag.Usage = func() {
		Banner()
		flag.PrintDefaults()
	}

	flag.BoolVar(
		&s,
		"s",
		false,
		"Send to or file, to disable or stdout",
	)

	flag.BoolVar(
		&i,
		"i",
		false,
		"Ignored case sensitive",
	)

	flag.BoolVar(
		&e,
		"e",
		false,
		"Remove empty lines",
	)

	flag.BoolVar(
		&t,
		"t",
		false,
		"Remove spaces between lines",
	)

	flag.Parse()
	file := flag.Arg(0)

	dados := Functions{
		sensitive: i,
		silent:    s,
		file:      file,
		empty:     e,
		trim:      t,
	}

	dados.Duplicate()
}
