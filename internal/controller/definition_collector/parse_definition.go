package definition_collector

import (
	"cross-words/external/vingt_minutes"
	"fmt"
	"strings"
)

func (d DefinitionCollector) ParseDefinitions(data vingt_minutes.GameData) ([]Definition, error) {
	gridInRow := Grid(data.Grid)
	gridInColumn := Grid(data.Grid).Reverse()

	wordsInRow, commandsInRow := gridInRow.OrderedWordsAndCommands()
	wordsInColumn, _ := gridInColumn.OrderedWordsAndCommands()

	definitions := []Definition{}

	definitionIndex := 0
	for row := 0; row < len(gridInRow); row++ {
		commandLine, ok := commandsInRow[row]
		if !ok {
			continue
		}
		for column := 0; column < len(gridInColumn); column++ {
			command, ok := commandLine[column]
			if !ok {
				continue
			}
			words, err := GetWordsFromCommand(command, wordsInRow, wordsInColumn, row, column)
			if err != nil {
				return nil, err
			}
			for _, word := range words {
				definitions = append(definitions, Definition{
					Definition: data.Definitions[definitionIndex],
					Strength:   data.Strength,
					Word:       word,
				})
				definitionIndex++
			}
		}
	}
	return definitions, nil
}

func GetWordsFromCommand(command string, wordsInRow, wordsInColumn WordWithCoordinate, row, column int) ([]string, error) {
	switch command {
	case "a":
		return []string{GetWordFromRow(wordsInRow, row, column+1)}, nil
	case "b":
		return []string{GetWordFromColumn(wordsInColumn, row+1, column)}, nil
	case "c":
		return []string{GetWordFromColumn(wordsInColumn, row, column+1)}, nil
	case "d":
		return []string{GetWordFromRow(wordsInRow, row+1, column)}, nil
	case "e", "f", "g", "h", "i":
		return []string{
			GetWordFromRow(wordsInRow, row, column+1),
			GetWordFromColumn(wordsInColumn, row+1, column),
		}, nil
	case "j", "k", "l", "m", "n":
		return []string{
			GetWordFromColumn(wordsInColumn, row, column+1),
			GetWordFromColumn(wordsInColumn, row+1, column),
		}, nil
	case "o", "p", "q", "r", "s":
		return []string{
			GetWordFromRow(wordsInRow, row, column+1),
			GetWordFromRow(wordsInRow, row+1, column),
		}, nil
	case "t", "u", "v", "w", "x":
		return []string{
			GetWordFromColumn(wordsInColumn, row, column+1),
			GetWordFromRow(wordsInRow, row+1, column),
		}, nil
	case "z":
		return nil, nil
	default:
		return nil, fmt.Errorf("unknown command %s", command)
	}
}

func GetWordFromRow(wordsInColumn WordWithCoordinate, row, column int) string {
	return wordsInColumn[column][row]
}

func GetWordFromColumn(wordsInRow WordWithCoordinate, row, column int) string {
	return wordsInRow[column][row]
}

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
