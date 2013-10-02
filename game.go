package quarto

type Game struct {
  Board  Board
  Stash  Stash
}

func (g *Game) IsWinningAt(pos Position) bool {
  return g.Board.IsWinningAt(pos)
}

func (g *Game) IsWinning() bool {
  return g.Board.IsWinningAt(Position{0, 0}) ||
         g.Board.IsWinningAt(Position{0, 1}) ||
         g.Board.IsWinningAt(Position{0, 2}) ||
         g.Board.IsWinningAt(Position{0, 3}) ||
         g.Board.IsWinningAt(Position{1, 0}) ||
         g.Board.IsWinningAt(Position{2, 0}) ||
         g.Board.IsWinningAt(Position{3, 0})
}

func (g *Game) PossibleMoves(pieceIndex uint8) []Move {
  var moves []Move
  count := int(g.Board.EmptyCellCount()) * (int(g.Stash.PieceCount()) - 1)

  if count == 0 {
    moves = make([]Move, 1)
    moves[0] = Move{g.Board.firstEmptyCellPosition(), 0}
  } else {
    moves = make([]Move, count)
    var i, j, k, n uint8
    for k = 0; k < 16; k++ {
      if g.Stash[k] != PIECE_EMPTY && k != pieceIndex {
        for i = 0; i < 4; i++ {
          for j = 0; j < 4; j++ {
            if g.Board[i][j] == PIECE_EMPTY {
              moves[n] = Move{Position{i, j}, k}
              n++
            }
          }
        }
      }
    }
  }

  return moves
}

func (g *Game) PlayWith(pieceIndex uint8) (Move, bool) {
  moves := g.PossibleMoves(pieceIndex)

  // If there is a wining move, play it
  for _, move := range moves {
    game := move.ApplyTo(pieceIndex, *g)

    if game.IsWinningAt(move.Pos) {
      return move, true
    }
  }

  // Finding the better between one is easy...
  if len(moves) == 1 { return moves[0], false }

MyMoves:
  for _, move := range moves {
    game := move.ApplyTo(pieceIndex, *g)

    opponentMoves := game.PossibleMoves(move.Idx)
    for _, opponentMove := range opponentMoves {
      opponentGame := opponentMove.ApplyTo(move.Idx, game)
      if opponentGame.IsWinningAt(opponentMove.Pos) { continue MyMoves }
    }

    return move, false
  }

  return moves[0], false
}
