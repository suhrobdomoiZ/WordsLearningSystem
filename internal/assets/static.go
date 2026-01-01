package assets

import "embed"

//go:embed static/*
var staticFS embed.FS

func GetStaticFS() embed.FS {
	return staticFS
}
