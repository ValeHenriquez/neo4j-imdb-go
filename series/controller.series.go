package series

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetSeries(c *fiber.Ctx) error {
	fmt.Println("Getting Series")
	series, err := getAllSeries()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(series)
}

func GetSerie(c *fiber.Ctx) error {
	fmt.Println("Getting Serie")
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	serie, err := getSerieById(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Serie not found"})
	}
	return c.Status(http.StatusOK).JSON(serie)
}

func GetSerieRecomendations(c *fiber.Ctx) error {
    fmt.Println("Getting Serie Recomendations")
    id, err := strconv.ParseInt(c.Params("id"), 10, 64)
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    serie, err := getSerieRecomendations(id)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(http.StatusOK).JSON(serie)
}