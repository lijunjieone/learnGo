package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main2() {
	createImage2()
}

func createImage2() {
	size := 300

	pic := image.NewGray(image.Rect(0, 0, size, size))

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / float64(size)

		y := float64(size)/2 - math.Sin(s)*float64(size)/2
		pic.SetGray(x, int(y), color.Gray{0})

	}

	file, err := os.Create("sin.png")

	if err != nil {
		log.Fatal(err)
	}

	png.Encode(file, pic)

	file.Close()

}
