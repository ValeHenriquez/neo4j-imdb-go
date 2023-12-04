package api

import (
	"fmt"
	"time"

	"github.com/ValeHenriquez/neo4j-imdb-go/database"
	"github.com/gofiber/fiber/v2"
)

func PopulateDB(c *fiber.Ctx) error {
	fmt.Println("Comenzando el llenado de la base de datos...")
	startTime := time.Now()

	err := fillDB()

	if err != nil {
		return c.SendString(err.Error())
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Println("Tiempo transcurrido:", elapsedTime)

	fmt.Println("¡Llenado de la base de datos completado!")
	return c.SendString("¡Llenado de la base de datos completado! Tiempo transcurrido: " + elapsedTime.String())
}

func DeleteDB(c *fiber.Ctx) error {
	fmt.Println("Borrando todos los nodos de la base de datos...")
	err := database.DeleteAll()
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString("¡Borrado de la base de datos completado!")
}
