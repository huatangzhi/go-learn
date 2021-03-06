package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 [] int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))
	s1 := [] int{1, 2, 3, 4, 5}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 6, 8)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[5])
	s2 = append(s2, 99)
	t.Log(s2[0], s2[1], s2[6])
}

func TestSliceGrowing(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceSharedMemory(t *testing.T)  {
	year := []string{
		"Jan", "Feb","Mar", "Apr", "May", "June", "July", "Aug","Sep","Nov","Dec"}

	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer:=year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unknow"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceComparing(t *testing.T)  {
	a := [] int {1,2,3,4}
	b := [] int {1,2,3}
	t.Log(a == b)
}
