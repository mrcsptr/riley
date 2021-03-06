package filesystem

import (
	"strings"
)

// ParsePosition checks the position occupied by the dude for the selected entry,
// and returns an error if character is missing or not what was expected.
func parsePosition(content string) (string, error) {
	switch {
	case strings.EqualFold(content, "a"):
		return "A", nil
	case strings.EqualFold(content, "m"):
		return "M", nil
	case strings.EqualFold(content, "d"):
		return "D", nil
	default:
		return "", newInvalidPositionError(content)
	}
}
