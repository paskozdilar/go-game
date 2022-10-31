package main

import (
	"image"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

type Game struct {
	fps    int
	width  int
	height int
}

func (game Game) Run(callback func(*image.RGBA)) {
	// Define corners
	topLeft := image.Point{0, 0}
	bottomRight := image.Point{game.width, game.height}

	// Create pixel matrix
	rgbaImage := image.NewRGBA(image.Rectangle{topLeft, bottomRight})

	// Create Fyne app and window
	mainApp := app.New()
	imageWindow := mainApp.NewWindow("Images")

	// Create full-window canvas
	canvasToWrite := canvas.NewRasterFromImage(rgbaImage)
	imageWindow.SetContent(canvasToWrite)
	imageWindow.Resize(fyne.NewSize(float32(game.width), float32(game.height)))

	go func() {
		// Run game callback in the background every 1/fps seconds
		for {
			callback(rgbaImage)
			canvas.Refresh(canvasToWrite)
			time.Sleep(time.Duration(1000/game.fps) * time.Millisecond)
		}
	}()

	// Run Fyne app
	imageWindow.ShowAndRun()
}
