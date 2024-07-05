import (
	"fmt"
	"strings"
)

func check(x string) bool {
	switch strings.ToLower(x) {
	case "a", "e", "i", "o", "u", "y":
		return true
	default:
		return false
	}
}
