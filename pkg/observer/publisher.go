package observer

// publisher interface
type Publisher interface {
	Attach(observer Observer)
	Unsubscribe(observer Observer)
	SetTitleAndData(title, data string)
	Notify()
}
