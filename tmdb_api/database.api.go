package api

import (
	"github.com/ValeHenriquez/neo4j-imdb-go/database"
	"github.com/ValeHenriquez/neo4j-imdb-go/models"
	"github.com/ValeHenriquez/neo4j-imdb-go/utils"
)

func FillDB() error {
	idsPopular, err := getPopularMoviesIds()
	if err != nil {
		return err
	}

	for _, id := range idsPopular {
		if err := fillDBWithRecommendations(id); err != nil {
			return err
		}
	}

	return nil
}

func fillDBMovie(movie models.Movie, actors []models.Actor, genres []models.Genre) error {
	if err := createMovieAndRelations(movie, actors, utils.ACTED_IN); err != nil {
		return err
	}

	for _, genre := range genres {
		if err := createGenreAndRelation(movie, genre, utils.CATEGORIZED); err != nil {
			return err
		}
	}

	return nil
}

func createMovieAndRelations(movie models.Movie, actors []models.Actor, relationType string) error {
	if err := database.Create(&movie); err != nil {
		return err
	}

	for _, actor := range actors {
		if err := database.Create(&actor); err != nil {
			return err
		}
		if err := database.CreateRelation(&movie, &actor, relationType); err != nil {
			return err
		}
	}

	return nil
}

func createGenreAndRelation(movie models.Movie, genre models.Genre, relationType string) error {
	if err := database.Create(&genre); err != nil {
		return err
	}
	return database.CreateRelation(&movie, &genre, relationType)
}

func fillDBWithRecommendations(movieID int64) error {
	mainMovie, mainActors, mainGenres, err := getMovieDetailsById(movieID)
	if err != nil {
		return err
	}

	if err := fillDBMovie(mainMovie, mainActors, mainGenres); err != nil {
		return err
	}

	idsRecommendations, err := getMovieRecommendationsIds(movieID)
	if err != nil {
		return err
	}

	for _, id := range idsRecommendations {
		movie, actors, genres, err := getMovieDetailsById(id)
		if err != nil {
			return err
		}
		if err := fillDBMovie(movie, actors, genres); err != nil {
			return err
		}

		if err := database.CreateBiRelation(mainMovie, movie, utils.RECOMMENDS); err != nil {
			return err
		}
	}

	return nil
}
