package main

import "fmt"

type Game struct {
  Board  Board
  Stash  Stash
}

func (g *Game) IsWinning(pos Position) bool {
  return g.Board.IsWinning(pos)
}

func (g *Game) PossibleMoves(pieceIndex uint8) []Move {
  count := int(g.Board.EmptyCellCount()) * (int(g.Stash.PieceCount()) - 1)
  moves := make([]Move, count)

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

  return moves
}

func (g *Game) PlayWith(pieceIndex uint8) Move {
  fmt.Printf("Stash: ")
  for _, p := range g.Stash { fmt.Printf("%3d ; ", p) }
  fmt.Printf("\n")

  moves := g.PossibleMoves(pieceIndex)

MyMoves:
  for _, move := range moves {
    game := move.ApplyTo(pieceIndex, *g)
    opponent_moves := game.PossibleMoves(move.Idx)

    for _, opponent_move := range opponent_moves {
      opponent_game := opponent_move.ApplyTo(move.Idx, game)
      if opponent_game.IsWinning(move.Pos) { continue MyMoves }
    }

    return move
  }

  return moves[0]
}

func main() {
  game := Game{Stash: MakeFullStash()}

  move := game.PlayWith(0)
  game  = move.ApplyTo(0, game)
  fmt.Printf("Play at (%d, %d), give #%d\n", move.Pos.i, move.Pos.j, move.Idx)

  o_move := game.PlayWith(move.Idx)
  game    = o_move.ApplyTo(move.Idx, game)
  fmt.Printf("Play at (%d, %d), give #%d\n", o_move.Pos.i, o_move.Pos.j, o_move.Idx)

  move  = game.PlayWith(o_move.Idx)
  game  = move.ApplyTo(o_move.Idx, game)
  fmt.Printf("Play at (%d, %d), give #%d\n", move.Pos.i, move.Pos.j, move.Idx)

  o_move = game.PlayWith(move.Idx)
  game   = o_move.ApplyTo(move.Idx, game)
  fmt.Printf("Play at (%d, %d), give #%d\n", o_move.Pos.i, o_move.Pos.j, o_move.Idx)

  move  = game.PlayWith(o_move.Idx)
  game  = move.ApplyTo(o_move.Idx, game)
  fmt.Printf("Play at (%d, %d), give #%d\n", move.Pos.i, move.Pos.j, move.Idx)

  o_move = game.PlayWith(move.Idx)
  game   = o_move.ApplyTo(move.Idx, game)
  fmt.Printf("Play at (%d, %d), give #%d\n", o_move.Pos.i, o_move.Pos.j, o_move.Idx)

  move  = game.PlayWith(o_move.Idx)
  game  = move.ApplyTo(o_move.Idx, game)
  fmt.Printf("Play at (%d, %d), give #%d\n", move.Pos.i, move.Pos.j, move.Idx)

  o_move = game.PlayWith(move.Idx)
  game   = o_move.ApplyTo(move.Idx, game)
  fmt.Printf("Play at (%d, %d), give #%d\n", o_move.Pos.i, o_move.Pos.j, o_move.Idx)

  move  = game.PlayWith(o_move.Idx)
  game  = move.ApplyTo(o_move.Idx, game)
  fmt.Printf("Play at (%d, %d), give #%d\n", move.Pos.i, move.Pos.j, move.Idx)
}
