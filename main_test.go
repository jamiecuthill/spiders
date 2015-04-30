package main

import (
	"fmt"
	"testing"
)

// tesalating tile can be used to create any size valid solution grid
//   1h
// 2t  2h
//   1t
var tesalatingTile = newTile(
	side{spider(1), head},
	side{spider(2), head},
	side{spider(1), tail},
	side{spider(2), tail})

// non-tesalating tile can be used to create any size invalid solution grid
//   1h
// 4t  2t
//   3h
var nonTesalatingTile = newTile(
	side{spider(1), head},
	side{spider(2), tail},
	side{spider(3), head},
	side{spider(4), tail})

func TestMatch(t *testing.T) {
	if !match(side{spider(1), head}, side{spider(1), tail}) {
		t.Error("should be a match")
	}
}

func TestMatchOnSpiderZero(t *testing.T) {
	if match(side{spider(0), head}, side{spider(0), tail}) {
		t.Error("should not be a match")
	}
}

func TestNotMatchOnPart(t *testing.T) {
	if match(side{spider(1), head}, side{spider(1), head}) {
		t.Error("should not be a match")
	}
}

func TestNotMatchOnSpider(t *testing.T) {
	if match(side{spider(1), head}, side{spider(2), tail}) {
		t.Error("should not be a match")
	}
}

func TestEdgesIsEmptyForEmptySolution(t *testing.T) {
	s := grid([]tile{})
	l := len(s.edges())
	if l != 0 {
		t.Errorf("unexpected number of edges %d, want %d", l, 0)
	}
}

func TestEdgesIsEmptyForSingleTile(t *testing.T) {
	s := grid([]tile{tesalatingTile})
	l := len(s.edges())
	if l != 0 {
		t.Errorf("unexpected number of edges %d, want %d", l, 0)
	}
}

func TestEdgesIsCorrectLengthForIncompleteSolution(t *testing.T) {
	s := grid([]tile{tesalatingTile, tesalatingTile, tesalatingTile,
		tesalatingTile})
	l := len(s.edges())
	if l != 3 {
		t.Errorf("unexpected number of edges %d, want %d", l, 3)
	}
}

