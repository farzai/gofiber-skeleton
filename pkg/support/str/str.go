package str

import "strings"

func SnakeCase(str string) string {
	return snakeCase(str)
}

// snakeCase converts a string to snake_case
// Ref: https://stackoverflow.com/questions/56616196/how-to-convert-camel-case-string-to-snake-case
func snakeCase(str string) string {
	var b strings.Builder
    diff := 'a' - 'A'
    l := len(str)
    for i, v := range str {
        // A is 65, a is 97
        if v >= 'a' {
            b.WriteRune(v)
            continue
        }

        // v is capital letter here
        // irregard first letter
        // add underscore if last letter is capital letter
        // add underscore when previous letter is lowercase
        // add underscore when next letter is lowercase
        if (i != 0 || i == l-1) && ( // head and tail
            (i > 0 && rune(str[i-1]) >= 'a') || // pre
            (i < l-1 && rune(str[i+1]) >= 'a')) { //next
            b.WriteRune('_')
        }

        b.WriteRune(v + diff)
    }

    return b.String()
}
