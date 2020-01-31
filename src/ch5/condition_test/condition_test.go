package condition_test

import "testing"

func TestWhileLoop(t *testing.T) {
	n:=0
	for n < 5 {
		t.Log(n)
		n++
	}
 }

func TestIfMultiSec(t *testing.T)  {
	if a:=1 == 1;a {
		t.Log("1==1")
	}
}
