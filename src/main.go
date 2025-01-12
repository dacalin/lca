package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
	"syscall"
)

func main() {
	var hashType string
	flag.StringVar(&hashType, "hash", "none", "hash type: md5, sha1, or none")
	flag.StringVar(&hashType, "H", "none", "shorthand for --hash")
	showAll 		:= flag.Bool("p", false, "show permissions and owners")
	recursionDepth 	:= flag.Uint("r", 0, "recursion depth")	  
	flag.Parse()

	// Validate hash type
	switch hashType {
	case "md5", "sha1", "sha256", "none":
		// valid
	default:
		log.Fatalf("Invalid hash type: %q (must be md5, sha1, sha256, or none)", hashType)
	}
	
	processDirectory(".", 0,  uint8(*recursionDepth), hashType, *showAll)
}

func processDirectory(path string, level uint8, depth uint8, hashType string, showAll bool) {
	if level > depth {
		return
	}

	display 	:= NewDisplay(showAll, "02 Jan 06 15:04")

	files, err 	:= os.ReadDir(path)
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
		
		fileInfo := NewFileInfo(
			getPermissions(info),
			getUserName(stat.Uid),
			getGroupName(stat.Gid),
			info.Size(),
			info.ModTime(),
			fullPath,
			info.IsDir(),
			hashType,
		)

		display.Print(fileInfo, level)

		if info.IsDir() {
			processDirectory(fullPath, level+1, depth, hashType, showAll)
		}
	}
}

func getPermissions(info os.FileInfo ) string {
	if info.IsDir() {
		perm := info.Mode().Perm().String()
		return "d"+strings.TrimPrefix(perm, "-")
	}
	return info.Mode().Perm().String()
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

