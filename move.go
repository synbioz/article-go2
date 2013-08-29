package quarto

type Position struct {
  i, j  uint8
}

type Move struct {
  Pos  Position
  Idx  uint8
}

func (m Move) ApplyTo(pieceIndex uint8, g Game) Game {
  piece := g.Stash.PickPieceAtIndex(pieceIndex)
  g.Board.DropPieceAt(m.Pos, piece)
  return g
}

func (m Move) ToRepr() (uint8, uint8, uint8) {
  return m.Pos.i, m.Pos.j, m.Idx
}
