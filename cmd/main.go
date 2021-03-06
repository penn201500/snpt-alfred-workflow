package main

import (
	"fmt"
	"os"
	"os/exec"

	copyCmd "github.com/mike182uk/snpt-alfred-workflow/cmd/command/copy"
	searchCmd "github.com/mike182uk/snpt-alfred-workflow/cmd/command/search"
	titleCmd "github.com/mike182uk/snpt-alfred-workflow/cmd/command/title"
	alfredHelper "github.com/mike182uk/snpt-alfred-workflow/cmd/helper/alfred"
)

const snptBin = "snpt"
const fzfBin = "fzf"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Print("Usage: snpt-alfred-workflow <search|cp|title>")

		os.Exit(1)
	}

	checkDeps(snptBin, fzfBin)

	switch args[0] {
	case "search":
		searchCmd.Run(args, snptBin, fzfBin)
	case "cp":
		copyCmd.Run(args, snptBin)
	case "title":
		titleCmd.Run(args)
	}
}

func checkDeps(snptBin string, fzfBin string) {
	whichSnptCmd := exec.Command("which", snptBin) // #nosec
	whichFzfCmd := exec.Command("which", fzfBin)   // #nosec

	_, err := whichSnptCmd.Output()

	if err != nil {
		fmt.Print(alfredHelper.GenerateErrorItemList("snpt was not found", "Make sure you have installed snpt to use this workflow."))

		os.Exit(1)
	}

	_, err = whichFzfCmd.Output()

	if err != nil {
		fmt.Print(alfredHelper.GenerateErrorItemList("fzf was not found", "Make sure you have installed fzf to use this workflow."))

		os.Exit(1)
	}
}
