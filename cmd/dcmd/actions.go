package main

import (
	"os"
	"fmt"
	"bufio"

	cli "gopkg.in/urfave/cli.v1"
	"github.com/ericliao79/dcmd"
	"github.com/fatih/color"
)

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

func list(c *cli.Context) error {
	composes := dcmd.LoadComposes()

	color.Green("Found %d Docker-compose(s)!", len(composes))
	fmt.Println("")
	for _, data := range composes {
		color.Green("%s" + data, ">  ")
	}
	return nil
}
