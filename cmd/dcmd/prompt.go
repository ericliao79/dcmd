package main

import (
	"github.com/c-bata/go-prompt"
	"github.com/ericliao79/dcmd"
)

func completer(d prompt.Document) []prompt.Suggest {
	composes := dcmd.LoadComposes()
	var s []prompt.Suggest

	for _, data := range composes  {
		s = append(s, prompt.Suggest{Text: data, Description: ""})
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}