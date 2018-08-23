package counter

import (
	"syscall/js"
)

type domCounter struct {
	count   int
	countEl js.Value
}

// NewDOMCounter returns new DOMCounter
func NewDOMCounter(elementID string) Counter {
	c := &domCounter{}
	c.initDOM(elementID)
	return c
}

func (c *domCounter) initDOM(elementID string) {
	doc := js.Global().Get("document")
	el := doc.Call("getElementById", elementID)
	el.Call("insertAdjacentHTML", "beforeend", `<div id="count"></div><div><button id="inc">+</button><button id="dec">-</button></div>`)
	c.countEl = doc.Call("getElementById", "count")
	inc := doc.Call("getElementById", "inc")
	dec := doc.Call("getElementById", "dec")
	inc.Call("addEventListener", "click", js.NewCallback(func([]js.Value) { c.Increment() }))
	dec.Call("addEventListener", "click", js.NewCallback(func([]js.Value) { c.Decrement() }))
	c.update()
}

func (c *domCounter) update() {
	c.countEl.Set("textContent", c.count)
}

func (c *domCounter) Increment() {
	c.count++
	c.update()
}

func (c *domCounter) Decrement() {
	c.count--
	c.update()
}

func (c *domCounter) Count() int {
	return c.count
}
