package main

import (
	"errors"
	"fmt"
	"github.com/tettoc/gosnake/direction"
	"github.com/tettoc/gosnake/key"
	"github.com/tettoc/gosnake/screenclear"
	"github.com/tettoc/gosnake/snake"
	"os"
	"strconv"
	"time"
)

func main() {
	size, err := getSize()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	chanReading := key.StartReading()
	snakeB := snake.SnakeBoard(size)
	snk := snake.Snake{}
	snk = append(snk, snake.Body{D: direction.East, X: 0, Y: 0})
	run := true
	for run {
		snakeB.Print(snk)
		var err = error(nil)
		select {
		case inst := <-chanReading:
			if !inst.Run {
				return
			}
			snk.ChangeDirection(inst.Dir)
		case <-time.After(1 * time.Millisecond):
		}
		if snk, err = snk.Move(snakeB); err != nil {
			key.Stop()
			println(err.Error())
			run = false
			return
		}
		time.Sleep(200 * time.Millisecond)
		screenclear.CleanScreen()

		snakeB.Feed(5)
	}
}

func getSize() (int, error) {
	const minsize = 10
	if len(os.Args) != 2 {
		message := fmt.Sprintf("Give a size for board (min %d)", minsize)
		return 0, errors.New(message)
	}

	if size, err := strconv.Atoi(os.Args[1]); err != nil || size < minsize {
		message := fmt.Sprintf("Size should be an integer greater than", minsize)
		return 0, errors.New(message)
	} else {
		return size, nil
	}
}
