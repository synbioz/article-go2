package main

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
