package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Functions struct {
	silent    bool
	file      string
	sensitive bool
	empty     bool
	trim      bool
}

func (l Functions) Duplicate() {
	scanner := bufio.NewScanner(os.Stdin)
	lista := []string{}
	for scanner.Scan() {
		linha := map[bool]string{true: scanner.Text(), false: strings.ToLower(scanner.Text())}[l.sensitive]
		linha = l.processLine(linha)
		if l.empty && linha == "" {
			continue
		}
		if !l.isDuplicate(linha, lista) {
			lista = append(lista, linha)
			switch l.silent {
			case false:
				fmt.Println(linha)
			case true:
				continue
			}
		}
	}
	if l.file != "" {
		CreateFile(l.file, lista)
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
	if l.trim {
		line = strings.TrimSpace(line)
	}
	return line
}
