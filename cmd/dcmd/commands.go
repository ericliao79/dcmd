package main

import (
	cli "gopkg.in/urfave/cli.v1"
)

func initCommands() []cli.Command {
	return []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i", "initialize"},
			Usage:   "init config",
			Action:  initialize,
		},
		{
			Name:    "up",
			Usage:   "up your containers.",
			Action:  up,
		},
		{
			Name:    "ls",
			Usage:   "shows a list of your services",
			Action:  list,
		},
		{
			Name:    "stop",
			Usage:   "stop your all containers.",
			Action:  stop,
		},
	}
}