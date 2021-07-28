package server

import (
	"fmt"
	"net/http"

	"github.com/nmcomp/aivo-spotify/controller"
	"github.com/nmcomp/aivo-spotify/repository"
	"github.com/nmcomp/aivo-spotify/service"
)

func resolveAlbumController() *controller.Controller {
	ctrl, err := controller.NewAlbumController(
		resolveAlbumService(),
	)
	if err != nil {
		panicHandler(err)
	}
	return ctrl
}
func panicHandler(err error) {
	fmt.Printf("(panic) error handled while creating instance: %v", err)
	panic(err)
}

func resolveAlbumService() controller.Service {
	serv, err := service.NewAlbumService(
		resolveAlbumRepository(),
	)
	if err != nil {
		panicHandler(err)
	}
	return serv
}

func resolveAlbumRepository() service.AlbumRepo {
	repo, err := repository.NewAlbumRepository(
		&http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		},
	)
	if err != nil {
		panicHandler(err)
	}
	return repo
}
