package fileio

// FileExists returns true if the specified file existing on the filesystem
func fileExists(filename string) bool {
	return Touch(filename)
}
