package movies

import (
	"errors"

	"github.com/ValeHenriquez/neo4j-imdb-go/database"
	"github.com/ValeHenriquez/neo4j-imdb-go/models"
	"github.com/ValeHenriquez/neo4j-imdb-go/utils"
)

func getRandomMovie() (models.MovieResponse, error) {
	var movieResponse models.MovieResponse

	dataMovie, err := database.GetRandom(models.Movie{})
	if err != nil {
		return models.MovieResponse{}, err
	}

	movie, genres, actors, err := getMovieData(dataMovie)
	if err != nil {
		return models.MovieResponse{}, err
	}

	movieResponse = models.MovieResponse{
		Id:           movie.Id,
		Title:        movie.Title,
		Overview:     movie.Overview,
		Director:     movie.Director,
		BackdropPath: movie.BackdropPath,
		Runtime:      movie.Runtime,
		PosterPath:   movie.PosterPath,
		ReleaseDate:  movie.ReleaseDate,
		Genres:       genres,
		Actors:       actors,
	}

	return movieResponse, nil
}

func getMovieData(movieData map[string]interface{}) (models.Movie, []models.Genre, []models.Actor, error) {
	movie := models.Movie{
		Id:           movieData["Id"].(int64),
		Title:        movieData["Title"].(string),
		Overview:     movieData["Overview"].(string),
		Director:     movieData["Director"].(string),
		BackdropPath: movieData["BackdropPath"].(string),
		Runtime:      movieData["Runtime"].(int64),
		PosterPath:   movieData["PosterPath"].(string),
		ReleaseDate:  movieData["ReleaseDate"].(string),
	}

	genresData, err := database.GetRelationships(movie, utils.CATEGORIZED)
	if err != nil {
		return models.Movie{}, nil, nil, err
	}
	genres := make([]models.Genre, len(genresData))
	for i, g := range genresData {
		genres[i] = models.Genre{
			Id:   g["Id"].(int64),
			Name: g["Name"].(string),
		}
	}

	actorsData, err := database.GetRelationships(movie, utils.ACTED_IN)
	if err != nil {
		return models.Movie{}, nil, nil, err
	}
	actors := make([]models.Actor, len(actorsData))
	for i, a := range actorsData {
		actors[i] = models.Actor{
			Id:        a["Id"].(int64),
			Name:      a["Name"].(string),
			Character: a["Character"].(string),
		}
	}

	return movie, genres, actors, nil
}

func getAllMovies() ([]models.MovieResponse, error) {
	var movies []models.MovieResponse

	dataMovies, err := database.GetAll(models.Movie{})
	if err != nil {
		return nil, err
	}

	for _, d := range dataMovies {
		movie, genres, actors, err := getMovieData(d)
		if err != nil {
			return nil, err
		}

		movieResponse := models.MovieResponse{
			Id:           movie.Id,
			Title:        movie.Title,
			Overview:     movie.Overview,
			Director:     movie.Director,
			BackdropPath: movie.BackdropPath,
			Runtime:      movie.Runtime,
			PosterPath:   movie.PosterPath,
			ReleaseDate:  movie.ReleaseDate,
			Genres:       genres,
			Actors:       actors,
		}

		movies = append(movies, movieResponse)
	}

	return movies, nil
}

func getMovieById(id int64) (models.MovieResponse, error) {
	var movieResponse models.MovieResponse

	movieData, err := database.GetOne(models.Movie{Id: id})
	if err != nil || movieData == nil {
		return models.MovieResponse{}, errors.New("movie not found")
	}

	movie, genres, actors, err := getMovieData(movieData)
	if err != nil {
		return models.MovieResponse{}, err
	}

	movieResponse = models.MovieResponse{
		Id:           movie.Id,
		Title:        movie.Title,
		Overview:     movie.Overview,
		Director:     movie.Director,
		BackdropPath: movie.BackdropPath,
		Runtime:      movie.Runtime,
		PosterPath:   movie.PosterPath,
		ReleaseDate:  movie.ReleaseDate,
		Genres:       genres,
		Actors:       actors,
	}

	return movieResponse, nil
}

func getMovieRecomendations(id int64) ([]models.MovieResponse, error) {
	var movies []models.MovieResponse

	movieData, err := database.GetOne(models.Movie{Id: id})
	if err != nil || movieData == nil {
		return nil, errors.New("movie not found")
	}

	movie, _, _, err := getMovieData(movieData)
	if err != nil {
		return nil, err
	}

	recomendationsData, err := database.GetRelationships(movie, utils.RECOMMENDS)
	if err != nil {
		return nil, err
	}

	for _, d := range recomendationsData {
		movie, genres, actors, err := getMovieData(d)
		if err != nil {
			return nil, err
		}

		movieResponse := models.MovieResponse{
			Id:           movie.Id,
			Title:        movie.Title,
			Overview:     movie.Overview,
			Director:     movie.Director,
			BackdropPath: movie.BackdropPath,
			Runtime:      movie.Runtime,
			PosterPath:   movie.PosterPath,
			ReleaseDate:  movie.ReleaseDate,
			Genres:       genres,
			Actors:       actors,
		}

		movies = append(movies, movieResponse)
	}

	if len(movies) == 0 {
		return []models.MovieResponse{}, nil
	}

	return movies, nil
}
