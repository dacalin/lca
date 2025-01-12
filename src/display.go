package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

// DisplayStyle holds flags or settings about how to display the info.
type Display struct {
	showAll        bool
	datetimeFormat string
}

var colorMap = map[string]string{
	"black":   "0",
	"red":     "1",
	"green":   "2",
	"yellow":  "3",
	"blue":    "4",
	"magenta": "5",
	"cyan":    "6",
	"white":   "7",
}

func NewDisplay(showAll bool, datetimeFormat string) Display {
	return Display{showAll, datetimeFormat}
}

// Prints a FileInfo according to the style settings.
func (ds Display) Print(f FileInfo, level uint8) {
	permissionsColor := "red"
	ownerColor := "cyan"
	groupColor := "cyan"
	sizeColor := "magenta"
	timeColor := "white"
	fileColor := "green"
	hashColor := "yellow"
	indentColor := "black"

	if f.IsDir() {
		permissionsColor = "red"
		ownerColor = "cyan"
		groupColor = "cyan"
		sizeColor = "magenta"
		timeColor = "white"
		fileColor = "white"
		hashColor = "white"
		indentColor = "white"
	}

	dir, fileName := filepath.Split(f.Name())
	dir = strings.TrimPrefix(dir, "./")

	permissionsDataFormat := col(permissionsColor) + col(ownerColor) + col(groupColor)
	defaultDataFormat := colWide(sizeColor, 8) + col(timeColor) + col(fileColor) +  col(hashColor)

	indent := strings.Repeat("    ", int(level))

	// Use ANSI colors
	if ds.showAll {
		permissions := f.Permissions().str + "("+f.Permissions().octal + ")"
		// show all
		fmt.Printf(
			col(indentColor)+permissionsDataFormat+defaultDataFormat+"\n",
			indent, permissions, f.Owner(), f.Group(), ds.formatSize(f.Size()),
			f.ModTime().Format(ds.datetimeFormat), dir+fileName, f.Hash(),
		)
	} else {
		// hide permissions columns
		fmt.Printf(
			col(indentColor)+defaultDataFormat+"\n", indent,
			ds.formatSize(f.Size()), f.ModTime().Format(ds.datetimeFormat), dir+fileName, f.Hash(),
		)
	}
}

func (ds Display) formatSize(size int64) string {
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
	)

	switch {
	case size >= TB:
		return fmt.Sprintf("%.2f TB", float64(size)/TB)
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/GB)
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/MB)
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/KB)
	default:
		return fmt.Sprintf("%d B", size)
	}
}

// col and colWide are small helpers for coloring strings.
func col(color string) string {
	colorNumber, ok := colorMap[color]
	if !ok {
		colorNumber = "7"
	}

	return "\033[3" + colorNumber + "m%s\033[0m" + "  "
}

func colWide(color string, wide uint8) string {
	colorNumber, ok := colorMap[color]
	if !ok {
		colorNumber = "7"
	}
	wideStr := strconv.Itoa(int(wide))

	return "\033[3" + colorNumber + "m%" + wideStr + "s\033[0m" + "  "
}
