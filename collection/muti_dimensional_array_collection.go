package collection

type MultiDimensionalArrayCollection struct {
	value [][]interface{}
	BaseCollection
}

func (c MultiDimensionalArrayCollection) Collapse() Collection {
	if len(c.value[0]) == 0 {
		return Collect([]interface{}{})
	}

	// length = (len(outer_slice)-1) * len(inner_slice) + len(last_inner_slice)
	outer := len(c.value) - 1
	inner := len(c.value[0])
	last := len(c.value[len(c.value)-1])
	length := outer*inner + last
	d := make([]interface{}, length)

	// range outer_slice, and inner_slice except last one
	for i := 0; i < outer; i++ {
		for j := 0; j < inner; j++ {
			d[i*inner+j] = c.value[i][j]
		}
	}

	// range last_inner slice
	if last != 0 {
		for i := 0; i < last; i++ {
			d[outer*inner+i] = c.value[len(c.value)-1][i]
		}
	}

	return Collect(d)
}

func (c MultiDimensionalArrayCollection) Concat(value interface{}) Collection {
	return MultiDimensionalArrayCollection{
		value:          append(c.value, value.([][]interface{})...),
		BaseCollection: BaseCollection{length: c.length + len(value.([][]interface{}))},
	}
}
