package main

import (
	"fmt"
	"math"
)

// a spider must not be == 0
type spider byte

type part string

var head = part("h")
var tail = part("t")

type side struct {
	spider
	part
}

func (s *side) toString() string {
	return fmt.Sprintf("%v%s", s.spider, s.part)
}

type tile struct {
	top    side
	right  side
	bottom side
	left   side
}

func (t tile) rotate(times int) tile {
	for times > 0 {
		t = newTile(t.left, t.top, t.right, t.bottom)
		times--
	}
	return t
}

type solution []tile

type edge struct {
	a side
	b side
}

func (s solution) isValid() bool {
	for _, e := range s.edges() {
		if !match(e.a, e.b) {
			return false
		}
	}
	return true
}

func (s solution) isComplete() bool {
	return len(s) == 9 && s.isValid()
}

// edges returns a slice of touching edges in the grid
func (s solution) edges() (edges []edge) {
	if (len(s)) == 1 {
		return make([]edge, 0)
	}
	size := int(math.Sqrt(float64(len(s))))
	for i := 0; i < len(s); i++ {
		// Adding vertical edges between tiles in a row
		if i%size != 0 {
			edges = append(edges, edge{s[i-1].right, s[i].left})
		}
		// Add horizontal top edges for second and third row
		if i >= size {
			edges = append(edges, edge{s[i-size].bottom, s[i].top})
		}
	}
	return
}

func (s solution) place(t tile) solution {
	return append(s, t)
}

// All possible permutations for this solution with the remaining tiles
func (s solution) permutations(tiles []tile) []solution {
	var n []solution
	for _, t := range tiles {
		n = append(n, s.variations(t)...)
	}
	return n
}

// places the tile in the next space
func (s solution) variations(t tile) []solution {
	var n []solution
	for i := 0; i < 4; i++ {
		s1 := s.place(t.rotate(i))
		if s1.isValid() {
			n = append(n, s1)
		}
	}
	return n
}

func (s solution) print() {
	s.printRow(0, 3)
	s.printRow(3, 6)
	s.printRow(6, 9)
}

func (s solution) printRow(from, to int) {
	for i := from; i < to; i++ {
		if i < len(s) {
			fmt.Print("  ", s[i].top.toString(), "  ")
		} else {
			fmt.Print("  --  ")
		}
	}
	fmt.Println("")
	for i := from; i < to; i++ {
		if i < len(s) {
			fmt.Print(s[i].left.toString(), "  ", s[i].right.toString())
		} else {
			fmt.Print("--  --")
		}
	}
	fmt.Println("")
	for i := from; i < to; i++ {
		if i < len(s) {
			fmt.Print("  ", s[i].bottom.toString(), "  ")
		} else {
			fmt.Print("  --  ")
		}
	}
	fmt.Println("")
}

// match true if the edge matches
func match(a, b side) bool {
	if a.spider == 0 || b.spider == 0 {
		return false
	}
	return a.spider == b.spider && a.part != b.part
}

// Gives a solution to the puzzle by placing the given tiles in a square grid
func solve(tiles []tile, solutions []solution) solution {
	if len(solutions) == 0 {
		// start from scratch
		for _, tile := range tiles {
			solutions = append(solutions, solution{tile})
		}
	}
	for _, s := range solutions {
		solve(tiles[1:], s.variations(tiles[0]))
	}
	// solve(tiles[1:], s.place(tiles[0]))
	return solutions[0]
}

func newTile(top, right, bottom, left side) tile {
	return tile{
		top:    top,
		right:  right,
		bottom: bottom,
		left:   left,
	}
}

func main() {
	fmt.Println("Initialising puzzle")

	// Tarantula - The one that is big and chunky
	tarantula := spider(1)

	// Cellar - The on with the long legged spindly
	cellar := spider(2)

	// Johnson - The fat one with the brown ass
	johnson := spider(3)

	// Wolf - The one with the checkered abdomen
	wolf := spider(4)

	tiles := make([]tile, 9)

	tiles[0] = newTile(
		side{wolf, head},
		side{tarantula, tail},
		side{tarantula, head},
		side{johnson, head})

	tiles[1] = newTile(
		side{cellar, head},
		side{tarantula, head},
		side{wolf, head},
		side{tarantula, head})

	tiles[2] = newTile(
		side{cellar, tail},
		side{johnson, tail},
		side{cellar, head},
		side{tarantula, tail})

	tiles[3] = newTile(
		side{tarantula, tail},
		side{cellar, tail},
		side{tarantula, head},
		side{johnson, head})

	tiles[4] = newTile(
		side{wolf, tail},
		side{johnson, head},
		side{cellar, tail},
		side{cellar, head})

	tiles[5] = newTile(
		side{cellar, tail},
		side{tarantula, tail},
		side{wolf, tail},
		side{johnson, tail})

	tiles[6] = newTile(
		side{tarantula, tail},
		side{johnson, tail},
		side{johnson, tail},
		side{wolf, head})

	tiles[7] = newTile(
		side{cellar, head},
		side{johnson, tail},
		side{johnson, head},
		side{tarantula, head})

	tiles[8] = newTile(
		side{wolf, head},
		side{johnson, head},
		side{tarantula, tail},
		side{wolf, tail})

	// picSol := solution{tiles[0], tiles[1], tiles[2],
	// 	tiles[3], tiles[4], tiles[5],
	// 	tiles[6], tiles[7], tiles[8]}
	// picSol.print()
	// fmt.Printf("Solved: %v\n", picSol.isComplete())

	fmt.Println("Solving spiders")
	sol := solve(tiles, make([]solution, 0))
	sol.print()

}
