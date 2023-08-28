package main

import "design_pattern/3_9_state/state"

func main() {
	context := state.NewContext()
	context.Close()
	context.Close()
	context.Run()
	context.Open()
	context.Close()
	context.Stop()
	context.Open()
}
