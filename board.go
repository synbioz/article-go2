package quarto

type Board [4][4]Piece

func (b *Board) DropPieceAt(pos Position, p Piece) {
  b[pos.i][pos.j] = p
}

func (b *Board) IsWinningAt(pos Position) bool {
  return b.isRowWinningAt(pos) ||
         b.isColWinningAt(pos) ||
         b.isDiag1WinningAt(pos) ||
         b.isDiag2WinningAt(pos)
}

func (b *Board) isRowWinningAt(pos Position) bool {
  return (b[pos.i][0] & b[pos.i][1] & b[pos.i][2] & b[pos.i][3]) > 0
}

func (b *Board) isColWinningAt(pos Position) bool {
  return (b[0][pos.j] & b[1][pos.j] & b[2][pos.j] & b[3][pos.j]) > 0
}

func (b *Board) isDiag1WinningAt(pos Position) bool {
  return (pos.i + pos.j == 3) && (b[0][3] & b[1][2] & b[2][1] & b[3][0] > 0)
}

func (b *Board) isDiag2WinningAt(pos Position) bool {
  return (pos.i == pos.j) && (b[0][0] & b[1][1] & b[2][2] & b[3][3] > 0)
}


func (b *Board) EmptyCellCount() (c uint8) {
  for i := 0; i < 4; i++ {
    for j := 0; j < 4; j++ {
      if b[i][j] == PIECE_EMPTY {
        c++
      }
    }
  }
  return
}
