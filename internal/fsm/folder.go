package fsm

import (
	"errors"
	"os"
	"path"

	errorCode "github.com/Zhousiru/file-explorer-backend/internal/error_code"
	"github.com/Zhousiru/file-explorer-backend/internal/util"
)

const (
	FLAG_FILE = 1 << iota
	FLAG_FOLDER
)

type FolderSub struct {
	Name     string
	Path     string
	IsFolder bool
}

type Folder struct {
	File
}

func (f *Folder) GetSub(flag int) ([]*FolderSub, error) {
	if !util.IsExist(f.Path) {
		return nil, errors.New(errorCode.FileNotFound)
	}

	sub, err := os.ReadDir(f.Path)
	if err != nil {
		return nil, err
	}

	list := []*FolderSub{}

	for _, v := range sub {
		isDir := v.IsDir()
		if isDir {
			// is folder
			if (flag & FLAG_FOLDER) == 0 {
				continue
			}
		} else {
			// is file
			if (flag & FLAG_FILE) == 0 {
				continue
			}
		}

		el := new(FolderSub)
		el.Name = v.Name()
		el.Path = path.Join(f.Path, v.Name())
		el.IsFolder = isDir

		list = append(list, el)
	}

	return list, nil
}
