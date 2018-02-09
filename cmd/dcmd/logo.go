package main

import (
	"github.com/ericliao79/dcmd"
	"github.com/fatih/color"
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
