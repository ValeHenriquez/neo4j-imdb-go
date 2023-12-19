package api

type CreatedBy struct {
	Id          int64  `json:"id"`
	CreditId    string `json:"credit_id"`
	Name        string `json:"name"`
	Gender      int64  `json:"gender"`
	ProfilePath string `json:"profile_path"`
}

type Network struct {
	Id            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

type Season struct {
	AirDate      string  `json:"air_date"`
	EpisodeCount int64   `json:"episode_count"`
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   string  `json:"poster_path"`
	SeasonNumber int64   `json:"season_number"`
	VoteAverage  float64 `json:"vote_average"`
}

type ProductionCompany struct {
	Id            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

type ProductionCountry struct {
	ISO31661 string `json:"iso_3166_1"`
	Name     string `json:"name"`
}

type SpokenLanguage struct {
	EnglishName string `json:"english_name"`
	ISO6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
}

type SerieResponse struct {
	Adult               bool                `json:"adult"`
	BackdropPath        string              `json:"backdrop_path"`
	CreatedBy           []CreatedBy         `json:"created_by"`
	EpisodeRunTime      []int64             `json:"episode_run_time"`
	FirstAirDate        string              `json:"first_air_date"`
	Genres              []Genre             `json:"genres"`
	Homepage            string              `json:"homepage"`
	Id                  int64               `json:"id"`
	InProduction        bool                `json:"in_production"`
	Languages           []string            `json:"languages"`
	LastAirDate         *string             `json:"last_air_date"`
	LastEpisodeToAir    *interface{}        `json:"last_episode_to_air"`
	Name                string              `json:"name"`
	NextEpisodeToAir    *interface{}        `json:"next_episode_to_air"`
	Networks            []Network           `json:"networks"`
	NumberOfEpisodes    int64               `json:"number_of_episodes"`
	NumberOfSeasons     int64               `json:"number_of_seasons"`
	OriginCountry       []string            `json:"origin_country"`
	OriginalLanguage    string              `json:"original_language"`
	OriginalName        string              `json:"original_name"`
	Overview            string              `json:"overview"`
	Popularity          float64             `json:"popularity"`
	PosterPath          string              `json:"poster_path"`
	ProductionCompanies []ProductionCompany `json:"production_companies"`
	ProductionCountries []ProductionCountry `json:"production_countries"`
	Seasons             []Season            `json:"seasons"`
	SpokenLanguages     []SpokenLanguage    `json:"spoken_languages"`
	Status              string              `json:"status"`
	Tagline             string              `json:"tagline"`
	Type                string              `json:"type"`
	VoteAverage         float64             `json:"vote_average"`
	VoteCount           int64               `json:"vote_count"`
	Credits             Credits             `json:"credits"`
}

type PopularMoviesResponseResult struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIds         []int64 `json:"genre_ids"`
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

type PopularSeriesResponse struct {
	Page         int64                         `json:"page"`
	Results      []PopularSeriesResponseResult `json:"results"`
	TotalPages   int64                         `json:"total_pages"`
	TotalResults int64                         `json:"total_results"`
}

type PopularSeriesResponseResult struct {
	Adult            bool     `json:"adult"`
	BackdropPath     string   `json:"backdrop_path"`
	GenreIds         []int64  `json:"genre_ids"`
	Id               int64    `json:"id"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	FirstAirDate     string   `json:"first_air_date"`
	Name             string   `json:"name"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int64    `json:"vote_count"`
}

type MovieResponse struct {
	Adult               bool             `json:"adult"`
	BackdropPath        string           `json:"backdrop_path"`
	BelongsToCollection BelongsTo        `json:"belongs_to_collection"`
	Budget              int64            `json:"budget"`
	Genres              []Genre          `json:"genres"`
	Homepage            string           `json:"homepage"`
	Id                  int64            `json:"id"`
	ImdbId              string           `json:"imdb_id"`
	OriginalLanguage    string           `json:"original_language"`
	OriginalTitle       string           `json:"original_title"`
	Overview            string           `json:"overview"`
	Popularity          float64          `json:"popularity"`
	PosterPath          string           `json:"poster_path"`
	ProductionCompanies []ProductionComp `json:"production_companies"`
	ProductionCountries []Country        `json:"production_countries"`
	ReleaseDate         string           `json:"release_date"`
	Revenue             int64            `json:"revenue"`
	Runtime             int64            `json:"runtime"`
	SpokenLanguages     []Language       `json:"spoken_languages"`
	Status              string           `json:"status"`
	Tagline             string           `json:"tagline"`
	Title               string           `json:"title"`
	Video               bool             `json:"video"`
	VoteAverage         float64          `json:"vote_average"`
	VoteCount           int64            `json:"vote_count"`
	Credits             Credits          `json:"credits"`
}

type SerieRecommendationsResponseResults struct {
	Adult            bool     `json:"adult"`
	BackdropPath     string   `json:"backdrop_path"`
	Id               int64    `json:"id"`
	Name             string   `json:"name"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	PosterPath       string   `json:"poster_path"`
	MediaType        string   `json:"media_type"`
	GenreIds         []int64  `json:"genre_ids"`
	Popularity       float64  `json:"popularity"`
	FirstAirDate     string   `json:"first_air_date"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int64    `json:"vote_count"`
	OriginCountry    []string `json:"origin_country"`
}

type SerieRecommendationsResponse struct {
	Page         int64                                 `json:"page"`
	Results      []SerieRecommendationsResponseResults `json:"results"`
	TotalPages   int64                                 `json:"total_pages"`
	TotalResults int64                                 `json:"total_results"`
}

type MovieRecommendationsResponse struct {
	Page    int64                                 `json:"page"`
	Results []MovieRecommendationsResponseResults `json:"results"`
}

type MovieRecommendationsResponseResults struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIds         []int64 `json:"genre_ids"`
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
	Gender             int64   `json:"gender"`
	Id                 int64   `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	CastId             int64   `json:"cast_id"`
	Character          string  `json:"character"`
	CreditId           string  `json:"credit_id"`
	Order              int64   `json:"order"`
}

type Crew struct {
	Adult              bool    `json:"adult"`
	Gender             int64   `json:"gender"`
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
