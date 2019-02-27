package counter

import (
	"syscall/js"
)

type vueCounter struct {
	vm js.Value
}

// NewVueCounter returns new VueCounter
func NewVueCounter(elementID string) Counter {
	c := &vueCounter{}
	c.initVM(elementID)
	return c
}

func (c *vueCounter) initVM(elementID string) {
	vue := js.Global().Get("Vue")
	c.vm = vue.New(map[string]interface{}{
		"el":       "#" + elementID,
		"template": `<div><div>{{count}}</div><div><button @click="increment">+</button><button @click="decrement">-</button></div></div>`,
		"data": map[string]interface{}{
			"count": 0,
		},
		"methods": map[string]interface{}{
			"increment": js.FuncOf(func(js.Value, []js.Value) interface{} { c.Increment(); return nil }),
			"decrement": js.FuncOf(func(js.Value, []js.Value) interface{} { c.Decrement(); return nil }),
		},
	})
}

func (c *vueCounter) Increment() {
	c.vm.Set("count", c.Count()+1)
}

func (c *vueCounter) Decrement() {
	c.vm.Set("count", c.Count()-1)
}

func (c *vueCounter) Count() int {
	return c.vm.Get("count").Int()
}
