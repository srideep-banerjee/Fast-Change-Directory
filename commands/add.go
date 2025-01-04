package commands

import (
    "strings"
    "me/fast-cd/validation"
)

type add struct {}

func (a add) Matches(str string) bool {
    spaceInd := strings.IndexRune(str, ' ')

    if spaceInd != -1 {
        str = str[:spaceInd]
    }

    return strings.HasPrefix("add", str)
}

func (a add) Validate(str string) string {
    formatString := "Must be of the format: /add [tag]=[location]"
    str = str[strings.IndexRune(str, ' ') + 1:]
    equalInd := strings.IndexRune(str, '=')
    if equalInd == -1 {
        return formatString
    }
    tag := strings.Trim(str[:equalInd], " ")
    if valid, ch := validation.IsTagValid(tag); !valid {
        return string(ch) + " is not allowed in tags"
    }
    location := strings.Trim(str[equalInd + 1:], " ")
    if len(location) == 0 {
        return "Location cannot be empty"
    }
    return ""
}

func (a add) Process(str string) error {
    return nil
}
