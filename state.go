package main

import (
	"github.com/ksabanty/gator/internal/config"
	"github.com/ksabanty/gator/internal/database"
)

type State struct {
	db  *database.Queries
	cfg *config.Config
}
