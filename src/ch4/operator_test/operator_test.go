package operator_test

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

func TestCompareArray(t *testing.T) {
	a := [...] int{1, 2, 3, 4}
	b := [...] int{1, 2, 4, 5}
	//c:=[...] int{1,2,3,4,5}
	d := [...] int{1, 2, 4, 3}

	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	a := 7
	a = a & ^ Executable
	t.Log(a)
}
