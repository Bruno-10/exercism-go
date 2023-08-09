package logs

import (
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
 applicationCodes := map[int32]string{
    10071: "recommendation",
    128269: "search",
    9728: "weather",
}

for _, v := range log {
    if applicationCodes[v] != "" {
        return applicationCodes[v]
    }
}
return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    out := []rune(log)

    for i, v := range out {
        if v == oldRune {
            out[i] = newRune
        }
    }
    return string(out)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    return utf8.RuneCountInString(log) <= limit
}
