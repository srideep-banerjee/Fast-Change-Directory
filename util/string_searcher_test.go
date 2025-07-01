package util

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
        if !testEq(t, actual, test.expected) {
            t.Error("String searcher gives [" + arrayToString(t, actual, ", ") + "] for \"" + test.input + "\" when it should be [" + arrayToString(t, test.expected, ", ") + "]")
        }
    }
}

func TestStringSearcherWithInitialization(t *testing.T) {
    var stringSearcher *StringSearcher[int] = NewStringSearcherWith([]StringSearcherEntry[int] {
        {key: "ab", value: 1},
        {key: "cd", value: 2},
    })
    actual := stringSearcher.GetAvailableValues("")
    if !testEq(t, actual, []int{1, 2}) {
        t.Error("String searcher with values initialization gives [" + arrayToString(t, actual, ", ") + "] instead of [1, 2]")
    }
}

func testEq(_ testing.TB, a, b []int) bool {
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

func arrayToString(_ testing.TB, a []int, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
