package main

import (
	"os"
	"fmt"
	"bufio"

	cli "gopkg.in/urfave/cli.v1"
	"github.com/ericliao79/dcmd"
	"github.com/fatih/color"
	"github.com/c-bata/go-prompt"
)

//Create dcmd config
func initialize(c *cli.Context) error {
	//check Config dir
	if _, err := dcmd.IsEmpty(dcmd.StorePath); !os.IsNotExist(err) {
		err := os.Remove(dcmd.StorePath)

		if err != nil {
			color.Red("%sFailed to remove existing empty store!", dcmd.CrossSymbol)
			return nil
		}
	}

	err := os.Mkdir(dcmd.StorePath, 0755)
	if err != nil {
		color.Red("%sFailed to initialize Config", dcmd.CrossSymbol)
		return nil
	}

	color.Blue("%sWhere is your docker-compose path?", dcmd.EditSymbol)
	fmt.Printf("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	path := scanner.Text()
	if scanner.Err() != nil {
		// handle error.
	}
	if _, err := dcmd.SetConfig(path); err != nil {
		color.Red("%sFailed to initialize Config", dcmd.CrossSymbol)
		return nil
	}

	color.Green("%sCongratulations. Initialized done. ", dcmd.CheckSymbol)

	return nil
}

func up(c *cli.Context) error {
	if c.NArg() > 0 {
		c := c.Args().Get(0)
		composes := dcmd.LoadComposes()
		if _ , ok := composes[c]; ok {
			dcmd.Start(c)
		}
	} else {
		color.White("Please select Project.")
		t := prompt.Input("> ", completer)
		if len(t) > 0 {
		} else {
			color.Red("%s Please select Project.", dcmd.CrossSymbol)
		}
	}

	return nil
}

func list(c *cli.Context) error {
	composes := dcmd.LoadComposes()

	color.Green("Found %d Docker-compose(s)!", len(composes))
	color.White("")
	for _, data := range composes {
		color.Green("%s"+data, ">  ")
	}
	return nil
}

func stop(c *cli.Context) error {
	dcmd.Stop()

	return nil
}
