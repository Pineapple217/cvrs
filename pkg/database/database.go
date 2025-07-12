package database

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/Pineapple217/cvrs/pkg/config"
	"github.com/Pineapple217/cvrs/pkg/ent"
	_ "github.com/Pineapple217/cvrs/pkg/ent/runtime"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Client *ent.Client
	Conf   config.Database
}

func NewDatabase(conf config.Database) (*Database, error) {
	var err error
	for _, p := range []string{
		conf.DataLocation,
		path.Join(conf.DataLocation, TEMP_DIR),
		path.Join(conf.DataLocation, IMG_DIR),
		path.Join(conf.DataLocation, BACKUP_DIR),
	} {
		err = CreateDir(p)
		if err != nil {
			return nil, err
		}
	}

	client, err := ent.Open("sqlite3", conf.SqliteOptions)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}
	db := &Database{
		Client: client,
		Conf:   conf,
	}
	return db, nil
}

func CreateDir(p string) error {
	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		slog.Info("Creating directory", "dir", p)
		err := os.Mkdir(p, 0755)
		if err != nil {
			slog.Error("Failed to create directory",
				"dir", p,
				"error", err,
			)
			return err
		}
	}
	return nil
}
