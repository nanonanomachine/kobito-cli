package main

import (
	. "github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func cmdWrite(rootCmd *Command) {
	rootCmd.AddCommand(&Command{
		Use:   "write",
		Short: "Write a new item",
		Long:  "Write a new item",
		Run: func(cmd *Command, args []string) {
			stdinInfo, err := os.Stdin.Stat()
			failOnError(err)

			if stdinInfo.Size() == 0 {
				return
			}

			file, err := ioutil.TempFile("/tmp", "kobito-cli-stdin-")
			failOnError(err)
			defer os.Remove(file.Name())

			_, err = io.Copy(file, os.Stdin)
			failOnError(err)

			_, err = exec.Command("open", "-a", "Kobito", file.Name()).Output()
			failOnError(err)

			os.Exit(0)
		}})
}
