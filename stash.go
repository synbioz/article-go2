package quarto

type Stash [16]Piece

func MakeFullStash() (s Stash) {
  s[0]  = MakePiece(PIECE_TALL,  PIECE_BLACK, PIECE_ROUND,  PIECE_FULL)
  s[1]  = MakePiece(PIECE_TALL,  PIECE_BLACK, PIECE_ROUND,  PIECE_SPARSE)
  s[2]  = MakePiece(PIECE_TALL,  PIECE_BLACK, PIECE_SQUARE, PIECE_FULL)
  s[3]  = MakePiece(PIECE_TALL,  PIECE_BLACK, PIECE_SQUARE, PIECE_SPARSE)
  s[4]  = MakePiece(PIECE_TALL,  PIECE_WHITE, PIECE_ROUND,  PIECE_FULL)
  s[5]  = MakePiece(PIECE_TALL,  PIECE_WHITE, PIECE_ROUND,  PIECE_SPARSE)
  s[6]  = MakePiece(PIECE_TALL,  PIECE_WHITE, PIECE_SQUARE, PIECE_FULL)
  s[7]  = MakePiece(PIECE_TALL,  PIECE_WHITE, PIECE_SQUARE, PIECE_SPARSE)
  s[8]  = MakePiece(PIECE_SMALL, PIECE_BLACK, PIECE_ROUND,  PIECE_FULL)
  s[9]  = MakePiece(PIECE_SMALL, PIECE_BLACK, PIECE_ROUND,  PIECE_SPARSE)
  s[10] = MakePiece(PIECE_SMALL, PIECE_BLACK, PIECE_SQUARE, PIECE_FULL)
  s[11] = MakePiece(PIECE_SMALL, PIECE_BLACK, PIECE_SQUARE, PIECE_SPARSE)
  s[12] = MakePiece(PIECE_SMALL, PIECE_WHITE, PIECE_ROUND,  PIECE_FULL)
  s[13] = MakePiece(PIECE_SMALL, PIECE_WHITE, PIECE_ROUND,  PIECE_SPARSE)
  s[14] = MakePiece(PIECE_SMALL, PIECE_WHITE, PIECE_SQUARE, PIECE_FULL)
  s[15] = MakePiece(PIECE_SMALL, PIECE_WHITE, PIECE_SQUARE, PIECE_SPARSE)
  return
}

func (s *Stash) PickPieceAtIndex(i uint8) (p Piece) {
  p    = s[i]
  s[i] = PIECE_EMPTY
  return
}

func (s *Stash) PieceCount() (c uint8) {
  for i := 0; i < 16; i++ {
    if s[i] != PIECE_EMPTY {
      c++
    }
  }
  return
}
