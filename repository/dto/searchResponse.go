package dto

type (
	SearchArtistResponse struct {
		Artists ArtistResponse `json:"artists"`
	}
	ArtistResponse struct {
		Items []ItemResponse `json:"items"`
	}
	ItemResponse struct {
		ArtistID string `json:"id"`
	}
)
