package main

import (
	"os"
	"time"
)

func main() {
	maze := [][]bool{
		{true, true, true, true, true, true, true, true, true, true, true, true, true},
		{false, false, true, false, false, false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false, false, false, false},
		{false, false, true, true, true, true, true, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false, false, false, false},
		{false, false, true, true, true, true, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false, false, false, false},
		{true, true, true, true, true, true, true, true, true, true, true, true, true},
	}

	ratmaze := NuevoRatMaze(maze)
	ratmaze.Animar(os.Stdout, 100*time.Millisecond)
}
