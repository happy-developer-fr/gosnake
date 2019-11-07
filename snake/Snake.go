package snake

import (
	"errors"
	"fmt"
	"github.com/tettoc/gosnake/direction"
)

type Snake []Body

const (
	NONE = iota
	HEAD
	BODY
)

type Body struct {
	D    direction.Direction
	X, Y int
}

func newSnake(head Body) Snake {
	return append(Snake{}, head)
}

func (s Snake) Is(x int, y int) int {
	for i, body := range s {
		if body.X == x && body.Y == y {
			if i == 0 {
				return HEAD
			} else {
				return BODY
			}
		}
	}
	return NONE
}
func (s Snake) ChangeDirection(d direction.Direction) {
	s[0].D = d
}

func (s Snake) Move(b Board) (Snake, error) {
	head := s.head()
	switch head.D {
	case direction.North:
		head.Y = (head.Y - 1) % len(b.Cells)
		if head.Y == -1 {
			head.Y = len(b.Cells) - 1
		}
	case direction.South:
		head.Y = (head.Y + 1) % len(b.Cells)
	case direction.East:
		head.X = (head.X + 1) % len(b.Cells)
	case direction.West:
		head.X = (head.X - 1) % len(b.Cells)
		if head.X == -1 {
			head.X = len(b.Cells) - 1
		}
	}
	newSnake := newSnake(head)
	var copySize int
	if b.HasFood(head.X, head.Y) {
		copySize = len(s)
		b.Cells[head.Y][head.X].Food = false
	} else {
		copySize = len(s) - 1
	}
	for _, body := range s[:copySize] {
		newSnake = append(newSnake, body)
	}
	//fmt.Print("head : ", head.X, head.Y)
	var err = error(nil)
	if s.looped() {
		err = errors.New(fmt.Sprintf("Game over, your score Is : %d", len(s)))
	}
	return newSnake, err
}

func (s Snake) head() Body {
	return s[0]
}

func (s Snake) tail() Body {
	return s[len(s)-1]
}

func (s Snake) looped() bool {
	head := s.head()
	for idx, bodyElement := range s {
		if head.X == bodyElement.X && head.Y == bodyElement.Y && idx != 0 {
			return true
		}
	}
	return false
}

func (s Snake) Print() {
	for i, _ := range s {
		switch i {
		case 0:
			fmt.Print(">")
		default:
			fmt.Print("-")
		}
	}
}
