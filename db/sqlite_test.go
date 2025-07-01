package db

import (
	"me/fast-cd/util"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenSqliteMemoryDatabase(t *testing.T) {
	db, err := OpenSqliteDatabase("file:sqlite/test.db?mode=memory")

	if err != nil {
		t.Error(err)
		return
	}

	err = db.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestOpenSqliteFileDatabase(t *testing.T) {

	file := filepath.Join(util.GetBinaryPath(), "..", "sqlite", "test.db")

	db, err := OpenSqliteDatabase("file:" + file + "?mode=rwc")
	
	if err != nil {
		t.Error(err, file)
		return
	}
	
	t.Cleanup(func() {
		db.Close()
		os.Remove(file)
	})
	
	assert.FileExists(t, file)
}

func TestAddAndListTag(t *testing.T) {
	db, err := OpenSqliteDatabase("file:sqlite/test.db?mode=memory")
	if err != nil {
		t.Error("Failed to establish connection: " + err.Error())
		return
	}
	
	defer db.Close()
	
	err = db.AddTag("tag1", "Dummy location1")
	if err != nil {
		t.Error("Failed to add tag1: " + err.Error())
		return
	}
	
	err = db.AddTag("Tag2", "Dummy location2")
	if err != nil {
		t.Error("Failed to add tag2: " + err.Error())
		return
	}
	
	err = db.AddTag("Ttg2", "Dum")
	if err != nil {
		t.Error("Failed to add tag2: " + err.Error())
		return
	}

	expected := []Tag{}
	expected = append(
		expected,
		Tag{tag: "tag1", location: "Dummy location1"},
		Tag{tag: "Tag2", location: "Dummy location2"},
	)
	
	tags, err := db.ListTags("ta")
	if err != nil {
		t.Error("Failed to list tags: " + err.Error())
		return
	}
	
	assert.Equal(t, expected, tags)
}