package main

import "github.com/lzdszdl/transfer.sh/cmd"

func main() {
	app := cmd.New()
	app.RunAndExitOnError()
}
