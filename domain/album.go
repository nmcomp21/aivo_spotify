package domain

type (
	Album struct {
		Name        string  `json:"name"`
		ReleaseDate string  `json:"released"`
		TracksQty   uint    `json:"tracks"`
		Covers      []Cover `json:"covers"`
	}
	Cover struct {
		Height uint   `json:"height"`
		Width  uint   `json:"width"`
		Url    string `json:"url"`
	}
)
