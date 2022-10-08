package storage

type store struct {
	Paths []Path
}

type Path struct {
	OriginalPath string
	DotFilesPath string
	Taint        bool
}
