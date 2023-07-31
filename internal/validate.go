package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Functions struct {
	Ignored bool
	Silent  bool
	File    string
	Empty   bool
	Trim    bool
}

func (l Functions) Duplicate() {
	scanner := bufio.NewScanner(os.Stdin)
	lista := []string{}
	for scanner.Scan() {
		linha := map[bool]string{true: scanner.Text(), false: strings.ToLower(scanner.Text())}[l.Ignored]
		linha = l.processLine(linha)
		if l.Empty && linha == "" {
			continue
		}
		if !l.isDuplicate(linha, lista) {
			lista = append(lista, linha)
			switch l.Silent {
			case false:
				fmt.Println(linha)
			case true:
				continue
			}
		}
	}
	if l.File != "" {
		CreateFile(l.File, lista)
	}
}

func (l Functions) isDuplicate(linha string, array []string) bool {
	for _, valor := range array {
		if valor == linha {
			return true
		}
	}
	return false
}

func (l Functions) processLine(line string) string {
	if l.Trim {
		line = strings.TrimSpace(line)
	}
	return line
}
