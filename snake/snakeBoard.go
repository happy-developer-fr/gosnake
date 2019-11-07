package snake

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Board struct {
	Cells [][]Cell
}

func (b Board) HasFood(col, line int) bool {
	if len(b.Cells) > col && len(b.Cells) > line {
		return b.Cells[line][col].Food
	}
	return false
}

func (b Board) countFeedeedCells() int {
	counter := 0
	for _, lineCells := range b.Cells {
		for _, cell := range lineCells {
			if cell.Food {
				counter++
			}
		}
	}
	return counter
}

func (b Board) RandomCellWithoutFood() (int, int) {
	rand.Seed(time.Now().UnixNano())
	for {
		randLine := rand.Intn(len(b.Cells))
		randCol := rand.Intn(len(b.Cells))
		if cell := b.Cells[randLine][randCol]; cell.Food == false {
			return randLine, randCol
		}
	}
}

func (b Board) Feed(maxFeed int) {
	if b.countFeedeedCells() < maxFeed {
		line, col := b.RandomCellWithoutFood()
		b.Cells[line][col].Food = true
	}
}

func (b Board) Print(s Snake) {
	for lineIdx, line := range b.Cells {
		fmt.Print("|")
		for colIdx, cell := range line {
			sbody := s.Is(colIdx, lineIdx)
			switch sbody {
			case HEAD:
				fmt.Print("+")
			case BODY:
				fmt.Print(".")
			case NONE:
				switch cell.Food {
				case true:
					fmt.Print("*")
				default:
					fmt.Print(" ")
				}
			}

		}
		fmt.Println("|")
	}
}

func (b Board) PrintJson() ([]byte, error) {
	if bytes, err := json.MarshalIndent(b, "", "\t"); err != nil {
		fmt.Printf("Error %q when marshalling Board %v", err, b)
		return nil, err
	} else {
		return bytes, nil
	}
}

func SnakeBoard(size int) Board {
	b := Board{Cells: make([][]Cell, size)}
	for line := 0; line < size; line++ {
		b.Cells[line] = make([]Cell, size)
		for col := 0; col < size; col++ {
			b.Cells[line][col] = Cell{Food: false}
		}
	}
	return b
}
