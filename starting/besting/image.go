package besting

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"math/rand"
	"os"
)

func generateImage() {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	//red := color.RGBA{255, 0, 0, 255}
	for x := 0; x < 200; x++ {
		for y := 0; y < 100; y++ {
			if x == 50 || y == 50 || x == y {
				img.Set(x, y, color.White)
			} else {
				img.Set(x, y, color.Black)
			}
		}
	}
	file, err := os.Create("data/red.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

func generateGif() {
	const (
		numberOfFrames   = 64
		timeForEachFrame = 100
	)
	myGif := gif.GIF{LoopCount: numberOfFrames}

	for i := 0; i < numberOfFrames; i++ {
		rect := image.Rect(0, 0, 100, 100)

		img := image.NewPaletted(
			rect,
			[]color.Color{
				color.White,
				color.Black,
				color.RGBA{R: 255, G: 0, B: 0, A: 255}})

		m := rand.Intn(100)
		fmt.Println(m)
		for x := 0; x < rect.Dx(); x++ {
			for y := 0; y < rect.Dy(); y++ {
				if y == m || m == x {
					img.Set(x, y, color.RGBA{R: 255, G: 0, B: 0, A: 255})
				}
			}
		}
		myGif.Image = append(myGif.Image, img)
		myGif.Delay = append(myGif.Delay, timeForEachFrame)
	}

	file, err := os.Create("data/red.gif")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	gif.EncodeAll(file, &myGif)
}
