package db

import (
	"fmt"
	"me/fast-cd/util"
	"os"
	"path/filepath"
)

type Database interface {
	AddTag(tag string, location string) error
	FetchTag(tag string) (location string, err error)
	ListTags(prefix string) ([]Tag, error)
	Close() error
}

func OpenDatabase() (Database, error) {
	execPath := util.GetBinaryPath()

	err := os.MkdirAll(filepath.Join(execPath, "sqlite"), os.ModePerm)
	if err != nil {
		return nil, err
	}

	file := filepath.Join(execPath, "sqlite", "tags.db")
	return OpenSqliteDatabase("file:" + file + "?mode=rwc")
}

type TagNotFoundErr struct {
	tag string
}

func (tnfe TagNotFoundErr) Error() string {
	return fmt.Sprintf("Tag %s wasn't found", tnfe.tag)
}

type Tag struct {
	tag string
	location string
}

func NewTag(tag string, location string) Tag {
	return Tag{tag: tag, location: location}
}

func (t Tag) GetTagName() string {
	return t.tag
}

func (t Tag) GetLocation() string {
	return t.location
}