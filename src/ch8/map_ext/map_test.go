package map_ext

import "testing"

func TestMapWithFunction(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2))
	t.Log(m[2](2))
	t.Log(m[3](2))
}

func TestMapForSet(t *testing.T)  {
	mySet := map[int] bool{}
	mySet[1] = true
	n:=1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[1] = true
	t.Log(len(mySet))
	delete(mySet, 1)

	if mySet[1] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
