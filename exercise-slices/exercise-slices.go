package main

import "golang.org/golang/gotour/pic"

func Pic(dx, dy int) [][]uint8 {
	rows := make([][]uint8, dy)
	for i := range rows {
		rows[i] = make([]uint8, dx)

		for j := range rows[i] {
			rows[i][j] = uint8((i + j) / 2)
		}
	}
	return rows
}

func main() {
	pic.Show(Pic)
}
