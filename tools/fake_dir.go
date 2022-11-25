package tools

import (
	"io/fs"
	"time"
)

type FakeDir struct {
	FileName string
}

func (f FakeDir) Name() string {
	return f.FileName
}

func (f FakeDir) Size() int64 {
	//TODO implement me
	panic("implement me")
}

func (f FakeDir) Mode() fs.FileMode {
	//TODO implement me
	panic("implement me")
}

func (f FakeDir) ModTime() time.Time {
	//TODO implement me
	panic("implement me")
}

func (f FakeDir) IsDir() bool {
	return true
}

func (f FakeDir) Sys() any {
	//TODO implement me
	panic("implement me")
}
