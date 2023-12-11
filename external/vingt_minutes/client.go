package vingt_minutes

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func GetData(ctx context.Context, id string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.rcijeux.fr/drupal_game/20minutes/grids/%s.mfj", id))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
