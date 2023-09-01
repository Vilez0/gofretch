package util

import (
	"io/fs"
	"path/filepath"
)

// from https://stackoverflow.com/questions/55300117/how-do-i-find-all-files-that-have-a-certain-extension-in-go-regardless-of-depth
// FindFileWithExtension returns a slice of strings containing the paths of all files with the specified extension
func FindFileWithExtension(root string, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}
