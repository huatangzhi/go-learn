package type_test

import "testing"
type myInt int64
func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a)
	t.Log(a, b)
	var c myInt
	c = myInt(b)
	t.Log(b,c)
}


func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s=="" {
		t.Log("Empty")
	}
}