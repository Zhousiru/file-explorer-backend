package fsm

import (
	"os"
	"path"
)

type File struct {
	Path string
}

func (f File) Delete() error {
	return os.Remove(f.Path)
}

func (f File) Rename(newFilename string) error {
	dir := path.Dir(f.Path)
	newPath := path.Join(dir, newFilename)

	return os.Rename(f.Path, newPath)
}

func (f File) Move(newDir string) error {
	filename := path.Base(f.Path)
	newPath := path.Join(newDir, filename)

	_, err := os.Stat(newDir)
	if err != nil {
		// new dir not exist
		err := os.MkdirAll(newDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return os.Rename(f.Path, newPath)
}
