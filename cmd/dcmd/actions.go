package main

import (
	"bufio"
	"fmt"
	"github.com/ericliao79/dcmd"
	"github.com/fatih/color"
	cli "gopkg.in/urfave/cli.v1"
	"os"
)

//Create dcmd config
func initialize(c *cli.Context) error {
	var service []string
	var con []dcmd.Container
	//check Config dir
	if _, err := dcmd.IsEmpty(dcmd.MyAppConfig.StorePath); !os.IsNotExist(err) {
		err := os.RemoveAll(dcmd.MyAppConfig.StorePath)
		if err != nil {
			color.Red("%sFailed to remove existing empty store!", dcmd.MyAppConfig.CrossSymbol)
			return nil
		}
	}

	err := os.Mkdir(dcmd.MyAppConfig.StorePath, 0755)
	if err != nil {
		color.Red("%sFailed to initialize Config", dcmd.MyAppConfig.CrossSymbol)
		return nil
	}

	color.Blue("%sWhere is your docker-compose path?", dcmd.MyAppConfig.EditSymbol)
	fmt.Printf("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	path := scanner.Text()

	composes := dcmd.LoadComposes(path)
	prompt := getSelect("Do you want set compose?", composes)
	_, result, _ := prompt.Run()

	services := dcmd.LoadComposeYaml(path, result)
	services["Exit"] = "Exit."
	for {
		prompt = getSelect("What is you want add service? select Exit. goto Next", services)
		_, ser, _ := prompt.Run()
		delete(services, ser)
		if ser == "Exit." {
			break
		} else {
			service = append(service, ser)
		}
	}

	con = append(con, dcmd.Container{
		Name:    result,
		Service: service,
	})

	if _, err := dcmd.SetConfig(path, &con); err != nil {
		color.Red("%sFailed to initialize Config", dcmd.MyAppConfig.CrossSymbol)
		return nil
	}

	color.Green("%sCongratulations. Initialized done. ", dcmd.MyAppConfig.CheckSymbol)
	return nil
}

func up(c *cli.Context) error {
	composes := dcmd.LoadComposes()
	if c.NArg() > 0 {
		c := c.Args().Get(0)
		if _, ok := composes[c]; ok {
			dcmd.Start(c)
		}
	} else {
		prompt := getSelect("Please select Project.", composes)
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return nil
		}

		if len(result) > 0 {
			if _, ok := composes[result]; ok {
				dcmd.Start(result)
			} else {
				color.Red("%s Please select Project.", dcmd.MyAppConfig.CrossSymbol)
			}
		} else {
			color.Red("%s Please select Project.", dcmd.MyAppConfig.CrossSymbol)
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
