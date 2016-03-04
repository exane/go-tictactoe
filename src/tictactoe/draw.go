package main

import (
  "fmt"
  "os"
  "os/exec"
)

func drawBoard(board *Board) {
  clear()
  fmt.Printf("Turn: %d", board.turn)
  for i, val := range board.field {
    //fmt.Println(i)
    if ( i % board.size == 0) {
      fmt.Printf("\n")
    }
    fmt.Printf("%s ", val)
  }
  fmt.Printf("\n")
}

func clear() {
  cmd := exec.Command("cmd", "/c", "cls")
  cmd.Stdout = os.Stdout
  cmd.Run()
}