package models

type Serie struct {
	Id               int64  `json:"id"`
	Name             string `json:"title"`
	Overview         string `json:"overview"`
	CreatedBy        string `json:"created_by"`
	BackdropPath     string `json:"backdrop_path"`
	PosterPath       string `json:"poster_path"`
	FirstAirDate     string `json:"first_air_date"`
	NumberOfEpisodes int64  `json:"number_of_episodes"`
	NumberOfSeasons  int64  `json:"number_of_seasons"`
}

type SerieResponse struct {
	Id               int64   `json:"id"`
	Name             string  `json:"title"`
	Overview         string  `json:"overview"`
	CreatedBy        string  `json:"created_by"`
	BackdropPath     string  `json:"backdrop_path"`
	PosterPath       string  `json:"poster_path"`
	FirstAirDate     string  `json:"first_air_date"`
	NumberOfEpisodes int64   `json:"number_of_episodes"`
	NumberOfSeasons  int64   `json:"number_of_seasons"`
	Genres           []Genre `json:"genres"`
	Actors           []Actor `json:"actors"`
}
