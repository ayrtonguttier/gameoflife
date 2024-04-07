package main

import (
	"io"
	"os"
	"os/exec"
	"time"
)

func main() {
    state := initSimpleState()

	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		print(state, os.Stdout)
		state = playRound(state)
		time.Sleep(300 * time.Millisecond)
	}
}
func initSimpleState() [][]int {
	state := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	return state
}
func initState(size int) [][]int {
	state := make([][]int, size)
	for i := range state {
		state[i] = make([]int, size)
	}
	return state
}

func print(state [][]int, stdout io.StringWriter) {
	for _, r := range state {
		for _, c := range r {
			if c > 0 {
				stdout.WriteString("x")
			} else {
				stdout.WriteString(" ")
			}
		}
		stdout.WriteString("\n")
	}

	stdout.WriteString("\n")
}

var dir = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},

	{0, -1},
	{0, 1},

	{1, -1},
	{1, 0},
	{1, 1},
}

func playRound(state [][]int) [][]int {
	out := make([][]int, len(state))

	for r, row := range state {
		out[r] = make([]int, len(state[r]))
		copy(out[r], state[r])
		for c := range row {
			nc := getNeighborCount(state, r, c)
			if nc < 2 {
				out[r][c] = 0
			}
			if nc > 3 {
				out[r][c] = 0
			}
			if nc == 3 {
				out[r][c] = 1
			}
		}
	}

	return out
}

func getNeighborCount(state [][]int, r, c int) int {
	count := 0
	for _, d := range dir {
		row, col := r+d[0], c+d[1]
		if row < 0 || row >= len(state) {
			continue
		}

		if col < 0 || col >= len(state[r]) {
			continue
		}

		if state[row][col] > 0 {
			count++
		}
	}

	return count
}
