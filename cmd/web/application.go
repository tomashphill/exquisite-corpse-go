package main

import (
	"log"

	"github.com/tomashphill/exquisite-corpse-go/pkg/models"
)

// Application
type Application struct {
	model    models.ExqCorpModeler
	errorLog *log.Logger
	infoLog  *log.Logger
}
