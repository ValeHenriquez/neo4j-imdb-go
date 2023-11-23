package api

type PopularMoviesResponseResult struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIDs         []int   `json:"genre_ids"`
	Id               int64   `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int64   `json:"vote_count"`
}

type PopularMoviesResponse struct {
	Page         int64                         `json:"page"`
	Results      []PopularMoviesResponseResult `json:"results"`
	TotalResults int64                         `json:"total_results"`
	TotalPages   int64                         `json:"total_pages"`
}

type MovieResponse struct {
	Adult               bool             `json:"adult"`
	BackdropPath        string           `json:"backdrop_path"`
	BelongsToCollection BelongsTo        `json:"belongs_to_collection"`
	Budget              int64            `json:"budget"`
	Genres              []Genre          `json:"genres"`
	Homepage            string           `json:"homepage"`
	Id                  int64            `json:"id"`
	ImdbID              string           `json:"imdb_id"`
	OriginalLanguage    string           `json:"original_language"`
	OriginalTitle       string           `json:"original_title"`
	Overview            string           `json:"overview"`
	Popularity          float64          `json:"popularity"`
	PosterPath          string           `json:"poster_path"`
	ProductionCompanies []ProductionComp `json:"production_companies"`
	ProductionCountries []Country        `json:"production_countries"`
	ReleaseDate         string           `json:"release_date"`
	Revenue             int64            `json:"revenue"`
	Runtime             int              `json:"runtime"`
	SpokenLanguages     []Language       `json:"spoken_languages"`
	Status              string           `json:"status"`
	Tagline             string           `json:"tagline"`
	Title               string           `json:"title"`
	Video               bool             `json:"video"`
	VoteAverage         float64          `json:"vote_average"`
	VoteCount           int64            `json:"vote_count"`
	Credits             Credits          `json:"credits"`
}

type MovieRecommendationsResponse struct {
	Page    int64 `json:"page"`
	Results []MovieRecommendationsResponseResults `json:"results"`
}

type MovieRecommendationsResponseResults struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	Id               int64   `json:"id"`
	Title            string  `json:"title"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	PosterPath       string  `json:"poster_path"`
	MediaType        string  `json:"media_type"`
	GenreIds         []int64 `json:"genre_ids"`
	Popularity       float64 `json:"popularity"`
	ReleaseDate      string  `json:"release_date"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int64   `json:"vote_count"`
}

type Genre struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type BelongsTo struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
}

type ProductionComp struct {
	Id            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

type Country struct {
	Iso31661 string `json:"iso_3166_1"`
	Name     string `json:"name"`
}

type Language struct {
	EnglishName string `json:"english_name"`
	Iso6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
}

type Credits struct {
	Cast []Cast `json:"cast"`
	Crew []Crew `json:"crew"`
}

type Cast struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	Id                 int64   `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	CastId             int64   `json:"cast_id"`
	Character          string  `json:"character"`
	CreditId           string  `json:"credit_id"`
	Order              int     `json:"order"`
}

type Crew struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	Id                 int64   `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	CreditId           string  `json:"credit_id"`
	Department         string  `json:"department"`
	Job                string  `json:"job"`
}
