package collection

import (
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
	//c := floatTest[0] + floatTest[1]
	//fmt.Println(c)

	assert.Equal(t, Collect(floatTest).Sum().String(), "129.11")
}

func TestCollection_Splice(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	c := Collect(a)
	assert.Equal(t, c.Splice(1, 3).ToStringArray(), []string{"e", "l", "l"})
	assert.Equal(t, c.Splice(1).ToStringArray(), []string{"e", "l", "l", "o"})

	assert.Equal(t, Collect(numbers).Splice(2, 1).ToNumberArray(),
		[]decimal.Decimal{nd(3)})

	assert.Equal(t, Collect(foo).Splice(1, 2).ToMapArray(), []map[string]interface{}{
		{
			"foo": 30,
		}, {
			"foo": 20,
		},
	})
}

func TestCollection_Take(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(a).Take(-2).ToStringArray(), []string{"l", "o"})

	assert.Equal(t, Collect(numbers).Take(4).ToNumberArray(),
		[]decimal.Decimal{nd(1), nd(2), nd(3), nd(4)})

	assert.Equal(t, Collect(foo).Take(2).ToMapArray(), []map[string]interface{}{
		{
			"foo": 10,
		}, {
			"foo": 30,
		},
	})
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

	m := Collect(numbers).Mode()
	assert.Equal(t, m[0].(decimal.Decimal).IntPart() == int64(8) ||
		m[0].(decimal.Decimal).IntPart() == int64(6), true)
	assert.Equal(t, m[1].(decimal.Decimal).IntPart() == int64(8) ||
		m[1].(decimal.Decimal).IntPart() == int64(6), true)

	assert.Equal(t, Collect(a).Mode(), []interface{}{"l"})
	assert.Equal(t, Collect(foo2).Mode("foo"), []interface{}{40})
}

func TestCollection_Chunk(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(foo).Chunk(2).value[0][0], map[string]interface{}{"foo": 10})
	assert.Equal(t, len(Collect(numbers).Chunk(3).value), 4)
	assert.Equal(t, Collect(a).Chunk(3).value[0][2], "l")
}

func TestCollection_Collapse(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(foo).Chunk(2).Collapse(), Collect(foo))
	assert.Equal(t, Collect(a).Chunk(3).Collapse(), Collect(a))
	assert.Equal(t, Collect(numbers).Chunk(3).Collapse(), Collect(numbers))
}

func TestBaseCollection_Concat(t *testing.T) {
	test_numbers := []int{1, 2, 3, 4, 5, 6, 6, 7, 8, 8, 9}
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, len(Collect(foo).Concat(
		[]map[string]interface{}{{"foo": 100}}).ToMapArray()), 5)
	assert.Equal(t, Collect(numbers).Concat(
		[]decimal.Decimal{newDecimalFromInterface(9)}), Collect(test_numbers))
	assert.Equal(t, Collect(a).Concat([]string{"world"}).All()[5], "world")
	assert.Equal(t, Collect(numbers).Chunk(2).Concat(
		[][]interface{}{}).Collapse(), Collect(numbers))
}
