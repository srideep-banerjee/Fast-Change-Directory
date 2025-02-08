package commands

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringSearcher(t *testing.T) {
    tests := []struct {
        add bool
        key string
        value int
        input string
        expected []int
    } {
        {add: false, input: "", expected: []int{} },
        {add: true, key: "", value: 0, input: "", expected: []int{0} },
        {add: true, key: "abcd", value: 1, input: "abcd", expected: []int{1} },
        {add: false, input: "abcde", expected: []int{} },
        {add: false, input: "ab", expected: []int{1} },
        {add: false, input: "cb", expected: []int{} },
        {add: true, key: "ab", value: 2, input: "a", expected: []int{1, 2} },
        {add: true, key: "cd", value: 3, input: "a", expected: []int{1, 2} },
        {add: false, input: "cb", expected: []int{} },
    }

    var stringSearcher *StringSearcher[int] = NewStringSearcher[int]()
    for _, test := range tests {
        if test.add {
            stringSearcher.AddEntry(test.key, test.value)
        }
        actual := stringSearcher.GetAvailableValues(test.input)
        if !testEq(actual, test.expected) {
            t.Error("String searcher gives [" + arrayToString(actual, ", ") + "] for \"" + test.input + "\" when it should be [" + arrayToString(test.expected, ", ") + "]")
        }
    }
}

func TestStringSearcherWithInitialization(t *testing.T) {
    var stringSearcher *StringSearcher[int] = NewStringSearcherWith(map[string]int {
        "ab": 1,
        "cd": 2,
    })
    actual := stringSearcher.GetAvailableValues("")
    if !testEq(actual, []int{1, 2}) {
        t.Error("String searcher with values initialization gives incorrect result")
    }
}

func testEq(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func arrayToString(a []int, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
