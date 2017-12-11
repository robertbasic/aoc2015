package main

import (
	"testing"
)

func TestLights(t *testing.T) {
	lights := BuildLights(10)
	r := len(lights)
	e := 100

	if r != e {
		t.Errorf("Lights not built correctly, got total %d, expected %d", r, e)
	}

	ins := "turn on 0,0 through 9,9"

	lights = SetupLights(ins, lights)

	r = CountOnLights(lights)
	e = 100

	if r != e {
		t.Errorf("Got %d lights on, expected %d", r, e)
	}
}
