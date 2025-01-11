package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"syscall"

	_fileinfo "github.com/dacalin/lca/fileinfo"
)

func main() {
	hashType := flag.String("h", "none", "hash type: md5 or sha1")
	showAll := flag.Bool("p", false, "show permissions and owners")
	recursionDepth := flag.Uint("r", 0, "recursion depth (0-5)")
	flag.Parse()

	if *recursionDepth < 0 || *recursionDepth > 5 {
		log.Fatalf("Recursion depth must be between 0 and 5")
		return
	}

	processDirectory(".", 0,  uint8(*recursionDepth), *hashType, *showAll)
}

func processDirectory(path string, level uint8, depth uint8, hashType string, showAll bool) {
	if level > depth {
		return
	}

	files, err := os.ReadDir(path)
	if err != nil {
		log.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		fullPath := fmt.Sprintf("%s/%s", path, file.Name())
		info, err := os.Stat(fullPath)
		if err != nil {
			log.Println("Error getting file info:", err)
			continue
		}

		stat := info.Sys().(*syscall.Stat_t)
		fileInfo := _fileinfo.NewFileInfo(
			info.Mode().Perm().String(),
			getUserName(stat.Uid),
			getGroupName(stat.Gid),
			info.Size(),
			info.ModTime(),
			fullPath,
			info.IsDir(),
			hashType,
		)

		fileInfo.Display(showAll, level)
		if info.IsDir() {
			processDirectory(fullPath, level+1, depth, hashType, showAll)
		}
	}
}



func getUserName(uid uint32) string {
	user, err := user.LookupId(fmt.Sprint(uid))
	if err != nil {
		return fmt.Sprint(uid)
	}
	return user.Username
}

func getGroupName(gid uint32) string {
	group, err := user.LookupGroupId(fmt.Sprint(gid))
	if err != nil {
		return fmt.Sprint(gid)
	}
	return group.Name
}
