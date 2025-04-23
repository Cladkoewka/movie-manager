package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/Cladkoewka/movie-manager/internal/repository"
	"github.com/kurin/blazer/b2"
)

type MovieTrailerService struct {
	movieRepository repository.MovieRepository
	bucket *b2.Bucket
	bucketURL string
}

func NewMovieTrailerService(movieRepository repository.MovieRepository, bucket *b2.Bucket, bucketURL string) *MovieTrailerService {
	return &MovieTrailerService{
		movieRepository: movieRepository,
		bucket: bucket,
		bucketURL: bucketURL,
	}
}

func (s *MovieTrailerService) UploadTrailer(movieID int64, file *multipart.FileHeader) error {
	f, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	objectName := fmt.Sprintf("trailers/%d_%s", movieID, file.Filename)

	writer := s.bucket.Object(objectName).NewWriter(context.Background())

	_, err = io.Copy(writer, f)
	if err != nil {
		return fmt.Errorf("failed to copy file content: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}

	trailerURL := fmt.Sprintf("%s/%s", s.bucketURL, objectName)

	return s.movieRepository.UpdateMovieTrailer(movieID, trailerURL)
}

func (s *MovieTrailerService) SetTrailerURL(movieID int64, trailerURL string) error {
	movie, err := s.movieRepository.GetMovieByID(movieID)
	if err != nil {
		return err
	}

	movie.TrailerURL = trailerURL

	_, err = s.movieRepository.UpdateMovie(*movie)
	if err != nil {
		return err
	}

	return nil
}