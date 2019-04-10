package silk

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestCollection_Avg(t *testing.T) {
	a := []int{1,2,3,4,5,6,6,7,8,8}
	assert.Equal(t, Collect(a).Sum().ToInt64(), int64(50))
	assert.Equal(t, Collect(a).Avg().ToFloat64(), float64(5))

	b := []map[string]interface{}{
		{
			"foo": 10,
		},{
			"foo": 30,
		},{
			"foo": 20,
		},{
			"foo": 40,
		},
	}

	assert.Equal(t, Collect(b).Sum("foo").ToInt64(), int64(100))
}

func TestNumber_Add(t *testing.T) {
	a := NewNumberFromInt64(5)
	b := NewNumberFromInt64(5)
	c := NewNumberFromInt64(3)

	assert.Equal(t, a.Add(b).Reduce(c).ToInt64(), int64(7))
}

func TestCollection_Max(t *testing.T) {
	a := []int{1,2,3,4,5,6,6,7,8,8}
	assert.Equal(t, Collect(a).Max().ToInt64(), int64(8))

	b := []map[string]interface{}{
		{
			"foo": 10,
		},{
			"foo": 30,
		},{
			"foo": 20,
		},{
			"foo": 40,
		},
	}

	assert.Equal(t, Collect(b).Max("foo").ToInt64(), int64(40))
}

func TestCollection_Min(t *testing.T) {
	a := []int{1,2,3,4,5,6,6,7,8,8}
	assert.Equal(t, Collect(a).Min().ToInt64(), int64(1))

	b := []map[string]interface{}{
		{
			"foo": 10,
		},{
			"foo": 30,
		},{
			"foo": 20,
		},{
			"foo": 40,
		},
	}

	assert.Equal(t, Collect(b).Min("foo").ToInt64(), int64(10))
}