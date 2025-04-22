package service

import (
	"bytes"
	"errors"
	"github.com/Cladkoewka/movie-manager/internal/repository"
	"github.com/Cladkoewka/movie-manager/internal/model"
	"io"
	"mime/multipart"
)

// MoviePosterService отвечает за операции с постерами фильмов
type MoviePosterService struct {
	repo repository.MoviePosterRepository
}

// NewMoviePosterService создает новый сервис для работы с постерами фильмов
func NewMoviePosterService(repo repository.MoviePosterRepository) *MoviePosterService {
	return &MoviePosterService{repo: repo}
}

// SavePoster сохраняет постер фильма в базе данных
func (s *MoviePosterService) SavePoster(movieID int64, file multipart.File, mimeType string) error {
	// Преобразуем файл в байты
	posterBytes, err := convertFileToBytes(file)
	if err != nil {
		return err
	}

	// Сохраняем постер в базе данных
	return s.repo.SavePoster(movieID, posterBytes, mimeType)
}

// GetPosterByMovieID получает постер фильма по его ID
func (s *MoviePosterService) GetPosterByMovieID(movieID int64) (*model.MoviePoster, error) {
	poster, err := s.repo.GetPosterByMovieID(movieID)
	if err != nil {
		return nil, errors.New("poster not found")
	}
	return poster, nil
}

// DeletePoster удаляет постер фильма по его ID
func (s *MoviePosterService) DeletePoster(movieID int64) error {
	return s.repo.DeletePoster(movieID)
}

// Преобразует файл в байтовый массив
func convertFileToBytes(file multipart.File) ([]byte, error) {
	var buffer bytes.Buffer
	_, err := io.Copy(&buffer, file)
	if err != nil {
		return nil, errors.New("failed to read file content")
	}
	return buffer.Bytes(), nil
}
