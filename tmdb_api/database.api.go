package api

import (
	"fmt"
	"sync"

	"github.com/ValeHenriquez/neo4j-imdb-go/database"
	"github.com/ValeHenriquez/neo4j-imdb-go/models"
	"github.com/ValeHenriquez/neo4j-imdb-go/utils"
)

func fillDB() error {
	var wg sync.WaitGroup
	wg.Add(2) 

	go func() {
		defer wg.Done()
		if err := fillDBCategory(getPopularMoviesIds, fillDBWithRecommendationsMovie); err != nil {
			fmt.Println("Error filling movies:", err)
		}
		fmt.Println("MOVIES CARGADAS")
	}()

	go func() {
		defer wg.Done()
		if err := fillDBCategory(getPopularSeriesIds, fillDBWithRecommendationsSerie); err != nil {
			fmt.Println("Error filling series:", err)
		}
		fmt.Println("SERIES CARGADAS")
	}()

	wg.Wait() 
	return nil
}

func fillDBCategory(getIDs func() ([]int64, error), fillFunc func(int64) error) error {
	ids, err := getIDs()
	if err != nil {
		return err
	}

	for _, id := range ids {
		if err := fillFunc(id); err != nil {
			return err
		}
	}

	return nil
}

func fillDBWithRecommendationsSerie(serieID int64) error {
	mainSerie, mainActors, mainGenres, err := getSeriesDetailsById(serieID)
	if err != nil {
		return err
	}

	if err := fillDBCategory(func() ([]int64, error) { return getSerieRecommendationsIds(serieID) }, func(id int64) error {
		serie, actors, genres, err := getSeriesDetailsById(id)
		if err != nil {
			return err
		}
		if err := fillDBSerie(serie, actors, genres); err != nil {
			return err
		}
		return database.CreateBiRelation(mainSerie, serie, utils.RECOMMENDS)
	}); err != nil {
		return err
	}

	return fillDBSerie(mainSerie, mainActors, mainGenres)
}

func fillDBSerie(serie models.Serie, actors []models.Actor, genres []models.Genre) error {
	if err := createCategoryAndRelations(&serie, actors, genres, utils.ACTED_IN, utils.CATEGORIZED); err != nil {
		return err
	}
	return nil
}

func fillDBMovie(movie models.Movie, actors []models.Actor, genres []models.Genre) error {
	if err := createCategoryAndRelations(&movie, actors, genres, utils.ACTED_IN, utils.CATEGORIZED); err != nil {
		return err
	}
	return nil
}

func createCategoryAndRelations(category interface{}, actors []models.Actor, genres []models.Genre, actorRelation, genreRelation string) error {
	if err := database.Create(category); err != nil {
		return err
	}

	for _, actor := range actors {
		if err := database.Create(&actor); err != nil {
			return err
		}
		if err := database.CreateRelation(category, &actor, actorRelation); err != nil {
			return err
		}
	}

	for _, genre := range genres {
		if err := database.Create(&genre); err != nil {
			return err
		}
		if err := database.CreateRelation(category, &genre, genreRelation); err != nil {
			return err
		}
	}

	return nil
}

func fillDBWithRecommendationsMovie(movieID int64) error {
	mainMovie, mainActors, mainGenres, err := getMovieDetailsById(movieID)
	if err != nil {
		return err
	}

	if err := fillDBCategory(func() ([]int64, error) { return getMovieRecommendationsIds(movieID) }, func(id int64) error {
		movie, actors, genres, err := getMovieDetailsById(id)
		if err != nil {
			return err
		}
		if err := fillDBMovie(movie, actors, genres); err != nil {
			return err
		}
		return database.CreateBiRelation(mainMovie, movie, utils.RECOMMENDS)
	}); err != nil {
		return err
	}

	return fillDBMovie(mainMovie, mainActors, mainGenres)
}
