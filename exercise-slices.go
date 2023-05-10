package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		picture[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// Choose the desired function to generate pixel values
			picture[y][x] = uint8((x + y) / 2) // Example: (x + y) / 2
			// picture[y][x] = uint8(x * y)    // Example: x * y
			// picture[y][x] = uint8(x ^ y)    // Example: x ^ y
		}
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
