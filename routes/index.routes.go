package routes

import (
	"github.com/ValeHenriquez/neo4j-imdb-go/movies"
	api "github.com/ValeHenriquez/neo4j-imdb-go/tmdb_api"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Server Running 👋!")
	})

	app.Get("/movies", movies.GetMovies)
	app.Get("/movies/:id", movies.GetMovie)
	app.Get("/movies/:id/recomendations", movies.GetMovieRecomendations)

	//Rutas auxiliares para poblar y eliminar la base de datos
	app.Post("/populateDB", api.PopulateDB)
	app.Delete("/delete", api.DeleteDB)

}
