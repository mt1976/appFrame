package spinner

import (
	"fmt"
	"time"
)

var activeStyle Styles

type framesIndex int

type Spinner struct {
	// ...
	style    framesIndex
	row      int
	column   int
	frames   []string
	cycle    int
	sequence int
	slow     time.Duration
}

// new returns a new Spinner, with defaults
func new() *Spinner {

	activeStyle.initialiseStyles()
	// ...
	sp := &Spinner{style: activeStyle.Default, row: 0, column: 0}
	sp.frames = getFrames(sp.style)
	sp.cycle = len(sp.frames)
	sp.sequence = 0
	sp.slow = 0
	return sp
}

// tick advances the spinner to the next state
func (s *Spinner) tick(msg string) {
	// ...
	//	log.Println("tick")
	s.sequence = (s.sequence + 1)
	if s.sequence >= s.cycle {
		s.sequence = 0
	}
	//	log.Println("sequence:", s.sequence)
	fmt.Print("\033[u\033[K[" + s.frames[s.sequence] + "] " + msg)
	if s.slow > 0 {
		//		log.Println("sleeping")
		time.Sleep(s.slow)
	}
}

// setStyle sets the style of the spinner
func (s *Spinner) setStyle(style framesIndex) *Spinner {
	// ...
	s.style = style
	s.frames = getFrames(style)
	s.cycle = len(s.frames)
	s.sequence = 0
	return s
}

// setLocation sets the location of the spinner
func (s *Spinner) setLocation(row int, column int) *Spinner {
	// ...
	s.row = row
	s.column = column
	return s
}

// getFrames returns the characters for a given style
func getFrames(style framesIndex) []string {

	rtn := []string{}

	switch style {
	case activeStyle.Default:
		// Do nothinbg
	case activeStyle.Plus:
		rtn = []string{"+", "x"}
	case activeStyle.Directions:
		rtn = []string{"v", "<", "^", ">"}
	case activeStyle.Dots:
		rtn = []string{".   ", " .  ", "  . ", "   ."}
	case activeStyle.Ball:
		rtn = []string{"◐", "◓", "◑", "◒"}
	case activeStyle.SquareClock:
		rtn = []string{"◰", "◳", "◲", "◱"}
	case activeStyle.Clock:
		rtn = []string{"◴", "◷", "◶", "◵"}
	case activeStyle.Snake:
		rtn = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	case activeStyle.ChasingDots:
		rtn = []string{".  ", ".. ", "...", " ..", "  .", "   "}
	case activeStyle.Arrows:
		rtn = []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}
	case activeStyle.Grow:
		rtn = []string{"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃"}
	case activeStyle.Cross:
		rtn = []string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"}
	case activeStyle.Flip:
		rtn = []string{"_", "_", "_", "-", "`", "`", "'", "´", "-", "_", "_", "_"}
	case activeStyle.Cylon:
		rtn = []string{"( ●    )",
			"(  ●   )",
			"(   ●  )",
			"(    ● )",
			"(     ●)",
			"(    ● )",
			"(   ●  )",
			"(  ●   )",
			"( ●    )",
			"(●     )"}
	case activeStyle.DirectionsSlow:
		rtn = []string{"<", "<", "∧", "∧", ">", ">", "v", "v"}
	default:
		rtn = []string{"-", "\\", "|", "/"}
	}
	// ...
	return rtn
}

// initialiseStyles sets the default styles
func (s *Styles) initialiseStyles() *Styles {
	// ...
	s.Default = 1
	s.Plus = 2
	s.Directions = 3
	s.Dots = 4
	s.Ball = 5
	s.SquareClock = 6
	s.Clock = 7
	s.Snake = 8
	s.ChasingDots = 9
	s.Arrows = 10
	s.Grow = 11
	s.Cross = 12
	s.Flip = 13
	s.Cylon = 14
	s.DirectionsSlow = 15
	return s
}

// Delay sets the delay between frames
func (s *Spinner) setDelay(seconds float64) *Spinner {
	nanos := time.Second.Nanoseconds()
	seconds = float64(nanos) * seconds
	s.slow = time.Duration(seconds)
	return s
}
