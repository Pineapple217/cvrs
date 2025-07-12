package database

import (
	"archive/tar"
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/Pineapple217/cvrs/pkg/config"
	"github.com/spf13/cobra"
)

const BACKUP_DIR = "/backups"

func (db *Database) CreateFullBackup(ctx context.Context) error {
	sqliteFilePath, err := db.CreateSQLiteBackup(ctx)
	if err != nil {
		return err
	}
	defer os.Remove(sqliteFilePath)

	timeString := time.Now().Format("2006-01-02_15-04-05")
	tarballName := path.Join(db.Conf.DataLocation, BACKUP_DIR, fmt.Sprintf("backup_%s.tar.gz", timeString))
	tarballFile, err := os.Create(tarballName)
	if err != nil {
		return err
	}
	defer tarballFile.Close()

	gzipWriter := tar.NewWriter(tarballFile)
	defer gzipWriter.Close()

	err = addFileToTar(gzipWriter, sqliteFilePath, "database.db")
	if err != nil {
		return err
	}

	folderPath := path.Join(db.Conf.DataLocation, IMG_DIR)
	err = filepath.Walk(folderPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filePath == folderPath {
			return nil
		}

		relPath, err := filepath.Rel(folderPath, filePath)
		if err != nil {
			return err
		}
		headerName := filepath.Join(IMG_DIR, filepath.ToSlash(relPath))
		return addFileToTar(gzipWriter, filePath, headerName)
	})

	return err
}

func (db *Database) CreateSQLiteBackup(ctx context.Context) (string, error) {
	timeString := time.Now().Format("2006-01-02_15-04-05")
	path := path.Join(db.Conf.DataLocation, TEMP_DIR, fmt.Sprintf("%s.db", timeString))
	_, err := db.Client.ExecContext(ctx, fmt.Sprintf("vacuum into \"%s\"", path))
	return path, err
}

func addFileToTar(gzipWriter *tar.Writer, filePath, headerName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(fileInfo, "")
	if err != nil {
		return err
	}

	if headerName != "" {
		header.Name = headerName
	} else {
		header.Name = fileInfo.Name()
	}

	if err := gzipWriter.WriteHeader(header); err != nil {
		return err
	}

	_, err = io.Copy(gzipWriter, file)
	return err
}

func GetBackupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "backup",
		Short: "create full instance backup",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := config.Load()
			if err != nil {
				return err
			}
			db, err := NewDatabase(conf.Database)
			if err != nil {
				return err
			}
			err = db.CreateFullBackup(cmd.Context())
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
