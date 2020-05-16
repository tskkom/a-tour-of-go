package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, 0)
	for y := 0; y < dy; y++ {
		s := make([]uint8, 0)
		for x := 0; x < dx; x++ {
			s = append(s, uint8(draw(x, y)))
		}
		res = append(res, s)
	}
	return res
}

func draw(x, y int) int {
	return (x + y) / 2
}

func main() {
	pic.Show(Pic)
}
