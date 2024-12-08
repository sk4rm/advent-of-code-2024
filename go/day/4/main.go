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

func main() {
	for _, c := range match[1:] {
		seeds = explore(seeds, c)
	}
	fmt.Println(len(seeds))
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
