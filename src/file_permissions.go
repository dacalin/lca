package main

import (
	"os"
	"strconv"
)

type FilePermissions struct {
	decimal uint32
	octal   string
	str     string
}

func  NewFilePermissions(perm uint32, isDir bool) FilePermissions {

	fileMode := os.FileMode(perm)

	if isDir {
		fileMode = fileMode | os.ModeDir
	}

	return FilePermissions{
		decimal: perm,
		octal:   strconv.FormatUint(uint64(perm)&511, 8), // 511 is 777 in octal
		str:     fileMode.String(),
	}
}

func (f FilePermissions) Octal() string {
	return f.octal
}

func (f FilePermissions) String() string {
	return f.str
}
