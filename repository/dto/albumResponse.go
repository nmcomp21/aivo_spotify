package dto

import (
	"github.com/nmcomp/aivo-spotify/domain"
)

type (
	AlbumResponse struct {
		Items []AlbumItemResponse `json:"items"`
	}
	AlbumItemResponse struct {
		Image       []AlbumImageResponse `json:"images"`
		Name        string               `json:"name"`
		ReleaseDate string               `json:"release_date"`
		TotalTracks uint                 `json:"total_tracks"`
	}
	AlbumImageResponse struct {
		Height uint   `json:"height"`
		Width  uint   `json:"width"`
		Url    string `json:"url"`
	}
)

func (a AlbumResponse) ToAlbum() []domain.Album {
	var albums []domain.Album
	for _, item := range a.Items {
		album := domain.Album{
			Name:        item.Name,
			ReleaseDate: item.ReleaseDate,
			TracksQty:   item.TotalTracks,
		}
		covers := []domain.Cover{}
		for _, cover := range item.Image {
			albumCover := domain.Cover{
				Height: cover.Height,
				Width:  cover.Width,
				Url:    cover.Url,
			}
			covers = append(covers, albumCover)
		}
		album.Covers = covers
		albums = append(albums, album)
	}
	return albums
}
