package movies

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetRandom(c *fiber.Ctx) error {
	fmt.Println("Getting Random Movie")
	movie, err := getRandomMovie()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(movie)
}

func GetMovies(c *fiber.Ctx) error {
	fmt.Println("Getting Movies")
	movies, err := getAllMovies()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(movies)
}

func GetMovie(c *fiber.Ctx) error {
	fmt.Println("Getting Movie")
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	movie, err := getMovieById(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Movie not found"})
	}
	return c.Status(http.StatusOK).JSON(movie)
}

func GetMovieRecomendations(c *fiber.Ctx) error {
	fmt.Println("Getting Movie Recomendations")
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	movies, err := getMovieRecomendations(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(movies)
}
