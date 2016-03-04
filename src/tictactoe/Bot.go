package main


type Bot struct {
  id int
}

func (this Bot) calcTurn(field []string, depth int, maximizing bool, alpha, beta int) int {
  b := Board{}
  b.field = make([]string, len(field))
  copy(b.field, field)

  if b.checkWinCondition(this.id) {
    return 10 - depth
  }
  if b.checkWinCondition(1 ^ this.id) {
    return -10 + depth
  }
  if b.isTied() {
    return 0
  }

  score := make(map[int]int)
  var v int

  if maximizing {
    v = -9000
    b.currPlayer = this.id
  } else {
    v = 9000
    b.currPlayer = 1 ^ this.id
  }

  var bestMove int
  for i := range field {
    b.field = make([]string, len(field))
    copy(b.field, field)

    if !b.isValidMove(i) {
      continue
    }

    b.setBoard(i)

    score[i] = this.calcTurn(b.field, depth + 1, !maximizing, alpha, beta)

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

func (this Bot) doTurn(b Board) int {
  return this.calcTurn(b.field, 0, true, -9000, 9000)
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