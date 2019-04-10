package silk

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestCollection_Avg(t *testing.T) {
	a := []int{1,2,3,4,5,6,6,7,8,8}
	assert.Equal(t, Collect(a).Sum().ToInt64(), int64(50))
	assert.Equal(t, Collect(a).Avg().ToFloat64(), float64(5))
}

func TestNumber_Add(t *testing.T) {
	a := NewNumberFromInt64(5)
	b := NewNumberFromInt64(5)
	c := NewNumberFromInt64(3)

	assert.Equal(t, a.Add(b).Reduce(c).ToInt64(), int64(7))
}