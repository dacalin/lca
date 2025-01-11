package _fileinfo

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash"
	"io"
	"os"
	"time"
)

type FileInfo struct {
    Permissions string
    Owner       string
    Group       string
    Size        int64
    ModTime     time.Time
    Name        string
    Hash        string
    IsDir       bool
}

func NewFileInfo(permissions, owner, group string, size int64, modTime time.Time, name string, isDir bool, hashType string) FileInfo {
    fileInfo := FileInfo{
        Permissions: permissions,
        Owner:       owner,
        Group:       group,
        Size:        size,
        ModTime:     modTime,
        Name:        name,
        IsDir:       isDir,
    }

    if hashType != "none" {
        fileInfo.Hash = fileInfo.calculateHash(hashType)
    }

    return fileInfo
}

func (f FileInfo) calculateHash(hashType string) string {
    if f.IsDir {
        return ""
    }

    var hasher hash.Hash
    if hashType == "md5" {
        hasher = md5.New()
    } else {
        hasher = sha1.New()
    }

    file, err := os.Open(f.Name)
    if err != nil {
        return "error"
    }
    defer file.Close()

    if _, err := io.Copy(hasher, file); err != nil {
        return "error"
    }

    return fmt.Sprintf("%x", hasher.Sum(nil))
}

func (f FileInfo) Display(showAll bool, level uint8) {
  d := NewDisplay(showAll, "02 Jan 06 15:04")

  d.Print(f, level)
}
