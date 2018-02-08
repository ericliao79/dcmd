package main

import (
	"github.com/fatih/color"
	"github.com/ericliao79/dcmd"
)

var (
	Version = dcmd.Version
	logo    = `

 V%s
`
)

func displayLogo() {
	color.Cyan(logo, Version)
}