package main

import "golang.org/x/tour/pic"

// Pic は
func Pic(dx, dy int) [][]uint8 {

	pic := make([][]uint8, dy)
	picx := make([]uint8, dx)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			picx[j] = uint8(i * j)
		}

		pic[i] = picx
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
