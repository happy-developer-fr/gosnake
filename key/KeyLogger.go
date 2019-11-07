package key

import (
	term "github.com/nsf/termbox-go"
	"github.com/tettoc/gosnake/direction"
)

type instruction struct {
	Dir direction.Direction
	Run bool
}

func StartReading() chan instruction {
	c := make(chan instruction)

	//go func() {
	//	scanner := bufio.NewReaderSize(os.Stdin, 1)
	//	for {
	//		timer := time.NewTimer(time.Millisecond * 30)
	//		<-timer.C
	//
	//		ascii, err := scanner.ReadByte()
	//		if err != nil {
	//			fmt.Println(err.Error())
	//			return
	//		}
	//		switch ascii {
	//		case 68:
	//			c <- direction.West
	//		case 67:
	//			c <- direction.East
	//		case 66:
	//			c <- direction.South
	//		case 65:
	//			c <- direction.North
	//		}
	//	}
	//}()

	err := term.Init()
	if err != nil {
		panic(err)
	}
	run := true
	go func() {
		for run {
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				switch ev.Key {
				case term.KeyArrowUp:
					//reset()
					c <- instruction{direction.North, run}
				case term.KeyArrowDown:
					//reset()
					c <- instruction{direction.South, run}
				case term.KeyArrowLeft:
					//reset()
					c <- instruction{direction.West, run}
				case term.KeyArrowRight:
					//reset()
					c <- instruction{direction.East, run}
				case term.KeyCtrlC, term.KeyEsc:
					Stop()
					run = false
					c <- instruction{direction.North, run}
				}
			}
		}
	}()
	return c
}

func Stop() {
	reset()
	term.Close()
}
func reset() {
	term.Sync() // cosmestic purpose
}
