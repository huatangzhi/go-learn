package const_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T)  {
	t.Log(Monday,Tuesday, Wednesday)
}

func TestConstantTry1(t *testing.T)  {
	a:=1 // 001
	t.Log(a& Readable)
	t.Log(a& Writable)
	t.Log(a& Executable)
}
