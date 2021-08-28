package main

import "fmt"

type hotkey struct {
	command command
}

func (h *hotkey) pressHotkey() {
	fmt.Println("Using hotkey combo to perform command.")
	h.command.execute()
}
