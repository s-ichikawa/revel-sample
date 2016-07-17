package controllers

import "github.com/revel/revel"

type Sample struct {
	*revel.Controller
}

func (c Sample) Index() revel.Result {
	message := "Hello World!"
	return c.Render(message)
}
