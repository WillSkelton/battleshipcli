package main

import (
	"battleship/backend/cli/board"
	"battleship/backend/cli/ship"
	"log"
)

func main() {

	var (
		err error
		b   *board.Board
	)

	if b, err = board.NewBoard("Will", "PC"); err != nil {
		log.Fatal(err)
	}

	if _, err = b.AddShip(b.Player1, ship.CARRIER, 1, 1, ship.HORIZ); err != nil {
		log.Fatal(err)
	}

	b.PrintBoard()
}
