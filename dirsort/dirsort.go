package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var extMap = map[string]string{
	".mp3":  "Music",
	".flac": "Music",
	".wav":  "Music",
	".jpeg": "Images",
	".jpg":  "Images",
	".png":  "Images",
	".svg":  "Images",
	".zip":  "Archives",
	".tar":  "Archives",
	".7z":   "Archives",
	".gz":   "Archives",
	".mp4":  "Videos",
	".mkv":  "Videos",
	".avi":  "Videos",
	".pdf":  "Docs",
	".ppt":  "Docs",
}

func sort(directory string) {
	f, err := os.Open(directory)
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.IsDir() != true {

			flag := false
			ext := filepath.Ext(file.Name())
			var key string

			for item := range extMap {
				if item == ext {
					flag = true
					key = item
				}
			}
			if flag == false {
				err := os.MkdirAll(filepath.Join(directory, "Misc"), os.ModePerm)
				if err == nil {
					e := os.Rename(filepath.Join(directory, file.Name()), filepath.Join(directory, "Misc", file.Name()))
					if e != nil {
						fmt.Printf("Failed with error: %v\n", e)
						os.Exit(1)
					}
				}
			} else {
				err := os.MkdirAll(filepath.Join(directory, extMap[key]), os.ModePerm)
				if err == nil {
					e := os.Rename(filepath.Join(directory, file.Name()), filepath.Join(directory, extMap[key], file.Name()))
					if e != nil {
						fmt.Printf("Failed with error: %v\n", e)
						os.Exit(1)
					}
				}
			}
		}
	}
}

func main() {
	var dir string
	args := os.Args

	if len(args) <= 1 {
		fmt.Printf("Please provide a valid path\nEg: /home/user/Downloads")
		os.Exit(1)

	} else {
		dir = args[1]
		sort(dir)
	}
}
