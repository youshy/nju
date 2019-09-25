package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const blogDir = "/Users/arturkondas/Desktop/git/youshy.github.io/_posts/"

func main() {
	action := os.Args[1]

	if len(os.Args) == 1 {
		fmt.Printf("Gimme title!\n")
		os.Exit(1)
	}

	title := os.Args[2]

	switch action {
	case "post":
		createPost(title)
	default:
		fmt.Printf("I do not understand what you need me to do.\n")
		os.Exit(1)
	}
}

func createPost(title string) {
	t := time.Now()
	year := strconv.Itoa(t.Year())
	monthInt := int(t.Month())
	dayInt := t.Day()
	var month, day string

	if monthInt < 10 {
		month = "0" + strconv.Itoa(monthInt)
	} else {
		month = strconv.Itoa(monthInt)
	}

	if dayInt < 10 {
		day = "0" + strconv.Itoa(dayInt)
	} else {
		day = strconv.Itoa(dayInt)
	}

	fileNameString := strings.ReplaceAll(title, " ", "-") + ".md"
	fileName := year + "-" + month + "-" + day + "-" + fileNameString

	header := "---\nlayout: post\ntitle: " + title + "\n---\n\n"
	f, err := os.Create(filepath.Join(blogDir, filepath.Base(fileName)))
	check(err)

	defer f.Close()

	b, err := f.WriteString(header)
	check(err)
	fmt.Printf("Wrote %d bytes\n", b)

	f.Sync()

	cmd := exec.Command("vim", filepath.Join(blogDir, filepath.Base(fileName)))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
