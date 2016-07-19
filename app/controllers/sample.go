package controllers

import "github.com/revel/revel"

type Sample struct {
	*revel.Controller
}

func (c Sample) Index() revel.Result {
	message := "Hello World!"
	return c.Render(message)
}

func (c Sample) Hello(myName string) revel.Result {
    c.Validation.Required(myName).Message("Your Name is required.")
    c.Validation.MinSize(myName, 3).Message("Your Name is not long enough.")

    if (c.Validation.HasErrors()) {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(Sample.Index)
    }
    return c.Render(myName)
}
