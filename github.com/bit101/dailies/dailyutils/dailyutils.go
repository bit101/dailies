package dailyutils

import "github.com/bit101/blgo/util"

// Filename returns the filename based on the directory and file type.
func Filename(filetype string) string {
	return util.ParentDir() + "." + filetype
}
