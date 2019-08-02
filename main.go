package main

import "github.com/justenwalker/resume/cmd"

//go:generate go run main.go text --file ./README.md -t ./templates/markdown.tpl

func main() {
	cmd.Execute()
}
