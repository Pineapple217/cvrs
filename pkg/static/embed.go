package static

import (
	"embed"
	"io/fs"
)

//go:embed build/*
var frontend embed.FS

func GetFrontend() fs.FS {
	f, _ := fs.Sub(frontend, "build")
	return f
}
