package collection

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/shopspring/decimal"
	"testing"
)

func TestStringArrayCollection_Join(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(a).Join(""), "hello")
}

var (
	numbers = []int{1, 2, 3, 4, 5, 6, 6, 7, 8, 8}
	foo     = []map[string]interface{}{
		{
			"foo": 10,
		}, {
			"foo": 30,
		}, {
			"foo": 20,
		}, {
			"foo": 40,
		},
	}
)

func TestNumberArrayCollection_Sum(t *testing.T) {
	assert.Equal(t, Collect(numbers).Sum().IntPart(), int64(50))

	var floatTest = []float64{143.66, -14.55}
	c := floatTest[0] + floatTest[1]
	fmt.Println(c)

	assert.Equal(t, Collect(floatTest).Sum().String(), "129.11")
}

func TestCollection_Splice(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(a).Splice(1, 3, []string{"a"}), Collect([]string{"h", "a", "o"}))
	assert.Equal(t, Collect(numbers).Splice(0, 10,
		[]decimal.Decimal{NewDecimalFromInterface(3)}), Collect([]int{3}))
	assert.Equal(t, Collect(foo).Splice(1, 2, nil), Collect([]map[string]interface{}{
		{
			"foo": 10,
		}, {
			"foo": 40,
		}}))
}

func TestCollection_Take(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(foo[0]).Take(1), Collect(map[string]interface{}{"foo": 10}))
	assert.Equal(t, Collect(a).Take(-2), Collect([]string{"l", "o"}))
	assert.Equal(t, Collect(numbers).Take(4), Collect([]int{1, 2, 3, 4}))
	assert.Equal(t, Collect(foo).Take(2), Collect([]map[string]interface{}{
		{
			"foo": 10,
		}, {
			"foo": 30,
		}}))
}

func TestCollection_All(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(a).All(), []interface{}{"h", "e", "l", "l", "o"})
	assert.Equal(t, len(Collect(numbers).All()), 10)
	assert.Equal(t, Collect(foo).All()[1], map[string]interface{}{"foo": 30})
}

func TestCollection_Mode(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o", "w", "o", "l", "d"}
	foo2 := []map[string]interface{}{
		{
			"foo": 10,
		}, {
			"foo": 30,
		}, {
			"foo": 20,
		}, {
			"foo": 40,
		}, {
			"foo": 40,
		},
	}

	assert.Equal(t, Collect(numbers).Mode(""),
		[]interface{}{NewDecimalFromInterface(6), NewDecimalFromInterface(8)})
	assert.Equal(t, Collect(a).Mode(""), []interface{}{"l"})
	assert.Equal(t, Collect(foo2).Mode("foo"), []interface{}{40})
}

func TestCollection_Chunk(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}
	assert.Equal(t, Collect(foo).Chunk(2).value[0][0], map[string]interface{}{"foo": 10})
	assert.Equal(t, len(Collect(numbers).Chunk(3).value), 4)
	assert.Equal(t, Collect(a).Chunk(3).value[0][2], "l")
}
