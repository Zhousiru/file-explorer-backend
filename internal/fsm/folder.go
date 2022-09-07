package fsm

import (
	"os"
	"path"
	"path/filepath"

	"github.com/Zhousiru/file-explorer-backend/internal/config"
)

const (
	FLAG_FILE = 1 << iota
	FLAG_FOLDER
)

type FolderSub struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	IsFolder bool   `json:"isFolder"`
}

type Folder struct {
	File
}

func (f *Folder) GetSub(flag int) ([]*FolderSub, error) {
	sub, err := os.ReadDir(f.fullPath)
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

		fullSubPath := path.Join(f.fullPath, v.Name())

		el := new(FolderSub)
		el.Name = v.Name()

		tempPath, err := filepath.Rel(config.Get(config.K_ROOT), fullSubPath)
		if err != nil {
			return nil, err
		}

		el.Path = filepath.ToSlash(tempPath)
		el.IsFolder = isDir

		list = append(list, el)
	}

	return list, nil
}
