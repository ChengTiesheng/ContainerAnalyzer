package attr

import "fmt"

func Quote(l []string) []string {
	var quoted []string

	for _, s := range l {
		quoted = append(quoted, fmt.Sprintf("%q", s))
	}

	return quoted
}

func In(list []string, el string) bool {
	return IndexOf(list, el) != -1
}

func IndexOf(list []string, el string) int {
	for i, x := range list {
		if el == x {
			return i
		}
	}
	return -1
}
