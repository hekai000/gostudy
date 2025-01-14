package bfs

import (
	"testing"
)

func TestNumIslands2(t *testing.T) {
	tests := []struct {
		grid   [][]byte
		result int
	}{
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			result: 3,
		},
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '1'},
				{'0', '0', '0', '0', '0'},
				{'1', '1', '1', '1', '1'},
			},
			result: 2,
		},
		{
			grid: [][]byte{
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			result: 0,
		},
		{
			grid: [][]byte{
				{'1'},
			},
			result: 1,
		},
		{
			grid: [][]byte{
				{'1', '0'},
				{'0', '1'},
			},
			result: 2,
		},
		{
			grid:   [][]byte{},
			result: 0,
		},
	}

	for _, test := range tests {
		if got := numIslands2(test.grid); got != test.result {
			t.Errorf("numIslands2(%v) = %d; want %d", test.grid, got, test.result)
		}
	}
}
