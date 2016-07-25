package lib

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type FileFInder interface {
	FindFiles(path string) ([]string, []string, error)
}

type FileWalker struct {
	StartTime time.Time
	EndTime   time.Time
	Recursive bool
	Include   bool
	Dir       bool
	DirOnly   bool
}

func (finder *FileWalker) FindFiles(root string) ([]string, []string, error) {
	if finder.Recursive {
		return finder.findFileRecursive(root)
	} else {
		return finder.findFile(root)
	}
}

func (finder *FileWalker) findFile(root string) ([]string, []string, error) {
	ret := []string{}
	retDir := []string{}
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, nil, err
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, nil, err
	}

	for _, f := range files {
		if !finder.Dir && f.IsDir() {
			continue
		}
		if finder.checkDate(f) {
			fp := filepath.Join(abs, f.Name())
			if f.IsDir() {
				retDir = append(retDir, fp)
			} else {
				if !finder.DirOnly {
					ret = append(ret, fp)
				}
			}
		}
	}
	return ret, retDir, nil
}

func (finder *FileWalker) findFileRecursive(root string) ([]string, []string, error) {
	ret := []string{}
	retDir := []string{}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if root == path {
			return nil
		}

		// If Dir param not be set. skip.
		if !finder.Dir && info.IsDir() {
			return nil
		}

		if err != nil {
			return err
		}
		if finder.checkDate(info) {
			fp, err := filepath.Abs(path)
			if err != nil {
				return err
			}

			if info.IsDir() {
				retDir = append(retDir, fp)
			} else {
				if !finder.DirOnly {
					ret = append(ret, fp)
				}
			}
		}
		return nil
	})
	return ret, retDir, err
}

func (finder *FileWalker) checkDate(info os.FileInfo) bool {
	mt := info.ModTime()
	if finder.Include {
		return mt.After(finder.StartTime) && mt.Before(finder.EndTime)
	} else {
		if finder.EndTime.After(time.Now()) {
			return mt.Before(finder.StartTime)
		} else {
			return mt.Before(finder.StartTime) && mt.After(finder.EndTime)
		}
	}
}

func NewFileFinder(startTime, endTime time.Time, recursive, include, dir, dironly bool) (FileFInder, error) {
	return &FileWalker{
		StartTime: startTime,
		EndTime:   endTime,
		Recursive: recursive,
		Include:   include,
		Dir:       dir,
		DirOnly:   dironly,
	}, nil
}
