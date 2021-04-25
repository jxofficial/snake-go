package main

import "testing"

func TestGridGetRandomUnoccupiedPixel(t *testing.T) {
	gSlice := make([][]bool, 3)
	for i := range gSlice {
		gSlice[i] = make([]bool, 3)
	}

	g := grid{gSlice, 3, 3}

	// all occupied except x = 1, y = 0
	g.g[0][0] = true
	g.g[0][2] = true
	g.g[1][0] = true
	g.g[1][1] = true
	g.g[1][2] = true
	g.g[2][0] = true
	g.g[2][1] = true
	g.g[2][2] = true

	unoccupiedPos := position{1, 0}

	pixel, err := g.getRandomUnoccupiedPixel()

	if pixel != unoccupiedPos || err != nil {
		t.Fatalf("Expected %+v, got %+v with error: %v\n", unoccupiedPos, pixel, err)
	}
}

func TestGridIsFullyOccupiedTrue(t *testing.T) {
	gridSlice := [][]bool{{true, true}, {true, true}}
	g := grid{g: gridSlice}

	isFull := g.isCompletelyOccupied()

	if !isFull {
		t.Fatalf("Expected grid is full to be %v, got %v\n", true, isFull)
	}
}

func TestGridIsFullyOccupiedFalse(t *testing.T) {
	gridSlice := [][]bool{{false, true}, {true, true}}
	g := grid{g: gridSlice}

	isFull := g.isCompletelyOccupied()

	if isFull {
		t.Fatalf("Expected grid is full to be %v, got %v\n", false, isFull)
	}
}
