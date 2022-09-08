package fsm

import (
	"errors"
	"os"
	"path"

	"github.com/Zhousiru/file-explorer-backend/internal/config"
	errorCode "github.com/Zhousiru/file-explorer-backend/internal/error_code"
	"github.com/Zhousiru/file-explorer-backend/internal/util"
)

type File struct {
	fullPath string
	filename string
}

func (f *File) SetPath(relPath string) error {
	fullPath := path.Join(config.Get(config.K_ROOT), relPath)
	if !util.IsExist(fullPath) {
		return errors.New(errorCode.FileNotFound)
	}

	f.fullPath = fullPath
	f.filename = path.Base(f.fullPath)
	return nil
}

func (f *File) Delete() error {
	return os.RemoveAll(f.fullPath)
}

func (f *File) Rename(newFilename string) error {
	dir := path.Dir(f.fullPath)
	newPath := path.Join(dir, newFilename)
	if util.IsExist(newPath) {
		return errors.New(errorCode.FileAlreadyExist)
	}

	return os.Rename(f.fullPath, newPath)
}

func (f *File) Move(newDir string) error {
	fullNewDir := path.Join(config.Get(config.K_ROOT), newDir)

	newFullPath := path.Join(fullNewDir, f.filename)

	if util.IsExist(newFullPath) {
		return errors.New(errorCode.FileAlreadyExist)
	}

	if !util.IsExist(fullNewDir) {
		// new dir not exist
		err := os.MkdirAll(fullNewDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return os.Rename(f.fullPath, newFullPath)
}

func (f *File) GetFullPath() string {
	return f.fullPath
}

func (f *File) GetFilename() string {
	return f.filename
}

func (f *File) GetFormattedModTime() (string, error) {
	info, err := os.Stat(f.fullPath)
	if err != nil {
		return "", err
	}
	return info.ModTime().Format("2006-01-02 15:04:05"), nil
}
