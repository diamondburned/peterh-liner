//go:build js
// +build js

package liner

type termios = noopMode

func TerminalSupported() bool {
	return false
}

type noopMode struct{}

func (n noopMode) ApplyMode() error {
	return nil
}

// TerminalMode returns a noop InputModeSetter on this platform.
func TerminalMode() (ModeApplier, error) {
	return noopMode{}, nil
}

func initLinerTerminal(s *State) {
	s.inputRedirected = true
	s.outputRedirected = true
}

func (s *State) supportedStartPrompt() {}

func (s *State) supportedStopPrompt() {}

func (s *State) close() {}
