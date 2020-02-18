package misc

import "os"

//Creates the directory if it doesn't already exist
func MakeDir(dirName string) {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		os.Mkdir(dirName, os.ModeDir)
	}
}
