package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const blogDir = "/Users/arturkondas/Desktop/git/youshy.github.io/_posts/"

const help = `This beauty helps you with writing posts. Or automating them. Whatever.

- post <title> - creates a post with <title>

- help - you're reading this nao

Happy posting!
`

func main() {
	action := os.Args[1]

	if len(os.Args) == 1 {
		fmt.Printf("Gimme title!\n")
		os.Exit(1)
	}

	switch action {
	case "post":
		title := os.Args[2]
		createPost(title)
	case "help":
		fmt.Printf(help)
		os.Exit(1)
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

	fmt.Printf("Is that post finished? Do you want to commit? (y/n)")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	switch input.Text() {
	case "n":
		fmt.Printf("Aight, finish it later!\n")
		os.Exit(1)
	case "y":
		fmt.Printf("Let's go with this post then!\n")
		add := exec.Command("git", "add", fileName)
		add.Dir = blogDir
		add.Run()

		commitName := "New Post: " + title
		commit := exec.Command("git", "commit", "-m", commitName)
		commit.Dir = blogDir
		commit.Run()

		push := exec.Command("git", "push")
		push.Dir = blogDir
		push.Run()
	default:
		fmt.Printf("I do not know what you want to do then... Bye!\n")
		os.Exit(1)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
