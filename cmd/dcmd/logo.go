package main

import (
	"github.com/fatih/color"
	"github.com/ericliao79/dcmd"
)

var (

	logo    = `

 V%s
`
)

func displayLogo() {
	color.Cyan(logo, dcmd.Version)
}