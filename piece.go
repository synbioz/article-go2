package quarto

type Piece uint8

const (
  PIECE_EMPTY   = 0
  PIECE_TALL    = 1
  PIECE_SMALL   = 2
  PIECE_BLACK   = 4
  PIECE_WHITE   = 8
  PIECE_ROUND   = 16
  PIECE_SQUARE  = 32
  PIECE_FULL    = 64
  PIECE_SPARSE  = 128
)

func MakePiece(c1, c2, c3, c4 uint8) Piece {
  return Piece( c1 | c2 | c3 | c4 )
}

func (p Piece) IsValid() bool {
  return p == 0 ||
         ((p & PIECE_TALL)  > 0) != ((p & PIECE_SMALL)  > 0) &&
         ((p & PIECE_BLACK) > 0) != ((p & PIECE_WHITE)  > 0) &&
         ((p & PIECE_ROUND) > 0) != ((p & PIECE_SQUARE) > 0) &&
         ((p & PIECE_FULL)  > 0) != ((p & PIECE_SPARSE) > 0)
}
