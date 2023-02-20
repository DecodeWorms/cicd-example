package greeting

import (
	"testing"
)

func TestGreet(t *testing.T) {
	res := Greeting()
	if res != "Hello Mr Bola of Age 22 and at the Address Adewole Estate" {
		t.Error("test failed")
	}
}
