package envpath

import (
	"os"
	"path/filepath"
	"strings"
)

// Append adds path to the end of the provided string.
//
//If the path to add already exists in the provided string, it will be moved to
//the end.
func Append(envVar, newPath string) string {
	paths := filepath.SplitList(envVar)
	paths = removeStringFromSlice(newPath, paths)
	paths = append(paths, newPath)
	return strings.Join(paths, string(os.PathListSeparator))
}

// Prepend inserts path at the start of the provided string.
//
//If the path to add already exists in the provided string, it will be moved to
//the start.
func Prepend(envVar, newPath string) string {
	paths := filepath.SplitList(envVar)
	paths = removeStringFromSlice(newPath, paths)
	paths = append([]string{newPath}, paths...)
	return strings.Join(paths, string(os.PathListSeparator))
}

func removeStringFromSlice(s string, elems []string) []string {
	updated := []string{}
	for _, v := range elems {
		if v != s {
			updated = append(updated, v)
		}
	}
	return updated
}
