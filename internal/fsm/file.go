package fsm

import (
	"errors"
	"os"
	"path"

	errorCode "github.com/Zhousiru/file-explorer-backend/internal/error_code"
	"github.com/Zhousiru/file-explorer-backend/internal/util"
)

type File struct {
	Path string
}

func (f File) Delete() error {
	if util.IsExist(f.Path) {
		return errors.New(errorCode.FileNotFound)
	}

	return os.Remove(f.Path)
}

func (f File) Rename(newFilename string) error {
	if util.IsExist(f.Path) {
		return errors.New(errorCode.FileNotFound)
	}

	dir := path.Dir(f.Path)
	newPath := path.Join(dir, newFilename)
	if util.IsExist(newPath) {
		return errors.New(errorCode.FileAlreadyExist)
	}

	return os.Rename(f.Path, newPath)
}

func (f File) Move(newDir string) error {
	if util.IsExist(f.Path) {
		return errors.New(errorCode.FileNotFound)
	}

	filename := path.Base(f.Path)
	newPath := path.Join(newDir, filename)

	if util.IsExist(newPath) {
		return errors.New(errorCode.FileAlreadyExist)
	}

	if !util.IsExist(newDir) {
		// new dir not exist
		err := os.MkdirAll(newDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return os.Rename(f.Path, newPath)
}
