package service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nmcomp/aivo-spotify/domain"
)

type (
	AlbumRepo interface {
		GetAlbumsByBandName(c *gin.Context, bandName string) ([]domain.Album, error)
	}
	AlbumService struct {
		albumRepository AlbumRepo
	}
)

func NewAlbumService(albumRepo AlbumRepo) (*AlbumService, error) {
	rc := &AlbumService{
		albumRepository: albumRepo,
	}
	if err := rc.validate(); err != nil {
		return nil, fmt.Errorf("invalid album service: %v", err)
	}
	return rc, nil
}

func (s *AlbumService) validate() error {
	if s.albumRepository == nil {
		return errors.New("album repository client should not be nil")
	}
	return nil
}

func (s AlbumService) GetAlbumsByBandName(c *gin.Context, bandName string) ([]domain.Album, error) {
	return s.albumRepository.GetAlbumsByBandName(c, bandName)
}
