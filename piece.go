package main

import "strings"

type PieceColor uint8

const (
	White PieceColor = iota // 0
	Black                   // 1
)

type PieceType uint8

const (
	Free   PieceType = iota // 000
	Pawn                    // 001
	Knight                  // 010
	Bishop                  // 011
	Rook                    // 100
	Queen                   // 101
	King                    // 110
)

type Piece uint8

func (p Piece) GetString() string {
	var str string
	switch PieceType(p & 0b00000111) {
	case Pawn:
		str = "P"
	case Knight:
		str = "N"
	case Bishop:
		str = "B"
	case Rook:
		str = "R"
	case Queen:
		str = "Q"
	case King:
		str = "K"
	default:
		return "."
	}
	if p&0b10000000 == 0 {
		return strings.ToLower(str)
	}
	return str
}

func (p Piece) IsWhite() bool {
	return (p & 0b1000000) == 1
}

func (p Piece) IsBlack() bool {
	return (p & 0b1000000) == 0
}

func (p Piece) IsType(t PieceType) bool {
	return t == PieceType(0b00000111&p)
}

func (p Piece) IsFree() bool {
	return p == 0
}

func (p Piece) GetType() PieceType {
	return PieceType(0b00000111 & p)
}

func NewPiece(pieceType PieceType, pieceColor PieceColor) Piece {
	return Piece((uint8(pieceColor) << 7) | uint8(pieceType))
}

type PieceBoard struct {
	Piece    Piece
	Bitboard Bitboard
}