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

func main() {
    board := Board{}

    board.start()
}

var players = []string{"X", "O"}

type Bot struct {
    id int
}

func (this Bot) calcTurn(field []string, depth int, maximizing bool) int {
    b := Board{}
    b.field = make([]string, len(field))
    copy(b.field, field)

    if b.checkWinCondition(bot.id) {
        return 10 - depth
    }
    if b.checkWinCondition(1 ^ bot.id) {
        return -10 + depth
    }
    if b.isTied() {
        return 0
    }

    score := make(map[int]int)

    if maximizing {
        b.currPlayer = bot.id
    } else {
        b.currPlayer = 1 ^ bot.id
    }

    for i := range field {
        b.field = make([]string, len(field))
        copy(b.field, field)

        if !b.isValidMove(i) {
            continue
        }

        b.setBoard(i)

        score[i] = this.calcTurn(b.field, depth + 1, !maximizing)

    }
    var bestMove int
    var bestScore int

    if maximizing {
        bestScore = -9000
    } else {
        bestScore = 9000
    }

    for index, val := range score {
        if maximizing && bestScore < val {
            bestScore = val
            bestMove = index
        }
        if !maximizing && bestScore > val {
            bestScore = val
            bestMove = index
        }
    }
    if (depth == 0) {
        return bestMove
    }
    return bestScore
}

func (this Bot) doTurn(b Board) int {
    return this.calcTurn(b.field, 0, true)
}

var bot = Bot{1}

type Board struct {
    field      []string
    currPlayer int
    turn       int
}

func (b Board) start() {
    b.field = []string{
        "_", "_", "_",
        "_", "_", "_",
        "_", "_", "_",
    }
    rand.Seed(time.Now().UnixNano())
    b.currPlayer = rand.Intn(2)
    b.turn = 1
    b.doTurn()
}

func (b Board) readPlayerInput() int {
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

func (b Board) isTied() bool {
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

func (b Board) doTurn() {
    drawBoard(b)
    var i int

    if (b.currPlayer == bot.id) {
        //i = rand.Intn(9)
        i = bot.doTurn(b)
        fmt.Printf("Player %s set: %d\n", players[b.currPlayer], i)
    } else {
        i = b.readPlayerInput()
    }

    if b.isValidMove(i) == false {
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
        reader := bufio.NewReader(os.Stdin)
        reader.ReadString('\n')
        b.start()
        return
    }

    b.nextTurn()
}

func (b Board) checkWinCondition(currPlayer int) bool {
    for i := 0; i < 3; i++ {
        if b.field[0 + i * 3] == players[currPlayer] && b.field[1 + i * 3] == players[currPlayer] && b.field[2 + i * 3] == players[currPlayer] {
            return true
        }
        if b.field[0 + i] == players[currPlayer] && b.field[3 + i] == players[currPlayer] && b.field[6 + i] == players[currPlayer] {
            return true
        }
        if b.field[4] == players[currPlayer] &&
        (b.field[0] == players[currPlayer] && b.field[8] == players[currPlayer] ||
        b.field[2] == players[currPlayer] && b.field[6] == players[currPlayer]) {
            return true
        }
    }
    return false
}

func (b Board) isValidMove(i int) bool {
    if i < 0 || i > 8 {
        return false
    }
    return b.field[i] == "_"
}

func (b Board) setBoard(i int) {
    b.field[i] = players[b.currPlayer]
}

func (b Board) nextTurn() {
    b.turn++
    b.currPlayer++
    if b.currPlayer == 2 {
        b.currPlayer = 0
    }

    b.doTurn()
}
