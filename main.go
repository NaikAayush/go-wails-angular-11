package main

import (
	_ "embed"

	"github.com/wailsapp/wails"
)

func basic() string {
	return "World!"
}

//go:embed frontend/dist/frontend/main.js
var js string

//go:embed frontend/dist/frontend/styles.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "go-gui",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Run()
}
