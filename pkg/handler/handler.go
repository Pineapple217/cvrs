package handler

import "github.com/Pineapple217/cvrs/pkg/database"

type Handler struct {
	DB *database.Database
}

func NewHandler(DB *database.Database) *Handler {
	return &Handler{
		DB: DB,
	}
}
