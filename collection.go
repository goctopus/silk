package silk

// interface{} must be one type of:
// Number, string, map[string]string, map[string]Number
type Collection []interface{}

func Collect(src interface{}) Collection {
	var c Collection

	switch src.(type) {
	case []map[string]interface{}:
		for _, v := range src.([]map[string]interface{}) {
			c = append(c, v)
		}
	case []int:
		for _, v := range src.([]int) {
			c = append(c, NewNumberFromInt(v))
		}
	case []int32:
		for _, v := range src.([]int32) {
			c = append(c, NewNumberFromInt32(v))
		}
	case []int64:
		for _, v := range src.([]int64) {
			c = append(c, NewNumberFromInt64(v))
		}
	case []float32:
		for _, v := range src.([]float32) {
			c = append(c, NewNumberFromFloat32(v))
		}
	case []float64:
		for _, v := range src.([]float64) {
			c = append(c, NewNumberFromFloat64(v))
		}
	default:
		panic("wrong type")
	}

	return c
}

func (c Collection) All() []interface{} {
	return []interface{}(c)
}

func (c Collection) Avg(key ...interface{}) Number {
	return c.Sum(key...).Divide(NewNumberFromInt(len(c)))
}

func (c Collection) Sum(key ...interface{}) Number {
	var sum = NewNumberFromInt(0)

	if len(key) == 0 {
		for i := 0; i < len(c); i++ {
			switch c[i].(type) {
			case Number:
				sum.Add(c[i].(Number))
			default:
				continue
			}
		}
	}

	return sum
}

func (c Collection) Min(key ...interface{}) interface{} {
	panic("implement it")
}

func (c Collection) Max(key ...interface{}) interface{} {
	panic("implement it")
}

func (c Collection) Mode(key interface{}) []interface{} {
	panic("implement it")
}

func (c Collection) Only(keys []interface{}) Collection {
	panic("implement it")
}

func (c Collection) Pluck(key interface{}) []interface{} {
	panic("implement it")
}

func (c Collection) Prepend(key interface{}, value interface{}) Collection {
	panic("implement it")
}

func (c Collection) Pull(key interface{}) Collection {
	panic("implement it")
}

func (c Collection) Put(key interface{}, value interface{}) Collection {
	panic("implement it")
}

func (c Collection) SortBy(key interface{}) Collection {
	panic("implement it")
}

func (c Collection) Spice(index int) Collection {
	panic("implement it")
}

func (c Collection) Take(num int) Collection {
	panic("implement it")
}

func (c Collection) ToJson() string {
	panic("implement it")
}

func (c Collection) Where(key interface{}, value interface{}) Collection {
	panic("implement it")
}

type Number struct {
	value *float64
}

func NewNumberFromInt(a int) Number {
	d := float64(a)
	return Number{
		value: &d,
	}
}

func NewNumberFromInt32(a int32) Number {
	d := float64(a)
	return Number{
		value: &d,
	}
}

func NewNumberFromInt64(a int64) Number {
	d := float64(a)
	return Number{
		value: &d,
	}
}

func NewNumberFromFloat32(a float32) Number {
	d := float64(a)
	return Number{
		value: &d,
	}
}

func NewNumberFromFloat64(a float64) Number {
	return Number{
		value: &a,
	}
}

func (n Number) Add(src Number) Number {
	*(n.value) += *(src.value)
	return n
}

func (n Number) Reduce(src Number) Number {
	*(n.value) -= *(src.value)
	return n
}

func (n Number) Plus(src Number) Number {
	*(n.value) *= *(src.value)
	return n
}

func (n Number) Divide(src Number) Number {
	*(n.value) /= *(src.value)
	return n
}

func (n Number) ToInt64() int64 {
	return int64(*(n.value))
}

func (n Number) ToFloat64() float64 {
	return *(n.value)
}