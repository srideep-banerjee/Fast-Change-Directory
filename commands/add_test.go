package commands

import (
	"me/fast-cd/db"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddValidate(t *testing.T) {
	formatString := "Must be of the format: /add [tag]=[location]"
	tests := []struct {
		input string
		expected string
	} {
		{input: "", expected: formatString},
		{input: "/add", expected: formatString},
		{input: "/add ", expected: formatString},
		{input: "/add ade", expected: formatString},
		{input: "/add abe djejej", expected: formatString},
		{input: "/add abe=jrjfk$", expected: ""},
		{input: "/add      abe    = 	  jrjfk$", expected: ""},
		{input: "/add a$be=jrjfk$", expected: "$ is not allowed in tags"},
		{input: "/add abe=", expected: "Location cannot be empty"},
		{input: "/add abe= ", expected: "Location cannot be empty"},
		{input: "/add =", expected: "Tag cannot be empty"},
	}
	
	a := add{}
	for _, test := range tests {
		actual := a.Validate(test.input)
		assert.Equalf(t, test.expected, actual, "Input = " + strconv.Quote(test.input))
	}
}

func TestAddProcess(t *testing.T) {
	database, err := db.OpenSqliteDatabase("file:test.db?mode=memory")
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Cleanup(func() {database.Close()})
	
	err = database.AddTag("tag1", "location 1")
	if err != nil {
		t.Error(err)
		return
	}
	
	err = database.AddTag("tag2", "location 2")
	if err != nil {
		t.Error(err)
		return
	}
	
	err = database.AddTag("tag3", "location 1")
	if err != nil {
		t.Error(err)
		return
	}
	
	locations, err := database.ListTags("")
	if err != nil {
		t.Error(err)
		return
	}
	
	expected := [] db.Tag {
		db.NewTag("tag1", "location 1"),
		db.NewTag("tag2", "location 2"),
		db.NewTag("tag3", "location 1"),
	}
	
	assert.Equal(t, expected, locations)
	
}