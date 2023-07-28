package spinner

import (
	"log"
	"time"
)

type Styles struct {
	// ...
	Default        FramesIndex
	Plus           FramesIndex
	Directions     FramesIndex
	Dots           FramesIndex
	Ball           FramesIndex
	SquareClock    FramesIndex
	Clock          FramesIndex
	Snake          FramesIndex
	ChasingDots    FramesIndex
	Arrows         FramesIndex
	Grow           FramesIndex
	Cross          FramesIndex
	Flip           FramesIndex
	Cylon          FramesIndex
	DirectionsSlow FramesIndex
}

type FramesIndex int

func New() *Spinner {
	// ...
	return new()
}

func (s *Spinner) Tick() {
	s.tick("")
}

func (s *Spinner) TickWithMessage(message string) {
	s.tick(message)
}

func (s *Spinner) Style(style FramesIndex) *Spinner {
	// ...

	return s.setStyle(style)
}

//func (s *Spinner) SetLocation(row int, column int) *Spinner {
//	// ...
//	return s.setLocation(row, column)
//}

func (s *Spinner) Debug() {
	// ...

	log.Println("Debug")
	log.Println("style:", s.style)
	//log.Println("row:", s.row)
	//log.Println("column:", s.column)
	log.Println("frames:", s.characters)
	log.Println("speed:", s.slow)

}

func (s *Spinner) Delay(seconds float64) {
	nanos := time.Second.Nanoseconds()
	seconds = float64(nanos) * seconds
	s.setDelay(time.Duration(seconds))
}
