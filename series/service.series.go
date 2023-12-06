package series

import (
	"errors"

	"github.com/ValeHenriquez/neo4j-imdb-go/database"
	"github.com/ValeHenriquez/neo4j-imdb-go/models"
	"github.com/ValeHenriquez/neo4j-imdb-go/utils"
)

func getSerieData(serieData map[string]interface{}) (models.Serie, []models.Genre, []models.Actor, error) {
	serie := models.Serie{
		Id:               serieData["Id"].(int64),
		Name:             serieData["Name"].(string),
		Overview:         serieData["Overview"].(string),
		CreatedBy:        serieData["CreatedBy"].(string),
		BackdropPath:     serieData["BackdropPath"].(string),
		PosterPath:       serieData["PosterPath"].(string),
		FirstAirDate:     serieData["FirstAirDate"].(string),
		NumberOfEpisodes: serieData["NumberOfEpisodes"].(int64),
		NumberOfSeasons:  serieData["NumberOfSeasons"].(int64),
	}

	genresData, err := database.GetRelationships(serie, utils.CATEGORIZED)
	if err != nil {
		return models.Serie{}, nil, nil, err
	}
	genres := make([]models.Genre, len(genresData))
	for i, g := range genresData {
		genres[i] = models.Genre{
			Id:   g["Id"].(int64),
			Name: g["Name"].(string),
		}
	}

	actorsData, err := database.GetRelationships(serie, utils.ACTED_IN)
	if err != nil {
		return models.Serie{}, nil, nil, err
	}
	actors := make([]models.Actor, len(actorsData))
	for i, a := range actorsData {
		actors[i] = models.Actor{
			Id:        a["Id"].(int64),
			Name:      a["Name"].(string),
			Character: a["Character"].(string),
		}
	}

	return serie, genres, actors, nil
}

func getAllSeries() ([]models.SerieResponse, error) {
	var series []models.SerieResponse

	dataSeries, err := database.GetAll(models.Serie{})
	if err != nil {
		return nil, err
	}

	for _, d := range dataSeries {
		serie, genres, actors, err := getSerieData(d)
		if err != nil {
			return nil, err
		}
		series = append(series, models.SerieResponse{
			Id:               serie.Id,
			Name:             serie.Name,
			Overview:         serie.Overview,
			CreatedBy:        serie.CreatedBy,
			BackdropPath:     serie.BackdropPath,
			PosterPath:       serie.PosterPath,
			FirstAirDate:     serie.FirstAirDate,
			NumberOfEpisodes: serie.NumberOfEpisodes,
			NumberOfSeasons:  serie.NumberOfSeasons,
			Genres:           genres,
			Actors:           actors,
		})
	}

	return series, nil
}

func getSerieById(id int64) (models.SerieResponse, error) {
	serieData, err := database.GetOne(models.Serie{Id: id})
	if err != nil {
		return models.SerieResponse{}, errors.New("serie not found")
	}

	serie, genres, actors, err := getSerieData(serieData)
	if err != nil {
		return models.SerieResponse{}, err
	}

	return models.SerieResponse{
		Id:               serie.Id,
		Name:             serie.Name,
		Overview:         serie.Overview,
		CreatedBy:        serie.CreatedBy,
		BackdropPath:     serie.BackdropPath,
		PosterPath:       serie.PosterPath,
		FirstAirDate:     serie.FirstAirDate,
		NumberOfEpisodes: serie.NumberOfEpisodes,
		NumberOfSeasons:  serie.NumberOfSeasons,
		Genres:           genres,
		Actors:           actors,
	}, nil
}

func getSerieRecomendations(id int64) ([]models.SerieResponse, error) {
    var series []models.SerieResponse

    serieData, err := database.GetOne(models.Serie{Id: id})
    if err != nil {
        return nil, errors.New("serie not found")
    }

    serie, _, _, err := getSerieData(serieData)
    if err != nil {
        return nil, err
    }

    serieRecomendations, err := database.GetRelationships(serie, utils.RECOMMENDS)
    if err != nil {
        return nil, err
    }

    for _, d := range serieRecomendations {
        serie, genres, actors, err := getSerieData(d)
        if err != nil {
            return nil, err
        }
        series = append(series, models.SerieResponse{
            Id:               serie.Id,
            Name:             serie.Name,
            Overview:         serie.Overview,
            CreatedBy:        serie.CreatedBy,
            BackdropPath:     serie.BackdropPath,
            PosterPath:       serie.PosterPath,
            FirstAirDate:     serie.FirstAirDate,
            NumberOfEpisodes: serie.NumberOfEpisodes,
            NumberOfSeasons:  serie.NumberOfSeasons,
            Genres:           genres,
            Actors:           actors,
        })
    }

    return series, nil
}