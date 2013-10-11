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
  // Adapt the depth to the stash size
  var depth int
  if g.Stash.PieceCount() > 10 {
    depth = 2
  } else {
    depth = 4
  }

  move, _, _ := g.playWith(pieceIndex, 0, depth)

  game := move.ApplyTo(pieceIndex, *g)

  return move, game.IsWinningAt(move.Pos)
}

func (g *Game) playWith(pieceIndex uint8, depth, maxDepth int) (best_move Move, win bool, opponentWin bool) {
  moves := g.PossibleMoves(pieceIndex)

  // If there is a wining move, play it
  for _, move := range moves {
    game := move.ApplyTo(pieceIndex, *g)

    if game.IsWinningAt(move.Pos) {
      return move, true, false
    }
  }

  // If there is only one move to play, tie
  // If the maximum depth is reached, unknown result
  if len(moves) == 1 || depth >= maxDepth {
    return moves[0], false, false
  }

  // Find a move that will not give the victory to the opponent
  best_move_found := false
  for _, move := range moves {
    game := move.ApplyTo(pieceIndex, *g)
    _, opponentWin, win := game.playWith(move.Idx, depth + 1, maxDepth)

    if win {
      return move, true, false
    }

    if !opponentWin {
      best_move = move
      best_move_found = true
    }
  }

  if best_move_found {
    return best_move, false, false
  }

  return moves[0], false, true
}
