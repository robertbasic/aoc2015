package main

import "testing"

func TestFindmd5(t *testing.T) {
	i := "abcdef"

	o := Findmd5(i)
	e := 609043

	if o != e {
		t.Errorf("Got %d, expected %d", o, e)
	}

	i = "pqrstuv"

	o = Findmd5(i)
	e = 1048970

	if o != e {
		t.Errorf("Got %d, expected %d", o, e)
	}
}
