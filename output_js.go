//go:build js
// +build js

package liner

// GetWinSize is a custom function that returns the row and column size of the
// terminal. By default, it returns 0, 0, false. This function is meant to be
// overridden by the user.
var GetWinSize = func() (uint16, uint16, bool) {
	return 0, 0, false
}

const cursorColumn = false // unsure

func (s *State) getColumns() bool {
	_, col, ok := GetWinSize()
	if !ok {
		return false
	}
	s.columns = int(col)
	return true
}
