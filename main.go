package main

import "fmt"

// a spider must not be == 0
type spider byte

type part string

var head = part("h")
var tail = part("t")

type side struct {
	spider
	part
}

type tile struct {
	top    side
	right  side
	bottom side
	left   side
}

type grid []tile
type solution struct {
	grid
	remainingTiles []tile
}

type edge struct {
	a side
	b side
}

func (t tile) rotate(times int) tile {
	for times > 0 {
		t = newTile(t.left, t.top, t.right, t.bottom)
		times--
	}
	return t
}

func (g grid) isValid() bool {
	for _, e := range g.edges() {
		if !match(e.a, e.b) {
			return false
		}
	}
	return true
}

func (g grid) isComplete() bool {
	return len(g) == 9 && g.isValid()
}

func (s solution) exhausted() bool {
	return len(s.remainingTiles) == 0
}

// edges returns a slice of touching edges in the grid
func (g grid) edges() (edges []edge) {
	if (len(g)) == 1 {
		return make([]edge, 0)
	}
	size := 3 // Locked to work for 3x3 grids
	for i := 0; i < len(g); i++ {
		// Adding vertical edges between tiles in a row
		if i%size != 0 {
			edges = append(edges, edge{g[i-1].right, g[i].left})
		}
		// Add horizontal top edges for second and third row
		if i >= size {
			edges = append(edges, edge{g[i-size].bottom, g[i].top})
		}
	}
	return
}

// All possible permutations for thig grid with the remaining tiles
func (s solution) permutations() []solution {
	var n []solution
	for i, t := range s.remainingTiles {
		tiles := make([]tile, len(s.remainingTiles))
		copy(tiles, s.remainingTiles)
		t1 := append(tiles[:i], tiles[i+1:]...)
		vars := s.variations(t)
		for j := 0; j < len(vars); j++ {
			n = append(n, solution{append([]tile(nil), vars[j]...), append([]tile(nil), t1...)})
		}
	}
	fmt.Println(len(n), "permutations for")
	s.print()
	fmt.Println("============================")
	fmt.Println("remainders", s.remainingTiles)
	return n
}

// places the tile in the next space in all possible rotations
func (g grid) variations(t tile) []grid {
	n := make([]grid, 0, 4)
	for i := 0; i < 4; i++ {
		g1 := append(g, t.rotate(i))
		// n = append(n, g1)
		if g1.isValid() {
			n = append(n, g1)
		}
	}
	return n
}

// match true if the edge matches
func match(a, b side) bool {
	if a.spider == 0 || b.spider == 0 {
		return false
	}
	return a.spider == b.spider && a.part != b.part
}

// search for solutions to the problem
func search(solutions []solution) []solution {
	var incomplete []solution
	var exhausted []solution
	for i := range solutions {
		for _, s := range solutions[i].permutations() {
			if s.exhausted() {
				if s.isComplete() {
					exhausted = append(exhausted, s)

				}
			} else {
				incomplete = append(incomplete, s)
			}
		}
	}
	fmt.Println("status: placed", len(solutions[0].grid)+1, "inc", len(incomplete), "exh", len(exhausted))
	if len(incomplete) == 0 {
		return exhausted
	}
	return search(incomplete)
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

	var tiles [9]tile

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

	fmt.Println("Solving spiders")
	sol := search([]solution{solution{grid{}, tiles[:]}})
	if len(sol) == 0 {
		fmt.Println("No solutions found")
		return
	}
	for _, s := range sol {
		s.print()
		fmt.Println("=====================================")
	}
}

func (g grid) print() {
	g.printRow(0, 3)
	g.printRow(3, 6)
	g.printRow(6, 9)
}

func (g grid) printRow(from, to int) {
	for i := from; i < to; i++ {
		if i < len(g) {
			fmt.Print("  ", g[i].top.toString(), "  ")
		} else {
			fmt.Print("  --  ")
		}
	}
	fmt.Println("")
	for i := from; i < to; i++ {
		if i < len(g) {
			fmt.Print(g[i].left.toString(), "  ", g[i].right.toString())
		} else {
			fmt.Print("--  --")
		}
	}
	fmt.Println("")
	for i := from; i < to; i++ {
		if i < len(g) {
			fmt.Print("  ", g[i].bottom.toString(), "  ")
		} else {
			fmt.Print("  --  ")
		}
	}
	fmt.Println("")
}

func (s *side) toString() string {
	return fmt.Sprintf("%v%s", s.spider, s.part)
}
