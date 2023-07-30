package internal

import (
	"bufio"
	"os"

	"github.com/joaoviictorti/line/internal/utils"
)

func CreateFile(file string, lista []string) {
	filename, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		utils.ErroFile()
	}
	defer filename.Close()
	scanner := bufio.NewScanner(filename)
	linesInFile := make(map[string]bool)
	for scanner.Scan() {
		linesInFile[scanner.Text()] = true
	}

	for _, linha := range lista {
		if _, exists := linesInFile[linha]; !exists {
			_, err := filename.WriteString(linha + "\n")
			if err != nil {
				utils.WriteFile()
			}
		}
	}
}
