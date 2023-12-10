package vingt_minutes

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

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

func GetData(ctx context.Context, id string) (GameData, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.rcijeux.fr/drupal_game/20minutes/grids/%s.mfj", id))
	if err != nil {
		return GameData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GameData{}, err
	}

	return ParseGameData(ctx, body)
}

func ParseGameData(ctx context.Context, body []byte) (GameData, error) {
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
