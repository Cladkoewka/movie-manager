package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Cladkoewka/movie-manager/internal/model"
	"github.com/Cladkoewka/movie-manager/internal/service"
)

func LoadMoviesFromJSON(movieService *service.MovieService, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var movies []model.Movie
	if err := json.Unmarshal(data, &movies); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	for _, movie := range movies {
		if _, err := movieService.CreateMovie(movie); err != nil {
			return fmt.Errorf("failed to create movie: %w", err)
		}
	}

	return nil
}