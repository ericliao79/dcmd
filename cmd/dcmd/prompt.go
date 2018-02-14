package main

import (
	"github.com/manifoldco/promptui"
)

func getSelect(label string, s map[string]string) promptui.Select {
	v := make([]string, 0, len(s))
	for _, value := range s {
		v = append(v, value)
	}
	prompt := promptui.Select{
		Label: label,
		Items: v,
	}

	return prompt
}