func TestEdgesHoldsCorrectDataForIncompleteSolution(t *testing.T) {
	s := grid([]tile{tesalatingTile, tesalatingTile, tesalatingTile,
		tesalatingTile})

	// Need to make this test not care about order of edges
	expectedEdges := []edge{
		edge{tesalatingTile.right, tesalatingTile.left},
		edge{tesalatingTile.right, tesalatingTile.left},
		edge{tesalatingTile.bottom, tesalatingTile.top}}

	for i, e := range s.edges() {
		err := assertEdge(e, expectedEdges[i].a, expectedEdges[i].b)
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func TestEdgesForFullGrid(t *testing.T) {
	g := grid([]tile{tesalatingTile, tesalatingTile, tesalatingTile,
		tesalatingTile, tesalatingTile, tesalatingTile,
		tesalatingTile, tesalatingTile, tesalatingTile})

	if len(g.edges()) != 12 {
		t.Errorf("unexpected number of edges %d, want %d", len(g.edges()), 12)
	}
}

func assertEdge(e edge, a, b side) error {
	if e.a != a && e.b != b {
		return fmt.Errorf("unxpected edge %v, want {%v, %v}", e, a, b)
	}
	return nil
}

// solution 1x1 is valid
func TestSolution1by1IsValid(t *testing.T) {
	s := grid([]tile{tesalatingTile})
	if !s.isValid() {
		t.Error("solution must be valid")
	}
}

// solution is valid
func TestSolutionIsValid(t *testing.T) {
	s := grid([]tile{tesalatingTile, tesalatingTile, tesalatingTile, tesalatingTile})
	if !s.isValid() {
		t.Error("solution must be valid")
	}
}

// solution is not valid
func TestSolutionIsNotValid(t *testing.T) {
	s := grid([]tile{nonTesalatingTile, nonTesalatingTile, nonTesalatingTile, nonTesalatingTile})
	if s.isValid() {
		t.Error("solution must not be valid")
	}
}

func TestSolutionIsComplete(t *testing.T) {
	s := grid([]tile{tesalatingTile, tesalatingTile, tesalatingTile, tesalatingTile, tesalatingTile, tesalatingTile, tesalatingTile, tesalatingTile, tesalatingTile})
	if !s.isComplete() {
		t.Error("solution must be complete")
	}
}

func TestSolutionIsNotCompleteNotValid(t *testing.T) {
	s := grid([]tile{nonTesalatingTile, nonTesalatingTile, nonTesalatingTile, nonTesalatingTile, nonTesalatingTile, nonTesalatingTile, nonTesalatingTile, nonTesalatingTile, nonTesalatingTile})
	if s.isComplete() {
		t.Error("solution must not be complete")
	}
}

func TestSolutionIsNotCompleteNot9(t *testing.T) {
	s := grid([]tile{tesalatingTile, tesalatingTile, tesalatingTile, tesalatingTile, tesalatingTile, tesalatingTile, tesalatingTile, tesalatingTile})
	if s.isComplete() {
		t.Error("solution must not be complete")
	}
}

func TestRotateTile(t *testing.T) {
	t1 := tesalatingTile.rotate(1)
	if t1.top != tesalatingTile.left {
		t.Errorf("top was %v, wanted %v after rotation", t1.top, tesalatingTile.left)
	}
	if t1.right != tesalatingTile.top {
		t.Errorf("right was %v, wanted %v after rotation", t1.right, tesalatingTile.top)
	}
	if t1.bottom != tesalatingTile.right {
		t.Errorf("bottom was %v, wanted %v after rotation", t1.bottom, tesalatingTile.right)
	}
	if t1.left != tesalatingTile.bottom {
		t.Errorf("left was %v, wanted %v after rotation", t1.left, tesalatingTile.bottom)
	}
}

func TestVariationsFromEmpty(t *testing.T) {
	s := grid([]tile{})
	neighbours := s.variations(tesalatingTile)
	if len(neighbours) != 4 {
		t.Errorf("unexpected number of neighbours %d, wanted %d", len(neighbours), 4)
		return
	}
	if neighbours[0][0] != tesalatingTile {
		t.Errorf("expected to get the tile in the neighbours")
	}
	if neighbours[1][0] != tesalatingTile.rotate(1) {
		t.Errorf("expected to get the tile in the neighbours")
	}
	if neighbours[2][0] != tesalatingTile.rotate(2) {
		t.Errorf("expected to get the tile in the neighbours")
	}
	if neighbours[3][0] != tesalatingTile.rotate(3) {
		t.Errorf("expected to get the tile in the neighbours")
	}
}

func TestVariationsOnlyValid(t *testing.T) {
	s := grid([]tile{tesalatingTile})
	neighbours := s.variations(tesalatingTile)
	if len(neighbours) != 1 {
		t.Errorf("unexpected number of neighbours %d, wanted %d", len(neighbours), 1)
		return
	}
	if neighbours[0][1] != tesalatingTile {
		t.Errorf("expected to get the tile in the neighbours")
	}
}

func TestIsValidAgain(t *testing.T) {
	s := grid([]tile{tesalatingTile, nonTesalatingTile})
	if s.isValid() {
		t.Errorf("solution should not be valid")
	}
}

func TestPermutationsPlacesAllTilesInAllVariationsStartingWithEmpty(t *testing.T) {
	s := solution{grid{}, []tile{tesalatingTile, nonTesalatingTile}}
	solutions := s.permutations()
	if len(solutions) != 8 {
		t.Errorf("unexpected number of solutions %d, want %d", len(solutions), 8)
	}
	for _, sol := range solutions {
		if !sol.isValid() {
			t.Errorf("solution permutation was not valid")
			return
		}
	}
}

func TestPermutationsPlacesAllTilesInAllValidVariations(t *testing.T) {
	s := solution{grid{tesalatingTile}, []tile{tesalatingTile, nonTesalatingTile}}
	solutions := s.permutations()
	if len(solutions) != 2 {
		t.Errorf("unexpected number of solutions %d, want %d", len(solutions), 2)
	}
	for _, sol := range solutions {
		if !sol.isValid() {
			t.Errorf("solution permutation was not valid")
			return
		}
	}
}

// func TestAllPermutationsAreFound(t *testing.T) {
// 	tarantula := spider(1)
// 	cellar := spider(2)
// 	johnson := spider(3)
// 	wolf := spider(4)
// 	s := solution{grid{}, []tile{
// 		newTile(
// 			side{wolf, head},
// 			side{tarantula, tail},
// 			side{tarantula, head},
// 			side{johnson, head}),
// 		newTile(
// 			side{cellar, head},
// 			side{tarantula, head},
// 			side{wolf, head},
// 			side{tarantula, head}),
// 		newTile(
// 			side{tarantula, tail},
// 			side{cellar, tail},
// 			side{tarantula, head},
// 			side{johnson, head}),
// 		newTile(
// 			side{cellar, head},
// 			side{johnson, tail},
// 			side{johnson, head},
// 			side{tarantula, head}),
// 	}}
// 	sols := search([]solution{s})
// 	for _, v := range sols {
// 		v.print()
// 	}
// }

func TestImmutable(t *testing.T) {
	g := grid{tesalatingTile}
	gs := g.variations(newTile(
		side{spider(1), head},
		side{spider(2), tail},
		side{spider(3), head},
		side{spider(2), tail}))
	for _, v := range gs {
		v.print()
	}
}
