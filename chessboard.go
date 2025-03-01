package main

import (
	"fmt"
)

func (cb Chessboard) ValidateMove(move Move) bool {
	// check if move is within the board
	if move.From < 0 || move.From > 63 || move.To < 0 || move.To > 63 {
		return false
	}

	fromPiece := cb.GetPiece(move.From)
	// check if piece exists
	if fromPiece.IsFree() {
		return false
	}

	toPiece := cb.GetPiece(move.To)
	// check if target position is free or has an enemy piece
	if toPiece.IsFree() || toPiece.IsBlack() == fromPiece.IsBlack() {
		return false
	}

	// check wether this movement is applicable for the given piece
	return move.ValidPieces(fromPiece, move.From, move.To, toPiece)
}

type Chessboard struct {
	pieces [12]PieceBoard
}

func (cb Chessboard) GetPiece(position int) Piece {
	for _, piece := range cb.pieces {
		if piece.Bitboard&(Bitboard(1)<<position) != 0 {
			return piece.Piece
		}
	}
	return 0
}

// Validates chessboard that every position has max 1 piece
func (cb Chessboard) ValidateDoublePosition() bool {
	bb := Bitboard(0)
	for _, piece := range cb.pieces {
		bb &= piece.Bitboard
	}
	return bb == 0
}

// get all possible moves for a given color
func (cb Chessboard) GetPossibleMoves(white bool) []Move {
	var moves []Move
	for _, piece := range cb.pieces {
		if piece.Piece.IsWhite() == white {
			moves = append(moves, piece.Piece.GetPossibleMoves(piece.Bitboard, cb)...)
		}
	}
	return moves
}

// Create a new chessboard with the starting position
func getStartChessboard() Chessboard {
	return Chessboard{
		pieces: [12]PieceBoard{
			{NewPiece(Pawn, White), createBitboard([]int{8, 9, 10, 11, 12, 13, 14, 15})},
			{NewPiece(Pawn, Black), createBitboard([]int{48, 49, 50, 51, 52, 53, 54, 55})},
			{NewPiece(Knight, White), createBitboard([]int{1, 6})},
			{NewPiece(Knight, Black), createBitboard([]int{57, 62})},
			{NewPiece(Bishop, White), createBitboard([]int{2, 5})},
			{NewPiece(Bishop, Black), createBitboard([]int{58, 61})},
			{NewPiece(Rook, White), createBitboard([]int{0, 7})},
			{NewPiece(Rook, Black), createBitboard([]int{56, 63})},
			{NewPiece(Queen, White), createBitboard([]int{3})},
			{NewPiece(Queen, Black), createBitboard([]int{59})},
			{NewPiece(King, White), createBitboard([]int{4})},
			{NewPiece(King, Black), createBitboard([]int{60})},
		}}
}

// Draw the chessboard
func (cb Chessboard) Draw() {
	pos := Bitboard(1) << 63

	for i := 0; i < 64; i++ {
		if i%8 == 0 {
			fmt.Println()
		}

		fmt.Print(cb.GetPiece(i).GetString() + " ")

		pos >>= 1
	}
	fmt.Println()
}
