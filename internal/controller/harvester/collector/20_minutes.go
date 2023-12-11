package collector

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"hexacrosswords/external/vingt_minutes"
	"hexacrosswords/internal/db"
)

type collector20Minutes struct {
}

func GetCollector20Minutes() Collector {
	return collector20Minutes{}
}

func (c collector20Minutes) GetDefinitions(ctx context.Context, date time.Time) ([]db.Definition, error) {
	formatedDate := date.Format("020106")
	fmt.Printf("collecting definition of date %v\n", formatedDate)

	data, err := vingt_minutes.GetData(ctx, formatedDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get game data : %s", err.Error())
	}

	if data == nil {
		return nil, nil
	}

	solutions, err := c.parse20MinutesDefinitions(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("failed to interprete data : %s", err.Error())
	}

	if len(solutions) == 0 {
		return nil, fmt.Errorf("failed to interprete data : no data")
	}

	return solutions, nil
}

func (c collector20Minutes) parse20MinutesDefinitions(ctx context.Context, gameDataByte []byte) ([]db.Definition, error) {
	gameData, err := c.ParseGameData(ctx, gameDataByte)
	if err != nil {
		return nil, err
	}
	gridInRow := Grid20Minutes(gameData.Grid)
	gridInColumn := Grid20Minutes(gameData.Grid).Reverse()

	wordsInRow, commandsInRow := gridInRow.orderedWordsAndCommands()
	wordsInColumn, _ := gridInColumn.orderedWordsAndCommands()

	definitions := []db.Definition{}

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
			words, err := getWordsFromCommand(command, wordsInRow, wordsInColumn, row, column)
			if err != nil {
				return nil, err
			}
			for _, word := range words {
				definitions = append(definitions, db.Definition{
					Statement: gameData.Definitions[definitionIndex],
					Strength:  gameData.Strength,
					Word:      word,
				})
				definitionIndex++
			}
		}
	}
	return definitions, nil
}

type Grid map[int]string

type GameData struct {
	Strength    int
	Grid        Grid
	Definitions []string
}

var (
	strengthRegex             = regexp.MustCompile(`^force:"(\d)*",$`)
	gridRegex                 = regexp.MustCompile(`^grille:\[$`)
	gridLineRegex             = regexp.MustCompile(`^"(.*)"(\])?(?:,)?$`)
	definitionsRegex          = regexp.MustCompile(`^definitions:\[$`)
	definitionsLineRegex      = regexp.MustCompile(`^\[(?:"([^"]*)",?)+\](?:\])?(?:,)?$`)
	definitionsLineWordRegex  = regexp.MustCompile(`"([^"]*)"`)
	definitionsEndOfListRegex = regexp.MustCompile(`^\[(?:".*",?)+\](\],)(?:,)?$`)
	endOfListRegex            = regexp.MustCompile(`^](?:,)?$`)
)

func (c collector20Minutes) ParseGameData(ctx context.Context, body []byte) (GameData, error) {
	strength := 0
	grid := Grid{}
	definitions := make([]string, 0)

	browsingGrid := false
	gridLineCounter := 0
	browsingDefinitions := false
	for _, line := range strings.Split(string(body), "\n") {
		line = strings.TrimSpace(line)
		if browsingGrid {
			if gridLineRegex.MatchString(line) {
				matches := gridLineRegex.FindStringSubmatch(line)
				grid[gridLineCounter] = matches[1]
				gridLineCounter++
				if len(matches) > 2 && matches[2] != "" {
					browsingGrid = false
				}
				continue
			}

			if endOfListRegex.MatchString(line) {
				browsingGrid = false
				continue
			}

			return GameData{}, fmt.Errorf("line %s should match grid line regex but does not", line)
		}

		if browsingDefinitions {
			if definitionsLineRegex.MatchString(line) {
				matches := definitionsLineWordRegex.FindAllStringSubmatch(line, -1)
				definitionWords := []string{}
				for _, match := range matches {
					definitionWords = append(definitionWords, match[1])
				}
				definition := strings.Join(definitionWords, " ")
				definition = strings.ReplaceAll(definition, "- ", "")
				definition = strings.ReplaceAll(definition, "– ", "")
				definition = strings.ReplaceAll(definition, "%", "")

				definitions = append(definitions, definition)

				if definitionsEndOfListRegex.MatchString(line) {
					browsingDefinitions = false
				}
				continue
			}

			if endOfListRegex.MatchString(line) {
				browsingDefinitions = false
				continue
			}

			return GameData{}, fmt.Errorf("line %s should match definition line regex but does not", line)
		}

		if gridRegex.MatchString(line) {
			browsingGrid = true
			continue
		}

		if definitionsRegex.MatchString(line) {
			browsingDefinitions = true
			continue
		}

		if strengthRegex.MatchString(line) {
			matches := strengthRegex.FindStringSubmatch(line)
			strength, _ = strconv.Atoi(matches[1])
			continue
		}
	}

	return GameData{
		Strength:    strength,
		Grid:        grid,
		Definitions: definitions,
	}, nil
}

type Grid20Minutes map[int]string

func (g Grid20Minutes) Reverse() Grid20Minutes {
	if len(g) == 0 {
		return Grid20Minutes{}
	}

	if len(g[0]) == 0 {
		return Grid20Minutes{}
	}

	newGrid := Grid20Minutes{}
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

func (g Grid20Minutes) orderedWordsAndCommands() (wordWithCoordinate, wordWithCoordinate) {
	words := wordWithCoordinate{}
	commands := wordWithCoordinate{}
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

type wordWithCoordinate map[int]map[int]string

func getWordsFromCommand(command string, wordsInRow, wordsInColumn wordWithCoordinate, row, column int) ([]string, error) {
	switch command {
	case "a":
		return []string{
			wordsInRow[row][column+1],
		}, nil
	case "b":
		return []string{
			wordsInColumn[column][row+1],
		}, nil
	case "c":
		return []string{
			wordsInColumn[column+1][row],
		}, nil
	case "d":
		return []string{
			wordsInRow[row+1][column],
		}, nil
	case "e", "f", "g", "h", "i":
		return []string{
			wordsInRow[row][column+1],
			wordsInColumn[column][row+1],
		}, nil
	case "j", "k", "l", "m", "n":
		return []string{
			wordsInColumn[column+1][row],
			wordsInColumn[column][row+1],
		}, nil
	case "o", "p", "q", "r", "s":
		return []string{
			wordsInRow[row][column+1],
			wordsInRow[row+1][column],
		}, nil
	case "t", "u", "v", "w", "x":
		return []string{
			wordsInColumn[column+1][row],
			wordsInRow[row+1][column],
		}, nil
	case "z":
		return nil, nil
	default:
		return nil, fmt.Errorf("unknown command %s", command)
	}
}
