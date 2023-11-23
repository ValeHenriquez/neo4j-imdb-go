package models

type Movie struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	Overview     string `json:"overview"`
	Director     string `json:"director"`
	BackdropPath string `json:"backdrop_path"`
	PosterPath   string `json:"poster_path"`
	ReleaseDate  string `json:"release_date"`
}

type MovieResponse struct {
	Id           int64   `json:"id"`
	Title        string  `json:"title"`
	Overview     string  `json:"overview"`
	Director     string  `json:"director"`
	BackdropPath string  `json:"backdrop_path"`
	PosterPath   string  `json:"poster_path"`
	ReleaseDate  string  `json:"release_date"`
	Genres       []Genre `json:"genres"`
	Actors       []Actor `json:"actors"`
}
