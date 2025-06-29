package database

import (
	"context"
	"fmt"

	"github.com/Pineapple217/cvrs/pkg/config"
	"github.com/Pineapple217/cvrs/pkg/ent"
	_ "github.com/Pineapple217/cvrs/pkg/ent/runtime"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Client *ent.Client
}

func NewDatabase(conf config.Database) (*Database, error) {
	client, err := ent.Open("sqlite3", conf.SqliteOptions)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}
	db := &Database{
		Client: client,
	}
	return db, nil
}
