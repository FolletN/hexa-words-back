package interprete

import (
	"fmt"
	"strings"
)

type Grid map[int]string

func (g Grid) Reverse() Grid {
	if len(g) == 0 {
		return Grid{}
	}

	if len(g[0]) == 0 {
		return Grid{}
	}

	newGrid := Grid{}
	for i := 0; i < len(g); i++ {
		line := g[i]
		for j, character := range line {
			newLine := newGrid[j]
			newLineRune := []rune(newLine)
			newLineRune = append(newLineRune, character)
			newGrid[j] = string(newLineRune)
		}
	}
	return newGrid
}

func (g Grid) Display() {
	for i := 0; i < len(g); i++ {
		fmt.Println(g[i])
	}
}

type WordWithCoordinate map[int]map[int]string

func (g Grid) OrderedWordsAndCommands() (WordWithCoordinate, WordWithCoordinate) {
	words := WordWithCoordinate{}
	commands := WordWithCoordinate{}
	for i := 0; i < len(g); i++ {
		words[i] = map[int]string{}
		commands[i] = map[int]string{}
		line := g[i]

		word := ""
		firstColumnWord := -1
		for j := 0; j < len(line); j++ {
			character := string(line[j])
			if !isCharacter(character) {
				commands[i][j] = character
				if len(word) > 1 {
					words[i][firstColumnWord] = word
				}
				word = ""
				firstColumnWord = -1
				continue
			}
			if firstColumnWord == -1 {
				firstColumnWord = j
			}

			word += character
		}
		if len(word) > 1 {
			words[i][firstColumnWord] = word
		}
	}
	return words, commands
}

func isCharacter(character string) bool {
	return strings.ToUpper(character) == character
}
