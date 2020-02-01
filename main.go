package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("dotmac starting")

	fmt.Print("Install binaries? ")
	if getYesNo() {
		run("brew", "install", "macvim", "tmux", "golangci/tap/golangci-lint")
	}

	fmt.Println("Make symbolic links")
	if getYesNo() {
		run("ln", "-s", "~/.dotmac/tmux/tmux.conf", "~/.tmux.conf")
		run("ln", "-s", "~/.dotmac/zsh/zshrc", "~/.zshrc")
		run("ln", "-s", "~/.dotmac/vimrc", "~/.vimrc")
	}
}

func getYesNo() bool {
	fmt.Print("Y/n: ")
	switch strings.ToLower(getInput()) {
	case "yes", "", "y":
		return true
	default:
		return false
	}
}

func getInput() string {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}

func run(cmd string, args ...string) {
	c := exec.Command(cmd, args...)
	out, err := c.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
