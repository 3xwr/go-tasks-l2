package main

type saveCommand struct {
	invoker invoker
}

func (c *saveCommand) execute() {
	c.invoker.save()
}
