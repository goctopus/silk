package collection

import (
	"testing"
	"github.com/magiconair/properties/assert"
	"fmt"
)

func TestStringArrayCollection_Join(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(a).Join(""), "hello")
}

var numbers = []int{1, 2, 3, 4, 5, 6, 6, 7, 8, 8}

func TestNumberArrayCollection_Sum(t *testing.T) {
	assert.Equal(t, Collect(numbers).Sum().IntPart(), int64(50))

	var floatTest = []float64{143.66, -14.55}
	c := floatTest[0] + floatTest[1]
	fmt.Println(c)

	assert.Equal(t, Collect(floatTest).Sum().String(),"129.11")
}
