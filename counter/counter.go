package counter

// Counter is basic counter
type Counter interface {
	Increment()
	Decrement()
	Count() int
}
