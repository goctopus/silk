package silk

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestCollection_Avg(t *testing.T) {
	a := []int{1,2,3,4,5,6,6,7,8,8}
	assert.Equal(t, Collect(a).Sum().IntPart(), int64(50))
	f, _ := Collect(a).Avg().Float64()
	assert.Equal(t, f, float64(5))

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

	assert.Equal(t, Collect(b).Sum("foo").IntPart(), int64(100))
}

func TestCollection_Max(t *testing.T) {
	a := []int{1,2,3,4,5,6,6,7,8,8}
	assert.Equal(t, Collect(a).Max().IntPart(), int64(8))

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

	assert.Equal(t, Collect(b).Max("foo").IntPart(), int64(40))
}

func TestCollection_Min(t *testing.T) {
	a := []int{1,2,3,4,5,6,6,7,8,8}
	assert.Equal(t, Collect(a).Min().IntPart(), int64(1))

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

	assert.Equal(t, Collect(b).Min("foo").IntPart(), int64(10))
}