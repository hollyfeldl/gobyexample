// a go testing example
// extends go by example

package plus

import "testing"

func TestPlus(t *testing.T) {

	a := 4
	b := 5
	if Plus(a, b) != 9 {
		t.Error("Expected 9 when added ", a , "and", b)
	}

}

func TestPlusPlus(t *testing.T) {

	a := 4
	b := 5
	c := 6
	if PlusPlus(a, b, c) != 15 {
		t.Error("Expected 15 when added ", a , ",", b, "and", c)
	}

}
