package interprete

import (
	"cross-words-harverter/httpclient"
	"fmt"
)

type Solution struct {
	Definition string
	Word       string
	Strength   int
}

func (s Solution) String() string {
	return fmt.Sprintf("{\n\tdefinition: %v,\n\tword: %v,\n\tstrength: %v\n}", s.Definition, s.Word, s.Strength)
}

type Interpretor struct {
	Data httpclient.GameData

	gridInRow    Grid
	gridInColumn Grid
	strength     int

	wordsInRow    *WordWithCoordinate
	wordsInColumn *WordWithCoordinate
	commandsInRow *WordWithCoordinate
	definitions   *[]string
	solutions     *[]Solution
}

func NewInterpretor(data httpclient.GameData) Interpretor {
	return Interpretor{
		Data: data,
	}
}

func (i Interpretor) Interprete() ([]Solution, error) {
	i.gridInRow = Grid(i.Data.Grid)
	i.gridInColumn = Grid(i.Data.Grid).Reverse()

	wordsInRow, commandsInRow := i.gridInRow.OrderedWordsAndCommands()
	wordsInColumn, _ := i.gridInColumn.OrderedWordsAndCommands()

	i.wordsInRow = &wordsInRow
	i.wordsInColumn = &wordsInColumn
	i.commandsInRow = &commandsInRow
	i.definitions = &i.Data.Definitions
	i.solutions = &[]Solution{}
	i.strength = i.Data.Strength

	for row := 0; row < len(i.gridInRow); row++ {
		commandLine, ok := commandsInRow[row]
		if !ok {
			continue
		}
		for column := 0; column < len(i.gridInColumn); column++ {
			command, ok := commandLine[column]
			if !ok {
				continue
			}
			switch command {
			case "a":
				i.SetNextSolutionFromRow(row, column+1)
			case "b":
				i.SetNextSolutionFromColumn(row+1, column)
			case "c":
				i.SetNextSolutionFromColumn(row, column+1)
			case "d":
				i.SetNextSolutionFromRow(row+1, column)
			case "e", "f", "g", "h", "i":
				i.SetNextSolutionFromRow(row, column+1)
				i.SetNextSolutionFromColumn(row+1, column)
			case "j", "k", "l", "m", "n":
				i.SetNextSolutionFromColumn(row, column+1)
				i.SetNextSolutionFromColumn(row+1, column)
			case "o", "p", "q", "r", "s":
				i.SetNextSolutionFromRow(row, column+1)
				i.SetNextSolutionFromRow(row+1, column)
			case "t", "u", "v", "w", "x":
				i.SetNextSolutionFromColumn(row, column+1)
				i.SetNextSolutionFromRow(row+1, column)
			case "z":
			default:
				return nil, fmt.Errorf("unknown command : %s", command)
			}
		}
	}
	return *i.solutions, nil
}

func (i Interpretor) SetNextSolutionFromColumn(row, column int) {
	(*i.solutions) = append((*i.solutions), Solution{
		Definition: (*i.definitions)[0],
		Word:       (*i.wordsInColumn)[column][row],
		Strength:   i.strength,
	})
	(*i.definitions) = (*i.definitions)[1:]
}

func (i Interpretor) SetNextSolutionFromRow(row, column int) {
	(*i.solutions) = append((*i.solutions), Solution{
		Definition: (*i.definitions)[0],
		Word:       (*i.wordsInRow)[row][column],
		Strength:   i.strength,
	})

	(*i.definitions) = (*i.definitions)[1:]
}
