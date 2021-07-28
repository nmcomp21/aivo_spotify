package controller

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nmcomp/aivo-spotify/domain"
)

type (
	Service interface {
		GetAlbumsByBandName(c *gin.Context, bandName string) ([]domain.Album, error)
	}

	Controller struct {
		albumService Service
	}
)

func NewAlbumController(albumService Service) (*Controller, error) {
	rc := &Controller{
		albumService: albumService,
	}
	if err := rc.validate(); err != nil {
		return nil, fmt.Errorf("invalid album controller: %v", err)
	}
	return rc, nil
}

func (ctrl *Controller) validate() error {
	if ctrl.albumService == nil {
		return errors.New("album service should not be nil")
	}
	return nil
}

func (ctrl Controller) GetAlbumByBandName(c *gin.Context) {
	bandName := c.Query("q")
	album, err := ctrl.albumService.GetAlbumsByBandName(c, bandName)
	if err != nil {
		c.JSON(400, errors.New("error getting albums"))
	}
	c.JSON(200, album)
}
