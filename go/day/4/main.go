// Search for NOTE comments for common gotchas

package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const match string = "XMAS"
const part1 = false
const part2 = true

// Set during initializeSeeds()
var (
	rows          []string
	width, height int
	seeds         []seed
)

type seed struct {
	// position
	x int
	y int

	// direction
	dx int
	dy int
}

func (s *seed) move() {
	s.x += s.dx
	s.y += s.dy
}

func init() {
	seeds = []seed{}

	// NOTE Check for CRLF or LF
	rows = strings.Split(input, "\n")

	height = len(rows)
	if height > 0 {
		width = len(rows[0])
	}

	if part1 {
		for i, row := range rows {
			for j, cell := range row {

				if byte(cell) != match[0] {
					continue
				}

				// Eight directions
				for dx := -1; dx <= 1; dx += 1 {
					for dy := -1; dy <= 1; dy += 1 {

						if dx == 0 && dy == 0 {
							continue
						}

						seed := seed{j, i, dx, dy}
						seeds = append(seeds, seed)
					}
				}

			}
		}
	}
}

func main() {
	if part1 {
		for _, c := range match[1:] {
			seeds = explore(seeds, c)
		}
		fmt.Println(len(seeds))
	}

	if part2 {
		Ai := []int{}
		Aj := []int{}

		for i, row := range rows {
			if i == 0 || i == len(rows)-1 {
				continue
			}
			for j, letter := range row {
				if j == 0 || j == len(row)-1 {
					continue
				}

				if letter == 'A' {
					Ai = append(Ai, i)
					Aj = append(Aj, j)
				}

			}
		}

		count := 0
		for idx := range Ai {
			if checkAll(Ai[idx], Aj[idx]) {
				count += 1
			}
		}

		fmt.Println(count)
	}
}

func explore(seeds []seed, c rune) []seed {
	for idx := 0; idx < len(seeds); idx += 1 {

		seeds[idx].move()

		i := seeds[idx].y
		j := seeds[idx].x

		if i < 0 || i >= height || j < 0 || j >= width {
			seeds = removeSeed(seeds, idx)
			idx -= 1
			continue
		}

		letter := rune(rows[i][j])

		if letter != c {
			seeds = removeSeed(seeds, idx)
			idx -= 1
			continue
		}
	}

	return seeds
}

func removeSeed(seeds []seed, idx int) []seed {
	if len(seeds) <= 1 {
		return []seed{}
	}

	// https://stackoverflow.com/a/37335777
	seeds[idx] = seeds[len(seeds)-1]
	return seeds[:len(seeds)-1]
}

func checkAll(i, j int) bool {
	return check1(i, j) || check2(i, j) || check3(i, j) || check4(i, j)
}

func check1(i, j int) bool {
	return rows[i-1][j-1] == 'M' && rows[i-1][j+1] == 'M' && rows[i+1][j-1] == 'S' && rows[i+1][j+1] == 'S'
}

func check2(i, j int) bool {
	return rows[i-1][j-1] == 'M' && rows[i-1][j+1] == 'S' && rows[i+1][j-1] == 'M' && rows[i+1][j+1] == 'S'
}

func check3(i, j int) bool {
	return rows[i-1][j-1] == 'S' && rows[i-1][j+1] == 'S' && rows[i+1][j-1] == 'M' && rows[i+1][j+1] == 'M'
}

func check4(i, j int) bool {
	return rows[i-1][j-1] == 'S' && rows[i-1][j+1] == 'M' && rows[i+1][j-1] == 'S' && rows[i+1][j+1] == 'M'
}
