package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"os"
	"strconv"
	"time"
)


type FileInfo struct {
    isDir       bool
    name        string
    permissions FilePermissions
    owner       string
    group       string
    size        int64
    modTime     time.Time
    hash        string
    hashType    string
}

func NewFileInfo(permissions uint32, owner, group string, size int64, modTime time.Time, name string, isDir bool, hashType string) FileInfo {

    fileInfo := FileInfo{
        permissions: NewFilePermissions(permissions, isDir),
        owner:       owner,
        group:       group,
        size:        size,
        modTime:     modTime,
        name:        name,
        isDir:       isDir,
        hashType: hashType,
    }

    if hashType != "none" && !fileInfo.isDir {
        fileInfo.hash = fileInfo.calculateHash(hashType)
    }

    return fileInfo
}

func (f FileInfo) formatPermissions(perm uint32) string {
    // Convert the decimal value 493 to an octal string (i.e., "755")
    octStr := strconv.FormatUint(uint64(perm), 8)
    fileMode := os.FileMode(perm) | os.ModeDir
    fmt.Println(fileMode.String()) 

    return octStr
}

func (f FileInfo) calculateHash(hashType string) string {

    var hasher hash.Hash
    switch hashType {
    case "md5":
        hasher = md5.New()
    case "sha256":
        hasher = sha256.New()
    default:
        // default to SHA-1 if no recognized hash is provided
        hasher = sha1.New()
    }

    file, err := os.Open(f.name)
    if err != nil {
        return "error"
    }
    defer file.Close()

    if _, err := io.Copy(hasher, file); err != nil {
        return "error"
    }

    return fmt.Sprintf("%x", hasher.Sum(nil))
}

//getters
func (f FileInfo) IsDir() bool {
    return f.isDir
}

func (f FileInfo) Name() string {
    return f.name
}

func (f FileInfo) Permissions() FilePermissions {
    return f.permissions
}

func (f FileInfo) Owner() string {
    return f.owner
}

func (f FileInfo) Group() string {
    return f.group
}

func (f FileInfo) Size() int64 {
    return f.size
}

func (f FileInfo) ModTime() time.Time {
    return f.modTime
}

func (f FileInfo) Hash() string {
    return f.hash
}

func (f FileInfo) HashType() string {
    return f.hashType
}