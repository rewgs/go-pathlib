package windows

import (
	"unicode"
)

// Using this instead of os.PathSeparator, which is always `/`.
const Separator string = "\\"

// Checks if a path begins with a drive letter, and returns it if true.
func GetDriveLetter(path string) (bool, string) {
	if !unicode.IsLetter(rune(path[0])) {
		return false, ""
	}
	if string(path[1]) != ":" {
		return false, ""
	}
	if string(path[2]) != "/" && string(path[2]) != "\\" {
		return false, ""
	}
	return true, path[:2]
}
