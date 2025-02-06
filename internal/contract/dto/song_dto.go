package dto

type (
	GetSongsDto struct {
		Page     int `query:"page" form:"page"`
		PageSize int `query:"pageSize" form:"pageSize"`

		Song        string `json:"song" form:"song"`
		Group       string `json:"group" form:"group"`
		ReleaseDate string `json:"releaseDate" form:"releaseDate"`
		Text        string `json:"text" form:"text"`
		Link        string `json:"link" form:"link"`
	}

	PostSongDto struct {
		Song  string `json:"song" form:"song"`
		Group string `json:"group" form:"group"`
	}

	SongDto struct {
		SongId      string `json:"songId" form:"songId"`
		Song        string `json:"song" form:"song"`
		Group       string `json:"group" form:"group"`
		ReleaseDate string `json:"releaseDate" form:"releaseDate"`
		Text        string `json:"text" form:"text"`
		Link        string `json:"link" form:"link"`
	}

	GetSongsResponseDto struct {
		Page     int       `json:"page"`
		PageSize int       `json:"pageSize"`
		Items    []SongDto `json:"items"`
	}

	PutSongDto struct {
		Song        string `json:"song" form:"song"`
		Group       string `json:"group" form:"group"`
		ReleaseDate string `json:"releaseDate" form:"releaseDate"`
		Text        string `json:"text" form:"text"`
		Link        string `json:"link" form:"link"`
	}
)
