package main

import (
	"os"

	cli "gopkg.in/urfave/cli.v1"
	"github.com/ericliao79/dcmd"
)

func main()  {
	parseArgs()
	app := cli.NewApp()
	app.Name = dcmd.Name
	app.Usage = dcmd.Usage
	app.Commands = initCommands()
	app.Version = dcmd.Version

	app.Run(os.Args)
}

// ParseArgs parses input arguments and displays the program logo
func parseArgs() {
	if len(os.Args) == 1 {
		displayLogo()
	} else if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "h" || os.Args[1] == "help" {
			displayLogo()
		}
	}
}