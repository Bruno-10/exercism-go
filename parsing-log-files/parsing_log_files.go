package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
    allRegex := map[string]string{
        "trc": `^\[TRC\].*`,
        "dbg": `^\[DBG\].*`,
        "inf": `^\[INF\].*`,
        "wrn": `^\[WRN\].*`,
        "err": `^\[ERR\].*`,
        "ftl": `^\[FTL\].*`,
    }

    correctFormatting := false;

    for _,v := range allRegex {
        re, err := regexp.Compile(v)
        if err != nil {
            panic(err)
        }
        if re.MatchString(text) {
            correctFormatting = true
        }
    }
    return correctFormatting
}

func SplitLogLine(text string) []string {
    re, err := regexp.Compile(`<([\**]|[~*]|[-*]|[=*])*>`)
    if err != nil {
        panic(err)
    }
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    re, err := regexp.Compile(`(?i)".*password.*"`)
    if err != nil {
        panic(err)
    }

    returnCount := 0

    for i := 0; i < len(lines) ; i++ {
        if re.MatchString(lines[i]) {
            returnCount++
        }
    }
    return returnCount
}

func RemoveEndOfLineText(text string) string {
    re, err := regexp.Compile(`end-of-line.\d*`)
    if err != nil {
        panic(err)
    }

    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    re, err := regexp.Compile(`User\s+(\w+)`)
    if err != nil {
        panic(err)
    }

    for i := 0; i < len(lines) ; i++ {
        isAMatch := re.FindStringSubmatch(lines[i])
        if len(isAMatch) > 1 && isAMatch[1] != "" {
            lines[i] =  "[USR] " + isAMatch[1] + " " + lines[i]
        }
    }
    return lines
}
