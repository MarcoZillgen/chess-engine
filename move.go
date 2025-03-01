package main

type Move struct {
	From int
	To   int
}

func (m Move) ValidPiece(fromPiece Piece, fromPos int, toPos int, toPiece Piece) bool {
	switch fromPiece.GetType() {
	case Pawn:
		return (fromPiece.IsBlack() &&
			(((fromPos-8 == toPos || (fromPos-16 == toPos && fromPos > 47)) && toPiece.IsFree()) ||
				((fromPos-9 == toPos || fromPos-7 == toPos) && !toPiece.IsFree()))) ||
			(fromPiece.IsBlack() &&
				(((fromPos+8 == toPos || (fromPos+16 == toPos && fromPos > 15)) && toPiece.IsFree()) ||
					((fromPos+9 == toPos || fromPos+7 == toPos) && !toPiece.IsFree())))
	case Knight:
		return (fromPos-17 == toPos || fromPos-15 == toPos || fromPos-10 == toPos || fromPos-6 == toPos ||
			fromPos+6 == toPos || fromPos+10 == toPos || fromPos+15 == toPos || fromPos+17 == toPos)
	case Bishop:
		return (fromPos-toPos)%9 == 0 || (fromPos-toPos)%7 == 0
	case Rook:
		return fromPos/8 == toPos/8 || fromPos%8 == toPos%8
	case Queen:
		return (fromPos/8 == toPos/8 || fromPos%8 == toPos%8) || (fromPos-toPos)%9 == 0 || (fromPos-toPos)%7 == 0
	case King:
		return (fromPos-1 == toPos || fromPos+1 == toPos || fromPos-8 == toPos || fromPos+8 == toPos ||
			fromPos-9 == toPos || fromPos-7 == toPos || fromPos+7 == toPos || fromPos+9 == toPos)
	}

	return false
}
