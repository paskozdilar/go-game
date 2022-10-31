package main

import (
	"image"
	"image/color"
	_ "image/color"
)

func main() {
	game := Game{
		fps:    60,
		width:  640,
		height: 480,
	}

	x, y := 0, 0
	red := color.RGBA{0xff, 0x00, 0x00, 0x00}

	game.Run(func(img *image.RGBA) {
		// Re-draw screen here
		img.Set(x, y, red)

		// Calculate next pixel
		x++
		if x == game.width {
			x = 0
			y++
			if y == game.height {
				y = 0
			}
		}
	})
}
