package spinner

import (
	"log"
)

// New returns a new Spinner
func New() *Spinner {
	// ...
	return new()
}

// Tick advances the spinner to the next state
func (s *Spinner) Tick() {
	s.tick("")
}

// TickWithMessage advances the spinner to the next state and displays a message
func (s *Spinner) TickWithMessage(message string) {
	s.tick(message)
}

// Style sets the style of the spinner
func (s *Spinner) Style(style framesIndex) *Spinner {
	return s.setStyle(style)
}

//func (s *Spinner) SetLocation(row int, column int) *Spinner {
//	// ...
//	return s.setLocation(row, column)
//}

// Debug prints debug information to stdout
func (s *Spinner) Debug() {
	log.Println("Debug")
	log.Println("style:", s.style)
	//log.Println("row:", s.row)
	//log.Println("column:", s.column)
	log.Println("frames:", s.frames)
	log.Println("speed:", s.slow)
	log.Println("Styles:", s.Styles)
	log.Println("cycle:", s.cycle)

}

// Delay sets the delay between frames
func (s *Spinner) Delay(seconds float64) *Spinner {
	s.setDelay(seconds)
	return s
}
