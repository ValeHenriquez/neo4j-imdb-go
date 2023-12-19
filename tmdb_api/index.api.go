package api

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/ValeHenriquez/neo4j-imdb-go/models"
	"github.com/ValeHenriquez/neo4j-imdb-go/utils"
)

var CANT_ACTORS = 5
var CANT_PAGES = 1

func getSeriesDetailsById(id int64) (models.Serie, []models.Actor, []models.Genre, error) {
	addedURL := "/tv/" + strconv.FormatInt(id, 10) + "?append_to_response=credits&language=en-US"

	res, err := utils.MakeRequest(addedURL)
	if err != nil {
		return models.Serie{}, nil, nil, err
	}
	defer res.Body.Close()

	var data SerieResponse
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return models.Serie{}, nil, nil, err
	}

	createdBy := findCreatedBy(data.CreatedBy)
	actors := getActors(data.Credits.Cast, CANT_ACTORS)
	genres := getGenres(data.Genres)

	series := models.Serie{
		Id:               data.Id,
		Name:             data.Name,
		Overview:         data.Overview,
		CreatedBy:        createdBy,
		BackdropPath:     data.BackdropPath,
		PosterPath:       data.PosterPath,
		FirstAirDate:     data.FirstAirDate,
		NumberOfEpisodes: data.NumberOfEpisodes,
		NumberOfSeasons:  data.NumberOfSeasons,
	}

	return series, actors, genres, nil
}

func findCreatedBy(createdBy []CreatedBy) string {
	for _, creator := range createdBy {
		return creator.Name
	}
	return ""
}

func getPopularSeriesIds() ([]int64, error) {
	addedURL := fmt.Sprintf("/tv/popular?language=en-US&page=%d", CANT_PAGES)
	res, err := utils.MakeRequest(addedURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data PopularSeriesResponse
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	ids := make([]int64, len(data.Results))
	for i, movie := range data.Results {
		ids[i] = movie.Id
	}

	return ids, nil
}

func getMovieDetailsById(id int64) (models.Movie, []models.Actor, []models.Genre, error) {
	addedURL := "/movie/" + strconv.FormatInt(id, 10) + "?append_to_response=credits&language=en-US"

	res, err := utils.MakeRequest(addedURL)
	if err != nil {
		return models.Movie{}, nil, nil, err
	}
	defer res.Body.Close()

	var data MovieResponse
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return models.Movie{}, nil, nil, err
	}

	director := findCrewMember(data.Credits.Crew, "Director")
	actors := getActors(data.Credits.Cast, CANT_ACTORS)
	genres := getGenres(data.Genres)

	movie := models.Movie{
		Id:           data.Id,
		Title:        data.Title,
		Overview:     data.Overview,
		Director:     director,
		BackdropPath: data.BackdropPath,
		Runtime:      data.Runtime,
		PosterPath:   data.PosterPath,
		ReleaseDate:  data.ReleaseDate,
	}

	return movie, actors, genres, nil
}

func getGenres(genres []Genre) []models.Genre {
	var genresResponse []models.Genre
	for _, genre := range genres {
		genresResponse = append(genresResponse, models.Genre{Id: genre.Id, Name: genre.Name})
	}
	return genresResponse
}

func getActors(cast []Cast, limit int) []models.Actor {
	var actors []models.Actor
	sort.Slice(cast, func(i, j int) bool {
		return cast[i].Order < cast[j].Order
	})

	for i, actor := range cast {
		if i == len(cast) || i == limit {
			break
		}
		actors = append(actors, models.Actor{Id: actor.Id, Name: actor.Name, Character: actor.Character})
	}
	return actors
}

func findCrewMember(crew []Crew, job string) string {
	for _, crewMember := range crew {
		if crewMember.Job == job {
			return crewMember.Name
		}
	}
	return ""
}

func getPopularMoviesIds() ([]int64, error) {
	addedURL := fmt.Sprintf("/movie/popular?language=en-US&page=%d", CANT_PAGES)
	res, err := utils.MakeRequest(addedURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data PopularMoviesResponse
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	ids := make([]int64, len(data.Results))
	for i, movie := range data.Results {
		ids[i] = movie.Id
	}

	return ids, nil
}

func getMovieRecommendationsIds(id int64) ([]int64, error) {
	addedURL := "/movie/" + strconv.FormatInt(id, 10) + "/similar?language=en-US&page=1"

	res, err := utils.MakeRequest(addedURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data MovieRecommendationsResponse

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	ids := make([]int64, len(data.Results))
	for i, movie := range data.Results {
		ids[i] = movie.Id
	}

	fmt.Println("IDS from get movie recommendations", ids)
	return ids, nil
}

func getSerieRecommendationsIds(id int64) ([]int64, error) {
	addedURL := "/tv/" + strconv.FormatInt(id, 10) + "/similar?language=en-US&page=1"

	res, err := utils.MakeRequest(addedURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data SerieRecommendationsResponse
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	ids := make([]int64, len(data.Results))
	for i, serie := range data.Results {
		ids[i] = serie.Id
	}

	return ids, nil
}
