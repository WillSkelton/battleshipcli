package ship

import (
	"battleship/backend/cli/player"
	"fmt"
)

const (
	ALIVE = "ALIVE"
	HIT   = "HIT"
	DEAD  = "DEAD"

	VERT  = 'V'
	HORIZ = 'H'

	BOARDWIDTH = 10

	CARRIER    = "Carrier"
	BATTLESHIP = "Battleship"
	DESTROYER  = "Destroyer"
	SUBMARINE  = "Submarine"
	PTB        = "Patrol Boat"
)

type Ship struct {
	DisplayChar rune
	ShipType    string
	Length      int
	Location    [][]int
	Orientation rune
	Row         int
	Col         int
	Hits        []int
	Health      int
	Status      string
	Owner       *player.Player
}

var AddShip = map[string]func(*player.Player, int, int, rune) (*Ship, error){
	CARRIER:    NewCarrier,
	BATTLESHIP: NewBattleship,
	DESTROYER:  NewDestroyer,
	SUBMARINE:  NewSubmarine,
	PTB:        NewPTB,
}

// NewCarrier makes a New Carrier struct
func NewCarrier(p *player.Player, row, col int, orientation rune) (c *Ship, err error) {
	c = &Ship{}
	c.DisplayChar = 'C'
	c.ShipType = "Carrier"
	c.Length = 5
	c.Location = make([][]int, 5)
	c.Hits = make([]int, 5)
	c.Health = 5
	c.Status = ALIVE
	c.Orientation = orientation
	c.Owner = p

	if err = c.Place(row, col, c.Orientation); err != nil {
		return nil, err
	}
	return
}

// NewBattleship makes a New Battleship struct
func NewBattleship(p *player.Player, row, col int, orientation rune) (b *Ship, err error) {
	b = &Ship{}
	b.DisplayChar = 'B'
	b.ShipType = "Battleship"
	b.Length = 4
	b.Location = make([][]int, 4)
	b.Hits = make([]int, 4)
	b.Health = 4
	b.Status = ALIVE
	b.Orientation = orientation
	b.Owner = p

	if err = b.Place(row, col, b.Orientation); err != nil {
		return nil, err
	}
	return
}

// NewDestroyer makes a New Destroyer struct
func NewDestroyer(p *player.Player, row, col int, orientation rune) (d *Ship, err error) {
	d = &Ship{}
	d.DisplayChar = 'D'
	d.ShipType = "Destroyer"
	d.Length = 3
	d.Location = make([][]int, 3)
	d.Hits = make([]int, 3)
	d.Health = 3
	d.Status = ALIVE
	d.Orientation = orientation
	d.Owner = p

	if err = d.Place(row, col, d.Orientation); err != nil {
		return nil, err
	}
	return
}

// NewSubmarine makes a New Submarine struct
func NewSubmarine(p *player.Player, row, col int, orientation rune) (s *Ship, err error) {
	s = &Ship{}
	s.DisplayChar = 'S'
	s.ShipType = "Submarine"
	s.Length = 3
	s.Location = make([][]int, 3)
	s.Hits = make([]int, 3)
	s.Health = 3
	s.Status = ALIVE
	s.Orientation = orientation
	s.Owner = p

	if err = s.Place(row, col, s.Orientation); err != nil {
		return nil, err
	}
	return
}

// NewPTB makes a New PTB struct
func NewPTB(p *player.Player, row, col int, orientation rune) (ptb *Ship, err error) {
	ptb = &Ship{}
	ptb.DisplayChar = 'P'
	ptb.ShipType = "Patrol Boat"
	ptb.Length = 2
	ptb.Location = make([][]int, 2)
	ptb.Hits = make([]int, 2)
	ptb.Health = 2
	ptb.Status = ALIVE
	ptb.Orientation = orientation
	ptb.Owner = p

	if err = ptb.Place(row, col, ptb.Orientation); err != nil {
		return nil, err
	}
	return
}

// Place will fill a ship's location array with the squares it takes up.
func (s *Ship) Place(startRow, startCol int, orientation rune) (err error) {

	if !s.canPlace(startRow, startCol, orientation) {
		return fmt.Errorf("can't place ship there")
	}

	if orientation == VERT {
		for i := 0; i < s.Length; i++ {
			s.Location[i] = []int{startRow + i, startCol}
		}
	} else {
		for i := 0; i < s.Length; i++ {
			s.Location[i] = []int{startRow, startCol + i}
		}
	}

	return nil

}

func (s *Ship) canPlace(startRow, startCol int, orientation rune) bool {

	if startRow >= BOARDWIDTH || startRow < 0 {
		return false
	}
	if startCol >= BOARDWIDTH || startCol < 0 {
		return false
	}

	if orientation == VERT {
		if startRow > BOARDWIDTH-s.Length {
			return false
		}
	} else {
		if startCol > BOARDWIDTH-s.Length {
			return false
		}
	}
	return true
}

// PrintLocation prints out the coordinates that the ship occupies
func (s *Ship) PrintLocation() {
	for i := 0; i < s.Length; i++ {
		fmt.Printf("( %v, %v ) | ", s.Location[i][0], s.Location[i][1])
	}
	fmt.Println()
}

// Hit will update a ship's health/status when hit
func (s *Ship) Hit(row, col int) {
	for i, coordinate := range s.Location {
		x, y := coordinate[0], coordinate[1]
		if row == x && col == y {
			if s.Hits[i] == 1 {
				break
			}
			s.Hits[i] = 1
			s.Health--
			s.Status = HIT
			if s.Health == 0 {
				s.Status = DEAD
			}
			break
		}
	}
}
