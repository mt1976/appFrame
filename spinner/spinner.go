package spinner

import (
	"fmt"
	"time"
)

var Style Styles

type Spinner struct {
	// ...
	style      FramesIndex
	row        int
	column     int
	characters []string
	cycle      int
	sequence   int
	slow       time.Duration
}

func new() *Spinner {

	Style.initialiseStyles()
	// ...
	sp := &Spinner{style: Style.Default, row: 0, column: 0}
	sp.characters = getCharacters(sp.style)
	sp.cycle = len(sp.characters)
	sp.sequence = 0
	sp.slow = 0
	return sp
}

func (s *Spinner) tick(msg string) {
	// ...
	//	log.Println("tick")
	s.sequence = (s.sequence + 1)
	if s.sequence >= s.cycle {
		s.sequence = 0
	}
	//	log.Println("sequence:", s.sequence)
	fmt.Print("\033[u\033[K[" + s.characters[s.sequence] + "] " + msg)
	if s.slow > 0 {
		//		log.Println("sleeping")
		time.Sleep(s.slow)
	}
}

func (s *Spinner) setStyle(style FramesIndex) *Spinner {
	// ...
	s.style = style
	s.characters = getCharacters(style)
	s.cycle = len(s.characters)
	s.sequence = 0
	return s
}

func (s *Spinner) setLocation(row int, column int) *Spinner {
	// ...
	s.row = row
	s.column = column
	return s
}

func getCharacters(style FramesIndex) []string {

	rtn := []string{}

	switch style {
	case Style.Default:
		// Do nothinbg
	case Style.Plus:
		rtn = []string{"+", "x"}
	case Style.Directions:
		rtn = []string{"v", "<", "^", ">"}
	case Style.Dots:
		rtn = []string{".   ", " .  ", "  . ", "   ."}
	case Style.Ball:
		rtn = []string{"◐", "◓", "◑", "◒"}
	case Style.SquareClock:
		rtn = []string{"◰", "◳", "◲", "◱"}
	case Style.Clock:
		rtn = []string{"◴", "◷", "◶", "◵"}
	case Style.Snake:
		rtn = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	case Style.ChasingDots:
		rtn = []string{".  ", ".. ", "...", " ..", "  .", "   "}
	case Style.Arrows:
		rtn = []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}
	case Style.Grow:
		rtn = []string{"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃"}
	case Style.Cross:
		rtn = []string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"}
	case Style.Flip:
		rtn = []string{"_", "_", "_", "-", "`", "`", "'", "´", "-", "_", "_", "_"}
	case Style.Cylon:
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
	case Style.DirectionsSlow:
		rtn = []string{"<", "<", "∧", "∧", ">", ">", "v", "v"}
	default:
		rtn = []string{"-", "\\", "|", "/"}
	}
	// ...
	return rtn
}

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

func (s *Spinner) setDelay(periodInMillis time.Duration) *Spinner {
	// ...
	s.slow = periodInMillis
	return s
}
