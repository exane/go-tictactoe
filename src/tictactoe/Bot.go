package main


type Bot struct {
  id int
}

func (b *Bot) calcTurn(field []string, depth int, maximizing bool, alpha, beta int) int {
  board := board
  board.field = make([]string, len(field))
  copy(board.field, field)

  if board.checkWinCondition(b.id) {
    return 100 - depth
  }
  if board.checkWinCondition(1 ^ b.id) {
    return -100 + depth
  }
  if board.isTied() {
    return 0
  }
  if depth >= 8 {
    return 0
  }

  score := make(map[int]int)
  var v int

  if maximizing {
    v = -9000
    board.currPlayer = b.id
  } else {
    v = 9000
    board.currPlayer = 1 ^ b.id
  }

  var bestMove int
  for i := range field {
    board.field = make([]string, len(field))
    copy(board.field, field)

    if !board.isValidMove(i) {
      continue
    }

    board.setBoard(i)

    score[i] = b.calcTurn(board.field, depth + 1, !maximizing, alpha, beta)

    if maximizing {
      if score[i] > v {
        v = score[i]
        bestMove = i
      }
      alpha = max(alpha, v)
      if beta <= alpha {
        break
      }
    } else {
      if score[i] < v {
        v = score[i]
        bestMove = i
      }
      beta = min(beta, v)
      if beta <= alpha {
        break
      }
    }

  }

  if (depth == 0) {
    return bestMove
  }
  return v
}

func (b *Bot) doTurn(board *Board) int {
  return b.calcTurn(board.field, 0, true, -9000, 9000)
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}
func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}