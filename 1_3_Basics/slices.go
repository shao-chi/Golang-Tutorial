// go install golang.org/x/website/tour@latest
// go mod init golang.org
// go mod tidy
// go get golang.org/x/tour/pic

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)

	for y := range img {
		img[y] = make([]uint8, dx)

		for x := range img[y] {
			// img[y][x] = uint8((x+y) / 2)
			// img[y][x] = uint8(x*y)
			img[y][x] = uint8(x ^ y)
		}
	}

	return img
}

func main() {
	pic.Show(Pic)
}
