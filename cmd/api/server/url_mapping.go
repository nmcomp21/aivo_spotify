package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nmcomp/aivo-spotify/controller"
)

type mapping struct {
	albumController *controller.Controller
}

func newMapping() *mapping {
	return &mapping{
		albumController: resolveAlbumController(),
	}
}

func (m mapping) mapUrlsToControllers(router *gin.Engine) {
	router.GET("/api/v1/albums", m.albumController.GetAlbumByBandName)
}
