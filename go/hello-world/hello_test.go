package hello_world

import "testing"

func Test_hello(t *testing.T) {

	got := hello()
	t.Logf("Test_hello got = %s", got)
}
