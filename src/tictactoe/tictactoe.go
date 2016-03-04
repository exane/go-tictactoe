package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
  "math/rand"
  "time"
)

var board = Board{}

func main() {
  //go serverSetup()
  board.size = 3
  board.winCondition = 3
  board.start()
}

var players = []string{"X", "O"}

var bot = Bot{1}
var bot2 = Bot{0}

type Board struct {
  field        []string
  currPlayer   int
  turn         int
  botOnly      bool
  size         int
  winCondition int
}

func (b *Board) start() {
  b.buildField()
  rand.Seed(time.Now().UnixNano())
  b.currPlayer = rand.Intn(2)
  b.turn = 1
  b.botOnly = true
  b.doTurn()
}

func (b *Board) buildField() {
  b.field = make([]string, b.size * b.size)
  for i := 0; i < len(b.field); i++ {
    b.field[i] = "_"
  }
}

func (b *Board) readPlayerInput() int {
  reader := bufio.NewReader(os.Stdin)
  fmt.Printf("Player %s set: ", players[b.currPlayer])
  input, _ := reader.ReadString('\n')
  i, ok := strconv.ParseInt(strings.Trim(input, "\n\r"), 10, 16)

  if ok != nil {
    fmt.Println(ok)
    return 0
  }
  return int(i)
}

func (b *Board) isTied() bool {
  if b.checkWinCondition(0) || b.checkWinCondition(1) {
    return false
  }
  for _, val := range b.field {
    if val == "_" {
      return false
    }
  }
  return true
}

func (b *Board) doTurn() {
  drawBoard(b)
  var i int

  if (b.currPlayer == bot.id) {
    i = bot.doTurn(b)
    fmt.Printf("Player %s set: %d\n", players[b.currPlayer], i)
  } else if b.botOnly && b.currPlayer == bot2.id {
    i = bot2.doTurn(b)
    fmt.Printf("Player %s set: %d\n", players[b.currPlayer], i)
  } else {
    i = b.readPlayerInput() - 1
  }

  if !b.isValidMove(i) {
    fmt.Println("Invalid Move!")
    b.doTurn()
    return
  }

  b.setBoard(i)

  if b.checkWinCondition(b.currPlayer) {
    drawBoard(b)
    fmt.Println(players[b.currPlayer], "won!")
    fmt.Println("Press enter to restart the game.")
    reader := bufio.NewReader(os.Stdin)
    reader.ReadString('\n')
    b.start()
    return
  }

  if b.isTied() {
    drawBoard(b)
    fmt.Println("Tied!")
    fmt.Println("Press enter to restart the game.")
    //reader := bufio.NewReader(os.Stdin)
    //reader.ReadString('\n')
    b.start()
    return
  }

  b.nextTurn()
}

func (b *Board) checkWinCondition(currPlayer int) bool {
  for i := range b.field {
    p := Point{}
    p.fromIndex(i, b.size)
    if b.checkWinDirections(p, currPlayer) {
      return true
    }
  }

  return false
}

const (
  RIGHT = iota
  DOWN
  DIAG_UP
  DIAG_DOWN
)

func (b *Board) checkWinDirections(p Point, currPlayer int) bool {
  return b.checkWinDirection(p, RIGHT, 0, currPlayer) ||
  b.checkWinDirection(p, DOWN, 0, currPlayer) ||
  b.checkWinDirection(p, DIAG_UP, 0, currPlayer) ||
  b.checkWinDirection(p, DIAG_DOWN, 0, currPlayer)
}
func (b *Board) checkWinDirection(p Point, direction, depth, currPlayer int) bool {
  player := players[currPlayer]

  if v, ok := b.getField(p); v != player || !ok {
    return false
  }

  if depth >= b.winCondition - 1 {
    return true
  }

  switch direction {
  case RIGHT:
    p.x++
    return b.checkWinDirection(p, RIGHT, depth + 1, currPlayer)
    break
  case DOWN:
    p.y++
    return b.checkWinDirection(p, DOWN, depth + 1, currPlayer)
    break
  case DIAG_UP:
    p.x++
    p.y++
    return b.checkWinDirection(p, DIAG_UP, depth + 1, currPlayer)
    break
  case DIAG_DOWN:
    p.x++
    p.y--
    return b.checkWinDirection(p, DIAG_DOWN, depth + 1, currPlayer)
    break
  }
  return false
}

func (b *Board) isValidMove(i int) bool {
  if i < 0 || i >= len(b.field) {
    return false
  }
  return b.field[i] == "_"
}

func (b *Board) setBoard(i int) {
  b.field[i] = players[b.currPlayer]
}

func (b *Board) nextTurn() {
  b.turn++
  b.currPlayer++
  if b.currPlayer == 2 {
    b.currPlayer = 0
  }

  b.doTurn()
}

func (b *Board) getField(p Point) (string, bool) {
  if (p.x >= b.size || p.y >= b.size) {
    return ".", false
  }
  if (p.x < 0 || p.y < 0) {
    return ".", false
  }
  return b.field[p.toIndex(b.size)], true
}

func (b *Board) setField(x, y int, val string) {
  b.field[x + b.size * y] = val
}

type Point struct {
  x, y int
}

func (p *Point) toIndex(size int) int {
  return p.x + size * p.y
}
func (p *Point) fromIndex(index, size int) {
  p.x = index % size
  p.y = (index - p.x) / size
}
