package silk

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

var foo = []map[string]interface{}{
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

var numbers = []int{1, 2, 3, 4, 5, 6, 6, 7, 8, 8}

func TestCollection_Avg(t *testing.T) {
	assert.Equal(t, Collect(numbers).Sum().IntPart(), int64(50))
	f, _ := Collect(numbers).Avg().Float64()
	assert.Equal(t, f, float64(5))

	assert.Equal(t, Collect(foo).Sum("foo").IntPart(), int64(100))
}

func TestCollection_Max(t *testing.T) {
	assert.Equal(t, Collect(numbers).Max().IntPart(), int64(8))
	assert.Equal(t, Collect(foo).Max("foo").IntPart(), int64(40))
}

func TestCollection_Min(t *testing.T) {
	assert.Equal(t, Collect(numbers).Min().IntPart(), int64(1))

	assert.Equal(t, Collect(foo).Min("foo").IntPart(), int64(10))
}

func TestCollection_Join(t *testing.T) {
	a := []string{"h", "e", "l", "l", "o"}

	assert.Equal(t, Collect(a).Join(""), "hello")
}

func TestCollection_Combine(t *testing.T) {
	a := Collect([]string{"name", "age"})
	b := []interface{}{"John", 18}

	assert.Equal(t, a.Combine(b).ToMap()["name"], "John")
}

func TestCollection_Pluck(t *testing.T) {
	assert.Equal(t, Collect(foo).Pluck("foo").ToNumberArray()[0].IntPart(), int64(10))
}

func TestCollection_Only(t *testing.T) {
	a := map[string]interface{}{
		"product_id": 1,
		"name":       "Desk",
		"price":      100,
		"discount":   false,
	}

	assert.Equal(t, Collect(a).Only([]string{"name", "price"}).ToMap(), map[string]interface{}{
		"name":  "Desk",
		"price": 100,
	})
}

func TestCollection_Prepend(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	assert.Equal(t, Collect(a).Prepend(0).ToNumberArray()[0].IntPart(), int64(0))
}

func TestCollection_Pull(t *testing.T) {
	a := map[string]interface{}{
		"product_id": 1,
		"name":       "Desk",
		"price":      100,
		"discount":   false,
	}

	assert.Equal(t, Collect(a).Pull("name").length, 3)
}

func TestCollection_Put(t *testing.T) {
	a := map[string]interface{}{
		"product_id": 1,
		"name":       "Desk",
		"price":      100,
		"discount":   false,
	}

	assert.Equal(t, Collect(a).Put("name1", 2121).length, 5)
}

func TestCollection_SortBy(t *testing.T) {

	m := make(map[int]map[string]interface{})
	c := map[string]interface{}{
		"product_id": 122,
		"name":       "Desk",
		"price":      100,
		"discount":   false,
	}

	m[0] = c
	d := map[string]interface{}{
		"product_id": 1,
		"name":       "Desk",
		"price":      100,
		"discount":   false,
	}
	m[1] = d
	fmt.Println(m)

	assert.Equal(t, Collect(m).SortBy("product_id").length, 5)
}
