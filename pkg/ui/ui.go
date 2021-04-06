package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed files/*
var files embed.FS

func StaticAssets() http.FileSystem {

	assets, err := fs.Sub(files, "files")

	if err != nil {
		panic(err)
	}

	return http.FS(assets)
}
