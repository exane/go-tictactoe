package main

import (
  "testing"
//"os"
//"fmt"
)

var board = Board{
  []string{
    "_", "_", "_",
    "_", "_", "_",
    "_", "_", "_",
  }, 0, 0,
}

func TestWinCondition(t *testing.T) {
  var testFields = []struct {
    field  []string
    expect bool
  }{
    {[]string{
      "X", "X", "X",
      "_", "_", "_",
      "_", "_", "_",
    }, true},
    {[]string{
      "_", "_", "_",
      "X", "X", "X",
      "_", "_", "_",
    }, true},
    {[]string{
      "_", "_", "_",
      "_", "_", "_",
      "X", "X", "X",
    }, true},
    {[]string{
      "X", "_", "_",
      "_", "X", "_",
      "_", "_", "X",
    }, true},
    {[]string{
      "_", "_", "X",
      "_", "X", "_",
      "X", "_", "_",
    }, true},
    {[]string{
      "X", "_", "_",
      "X", "_", "_",
      "X", "_", "_",
    }, true},
    {[]string{
      "_", "_", "X",
      "_", "X", "O",
      "_", "_", "X",
    }, false},
    {[]string{
      "_", "X", "_",
      "_", "X", "_",
      "X", "0", "X",
    }, false},
    {[]string{
      "X", "O", "O",
      "O", "X", "X",
      "O", "_", "O",
    }, false},
  }

  b := board
  b.currPlayer = 0

  for _, val := range testFields {
    b.field = val.field
    res := b.checkWinCondition(b.currPlayer)
    if res != val.expect {
      t.Error("Expected", val.expect, "got", res)
    }
  }
}

func TestMoveValidation(t *testing.T) {
  b := board

  b.field[1] = "O"

  var moves = []struct {
    input  int
    expect bool
  }{
    {0, true, },
    {8, true, },
    {-1, false, },
    {9, false, },
    {9, false, },
    {1, false, },
    {2, true, },
  }

  for _, val := range moves {
    if b.isValidMove(val.input) != val.expect {
      t.Error("expect", val.input, "to be", val.expect)
    }
  }

}

func TestNextTurn(t *testing.T) {
  b := board
  //turn, player := b.turn, b.currPlayer
  //b.doTurn  func() {}
  //b.nextTurn()

  if b.turn != 1 || b.currPlayer != 1 {
    //t.Error("expect turn to be 1 and player to be 1")
  }
}

func TestAssignment(t *testing.T) {
  field := []string{
    "_", "_", "_",
    "_", "_", "_",
    "_", "_", "_",
  }
  b := Board{}

  b.field = make([]string, len(field))
  copy(b.field, field)
  b.field[0] = "X"

  if field[0] != "_" {
    t.Error("Expected field[0] to be _ got", field[0])
  }

  if b.field[0] != "X" {
    t.Error("Expected b.field[0] to be X got", b.field[0])
  }

}

func TestMiniMax(t *testing.T) {
  b := Board{
    []string{
      "X", "O", "X",
      "X", "O", "_",
      "_", "_", "X",
    }, 1, 0,
  }
  bot := Bot{1}

  testTable := []struct {
    field  []string
    expect int
  }{
    /*{
      []string{
        "X", "O", "X",
        "X", "O", "_",
        "_", "_", "X",
      }, 7,
    },
    {
      []string{
        "_", "_", "X",
        "X", "_", "_",
        "_", "_", "X",
      }, 5,
    },*/
    {
      []string{
        "_", "_", "O",
        "O", "X", "X",
        "O", "X", "_",
      }, 0,
    },
    {
      []string{
        "_", "_", "O",
        "O", "X", "X",
        "_", "X", "O",
      }, 1,
    },
  }

  //bot.calcTurn(b.field, 0, true)

  for _, val := range testTable {
    b.field = val.field
    res := bot.doTurn(b)
    if val.expect != res {
      t.Error("Expected", val.expect, "got", res)
    }
  }

}