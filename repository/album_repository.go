package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nmcomp/aivo-spotify/domain"
	"github.com/nmcomp/aivo-spotify/repository/dto"
)

type (
	AlbumRepository struct {
		client *http.Client
	}
)

const (
	albumByBandURL  = "https://api.spotify.com/v1/artists/%s/albums"
	searchArtistURL = "https://api.spotify.com/v1/search?q=%s"
	bearerToken     = "BQCHexrSTbp2KZJEAHLHDz2CNfb4QT9mHUaonrGpgDOV-pjzox_kbxLf-NOJFUwBGHzFFquZxL8pz_1BG4BoBahdOvnZrCoEKEe-Jf9dhx8nhRBqz1CkGoA7WfBg-KTFAAis-8C4-cU_JG-ILbHOItAqEWN-neHtUqY"
)

func NewAlbumRepository(client *http.Client) (*AlbumRepository, error) {
	r := &AlbumRepository{
		client: client,
	}
	if err := r.validate(); err != nil {
		return nil, fmt.Errorf("invalid album repository: %v", err)
	}

	return r, nil
}

func (r *AlbumRepository) validate() error {
	if r.client == nil {
		return errors.New("rest client should not be nil")
	}
	return nil
}

func (r *AlbumRepository) getArtistIDByArtistName(bandName string) (string, error) {
	client := &http.Client{}
	url := fmt.Sprintf(searchArtistURL, bandName+"&type=artist")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Bearer "+bearerToken)
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var artistResponse dto.SearchArtistResponse
	err = json.Unmarshal(responseData, &artistResponse)
	if err != nil {
		return "", errors.New("error unmarshalling response")
	}
	if len(artistResponse.Artists.Items) == 0 {
		return "", nil
	}
	return artistResponse.Artists.Items[0].ArtistID, nil
}

func (r *AlbumRepository) GetAlbumsByBandName(c *gin.Context, bandName string) ([]domain.Album, error) {
	artistID, err := r.getArtistIDByArtistName(bandName)
	if err != nil {
		return nil, err
	}
	if artistID == "" {
		return nil, errors.New("no artist found")
	}
	client := &http.Client{}
	url := fmt.Sprintf(albumByBandURL, artistID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+bearerToken)
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var albumResponse dto.AlbumResponse
	err = json.Unmarshal(responseData, &albumResponse)
	if err != nil {
		return nil, errors.New("error unmarshalling response")
	}

	return albumResponse.ToAlbum(), nil
}
