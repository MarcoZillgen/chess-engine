package main

import (
	"fmt"
)

type Bitboard uint64

// Create a bitboard from a list of positions
func createBitboard(positions []int) Bitboard {
	bb := Bitboard(0)

	for _, position := range positions {
		// Check if the position is valid
		if position < 0 || position > 63 {
			fmt.Printf("Invalid position: %d\n", position)
			continue
		}

		// Set the bit at the given position
		bb |= Bitboard(1) << position
	}

	return bb
}
