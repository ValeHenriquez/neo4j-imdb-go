package database

import (
	"context"
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var driver neo4j.DriverWithContext
var ctx = context.Background()

func Setup() {
	var err error
	dbUri := os.Getenv("DB_URI")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	driver, err = neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))
	if err != nil {
		log.Fatal("DB CONNECTION ERROR", err)
	}
	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		log.Fatal("VERIFY DB CONN ERROR", err)
	}

}

func Close() {
	err := driver.Close(ctx)
	if err != nil {
		log.Fatal("DB NOT CLOSED", err)
	}
}

func GetDriver() neo4j.DriverWithContext {
	return driver
}
