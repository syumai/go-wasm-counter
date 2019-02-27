package counter

import (
	"syscall/js"
)

type basicCounter struct {
	count   int
	countEl js.Value
}

// NewBasicCounter returns new DOMCounter
func NewBasicCounter(elementID string) Counter {
	c := &basicCounter{}
	c.initDOM(elementID)
	return c
}

func (c *basicCounter) initDOM(elementID string) {
	doc := js.Global().Get("document")
	el := doc.Call("getElementById", elementID)
	el.Call("insertAdjacentHTML", "beforeend", `<div id="count"></div><div><button id="inc">+</button><button id="dec">-</button></div>`)
	c.countEl = doc.Call("getElementById", "count")
	inc := doc.Call("getElementById", "inc")
	dec := doc.Call("getElementById", "dec")
	inc.Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} { c.Increment(); return nil }))
	dec.Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} { c.Decrement(); return nil }))
	c.update()
}

func (c *basicCounter) update() {
	c.countEl.Set("textContent", c.count)
}

func (c *basicCounter) Increment() {
	c.count++
	c.update()
}

func (c *basicCounter) Decrement() {
	c.count--
	c.update()
}

func (c *basicCounter) Count() int {
	return c.count
}
