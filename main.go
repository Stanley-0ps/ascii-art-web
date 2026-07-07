package main

import (
	"ascii-art/server"
	"embed"
)

//go:embed templates
var templateFs embed.FS

//go:embed static
var cssFiles embed.FS

// go build -ldflags "-X 'main.Style=K'" . (choose between A and k style)
var Style = "A"

func main() {
	server.TemplatesFS = templateFs
	server.Style = Style
	server.CssFS = cssFiles
	server.NewServer()
}
