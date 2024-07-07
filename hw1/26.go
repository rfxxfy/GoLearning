import (
	"fmt"
	"strings"
)

func anagram(data string, data2 string) bool {
	counterFirst := make(map[rune]int)
	counterSecond := make(map[rune]int)

	for _, char := range []rune(strings.ToLower(data)) {
		counterFirst[char]++
	}
	for _, char := range []rune(strings.ToLower(data2)) {
		counterSecond[char]++
	}

	if len(counterFirst) != len(counterSecond) {
		return false
	}
	for key, value := range counterFirst {
		if value != counterSecond[key] {
			return false
		}
	}
	return true
}
