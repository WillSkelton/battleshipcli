package board

import (
	"battleship/backend/cli/player"
	"battleship/backend/cli/ship"
	"fmt"
)

type Board struct {
	GameMap map[int]map[int]*ship.Ship

	Player1 *player.Player
	Player2 *player.Player
}

const (
	BOARDWIDTH = 10
)

func NewBoard(p1Name, p2Name string) (b *Board, err error) {
	b = &Board{}

	if b.Player1, err = player.NewPlayer(p1Name); err != nil {
		return nil, err
	}

	if b.Player2, err = player.NewPlayer(p2Name); err != nil {
		return nil, err
	}

	b.GameMap = make(map[int]map[int]*ship.Ship)

	return b, nil
}

func (b *Board) AddShip(p *player.Player, shipType string, startRow, startCol int, orientation rune) (s *ship.Ship, err error) {
	if s, err = ship.AddShip[shipType](p, startRow, startCol, orientation); err != nil {
		return nil, err
	}

	b.spaceOccupied(s)

	for _, location := range s.Location {

		row, col := location[0], location[1]

		if b.GameMap[col] == nil {
			b.GameMap[col] = make(map[int]*ship.Ship)
		}
		b.GameMap[col][row] = s

	}

	return s, nil
}

func (b *Board) PrintBoard() {
	for row := 0; row < BOARDWIDTH; row++ {
		for col := 0; col < BOARDWIDTH; col++ {

			if _, exists := b.GameMap[col][row]; exists {
				fmt.Printf("%v ", string(b.GameMap[col][row].DisplayChar))
			} else {
				fmt.Printf("~ ")
			}
		}
		fmt.Println()
	}
}

func (b *Board) spaceOccupied(s *ship.Ship) bool {
	for _, loc := range s.Location {
		row, col := loc[0], loc[1]
		if _, exists := b.GameMap[col][row]; exists {
			return false
		}
	}
	return true
}
