package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/tomashphill/exquisite-corpse-go/pkg/models"
	"github.com/tomashphill/exquisite-corpse-go/pkg/models/sqlite"
)

func main() {
	dbType := flag.String("dbType", "sqlite", "Database Type")
	dsn := flag.String("dsn", "./data/exquisite-corpse.db", "Database DSN")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	var model models.ExqCorpModeler
	var err error

	// open db based on dbType env variable
	switch *dbType {
	case "sqlite":
		model, err = sqlite.OpenDB(*dsn)
	default:
		err = fmt.Errorf("%s is not a valid database type", *dbType)
	}
	if err != nil {
		errorLog.Fatal(err)
	}
	defer model.Close()

	// configure the app
	app := &Application{
		errorLog: errorLog,
		infoLog:  infoLog,
		model:    model,
	}

	e := echo.New()
	// server routes
	e.GET("/corpse/:name", app.getCorpse)

	e.Logger.Fatal(e.Start(":1234"))
}
