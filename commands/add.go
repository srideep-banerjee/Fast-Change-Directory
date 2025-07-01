package commands

import (
	"errors"
	"me/fast-cd/db"
	"me/fast-cd/validation"
	"strings"
)

type add struct {}

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
    if len(tag) == 0 {
        return "Tag cannot be empty"
    }
    location := strings.Trim(str[equalInd + 1:], " ")
    if len(location) == 0 {
        return "Location cannot be empty"
    }
    return ""
}

func (a add) Process(database db.Database, str string) error {
    if str := a.Validate(str); str != "" {
        return errors.New(str)
    }

    str = str[strings.IndexRune(str, ' ') + 1:]
    equalInd := strings.IndexRune(str, '=')
    tag := strings.Trim(str[:equalInd], " ")
    location := strings.Trim(str[equalInd + 1:], " ")

    return database.AddTag(tag, location)
}
