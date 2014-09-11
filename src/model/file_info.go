package model

import (
	"os"
	"time"
)

type FileInfo struct {
	Name       string
	Size       int64
	Modifytime time.Time
	Perm       os.FileMode
	Md5        string
	Type       bool
}
