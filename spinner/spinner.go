package spinner

import (
	"fmt"
	"time"
)

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
	Styles   *spinnerStyles
}

type spinnerStyles struct {
	// ...
	Default        framesIndex
	Plus           framesIndex
	Directions     framesIndex
	Dots           framesIndex
	Ball           framesIndex
	SquareClock    framesIndex
	Clock          framesIndex
	Snake          framesIndex
	ChasingDots    framesIndex
	Arrows         framesIndex
	Grow           framesIndex
	Cross          framesIndex
	Flip           framesIndex
	Cylon          framesIndex
	DirectionsSlow framesIndex
}

// new returns a new Spinner, with defaults
func new() *Spinner {
	sp := &Spinner{row: 0, column: 0}
	sp.sequence = 0
	sp.slow = 0
	sp.Styles = initialiseStyles()
	sp.style = sp.Styles.Default
	sp.frames = sp.getFrames(sp.style)
	sp.cycle = len(sp.frames)
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
	s.frames = s.getFrames(style)
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
func (s *Spinner) getFrames(style framesIndex) []string {

	rtn := []string{"-", "\\", "|", "/"}

	switch style {
	case s.Styles.Default:
	//	rtn = []string{"-", "\\", "|", "/"}
	case s.Styles.Plus:
		rtn = []string{"+", "x"}
	case s.Styles.Directions:
		rtn = []string{"v", "<", "^", ">"}
	case s.Styles.Dots:
		rtn = []string{".   ", " .  ", "  . ", "   ."}
	case s.Styles.Ball:
		rtn = []string{"◐", "◓", "◑", "◒"}
	case s.Styles.SquareClock:
		rtn = []string{"◰", "◳", "◲", "◱"}
	case s.Styles.Clock:
		rtn = []string{"◴", "◷", "◶", "◵"}
	case s.Styles.Snake:
		rtn = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	case s.Styles.ChasingDots:
		rtn = []string{".  ", ".. ", "...", " ..", "  .", "   "}
	case s.Styles.Arrows:
		rtn = []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}
	case s.Styles.Grow:
		rtn = []string{"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃"}
	case s.Styles.Cross:
		rtn = []string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"}
	case s.Styles.Flip:
		rtn = []string{"_", "_", "_", "-", "`", "`", "'", "´", "-", "_", "_", "_"}
	case s.Styles.Cylon:
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
	case s.Styles.DirectionsSlow:
		rtn = []string{"<", "<", "∧", "∧", ">", ">", "v", "v"}
	default:
		//rtn = []string{"-", "\\", "|", "/"}
	}
	// ...
	return rtn
}

// initialiseStyles sets the default styles
func initialiseStyles() *spinnerStyles {
	// ...
	s := &spinnerStyles{}
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
