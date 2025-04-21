package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/Pineapple217/cvrs/pkg/ent"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Client *ent.Client
}

func NewDatabase(sqliteOptions string) (*Database, error) {
	client, err := ent.Open("sqlite3", sqliteOptions)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed opening connection to sqlite: %v", err))
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, errors.New(fmt.Sprintf("failed creating schema resources: %v", err))
	}
	db := &Database{
		Client: client,
	}
	return db, nil
}
