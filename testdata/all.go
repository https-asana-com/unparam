package foo

import "net/http"

type FooType int

func AllUsed(a, b FooType) FooType {
	return a + b
}

func OneUnused(a, b FooType) FooType {
	return a
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

type FooIface interface {
	foo(w http.ResponseWriter, code FooType) error
}

func FooImpl(w http.ResponseWriter, code FooType) error {
	w.Write([]byte("hi"))
	return nil
}

func (f FooType) AllUsed(a, b FooType) FooType {
	return a + b
}

func DummyImpl(f FooType) {}

func PanicImpl(f FooType) { panic("dummy") }