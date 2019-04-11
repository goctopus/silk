package silk

import (
	"github.com/shopspring/decimal"
)

// interface{} must be one type of:
// []decimal.Decimal, []string, []map[string]interface{}, map[string]interface{}
type Collection struct {
	value  interface{}
	length int
}

func Collect(src interface{}) Collection {
	var c Collection

	switch src.(type) {
	case []map[string]interface{}:
		c.value = src
		c.length = len(src.([]map[string]interface{}))
	case []string:
		c.value = src
		c.length = len(src.([]string))
	case map[string]interface{}:
		c.value = src
		c.length = len(src.(map[string]interface{}))
	case []int:
		var d = make([]decimal.Decimal, 0)
		for _, v := range src.([]int) {
			d = append(d, decimal.New(int64(v), 0))
		}
		c.value = d
		c.length = len(src.([]int))
	case []int8:
		var d = make([]decimal.Decimal, 0)
		for _, v := range src.([]int8) {
			d = append(d, decimal.New(int64(v), 0))
		}
		c.value = d
		c.length = len(src.([]int8))
	case []int16:
		var d = make([]decimal.Decimal, 0)
		for _, v := range src.([]int16) {
			d = append(d, decimal.New(int64(v), 0))
		}
		c.value = d
		c.length = len(src.([]int16))
	case []int32:
		var d = make([]decimal.Decimal, 0)
		for _, v := range src.([]int32) {
			d = append(d, decimal.New(int64(v), 0))
		}
		c.value = d
		c.length = len(src.([]int32))
	case []int64:
		var d = make([]decimal.Decimal, 0)
		for _, v := range src.([]int64) {
			d = append(d, decimal.New(v, 0))
		}
		c.value = d
		c.length = len(src.([]int64))
	case []float32:
		var f = make([]decimal.Decimal, 0)
		for _, v := range src.([]float32) {
			f = append(f, decimal.NewFromFloat32(v))
		}
		c.value = f
		c.length = len(src.([]float32))
	case []float64:
		var f = make([]decimal.Decimal, 0)
		for _, v := range src.([]float64) {
			f = append(f, decimal.NewFromFloat(v))
		}
		c.value = f
		c.length = len(src.([]float64))
	default:
		panic("wrong type")
	}

	return c
}

func (c Collection) All() []interface{} {
	panic("implement it")
}

func (c Collection) Avg(key ...string) decimal.Decimal {
	return c.Sum(key...).Div(decimal.New(int64(c.length), 0))
}

func (c Collection) Sum(key ...string) decimal.Decimal {
	var sum = decimal.New(0, 0)

	if len(key) == 0 {
		if n, ok := c.value.([]decimal.Decimal); ok {
			for i := 0; i < len(n); i++ {
				sum = sum.Add(n[i])
			}
		}
	} else {
		if n, ok := c.value.([]map[string]interface{}); ok {
			for i := 0; i < len(n); i++ {
				sum = sum.Add(NewDecimalFromInterface(n[i][key[0]]))
			}
		}
	}

	return sum
}

func (c Collection) Min(key ...string) decimal.Decimal {

	var smallest = decimal.New(0, 0)

	if len(key) == 0 {
		if n, ok := c.value.([]decimal.Decimal); ok {
			for i := 0; i < len(n); i++ {
				if i == 0 {
					smallest = n[i]
					continue
				}
				if smallest.GreaterThan(n[i]) {
					smallest = n[i]
				}
			}
		}
	} else {

		var number decimal.Decimal

		if n, ok := c.value.([]map[string]interface{}); ok {
			for i := 0; i < len(n); i++ {
				number = NewDecimalFromInterface(n[i][key[0]])
				if i == 0 {
					smallest = number
					continue
				}
				if smallest.GreaterThan(number) {
					smallest = number
				}
			}
		}
	}

	return smallest
}

func (c Collection) Max(key ...string) decimal.Decimal {

	var biggest = decimal.New(0, 0)

	if len(key) == 0 {
		if n, ok := c.value.([]decimal.Decimal); ok {
			for i := 0; i < len(n); i++ {
				if i == 0 {
					biggest = n[i]
					continue
				}
				if biggest.LessThan(n[i]) {
					biggest = n[i]
				}
			}
		}
	} else {

		var number decimal.Decimal

		if n, ok := c.value.([]map[string]interface{}); ok {
			for i := 0; i < len(n); i++ {
				number = NewDecimalFromInterface(n[i][key[0]])
				if i == 0 {
					biggest = number
					continue
				}
				if biggest.LessThan(number) {
					biggest = number
				}
			}
		}
	}

	return biggest
}

func (c Collection) Join(delimiter string) string {
	s := ""

	if n, ok := c.value.([]string); ok {
		for i := 0; i < len(n); i++ {
			if i != len(n)-1 {
				s += n[i] + delimiter
			} else {
				s += n[i]
			}
		}
	} else if n, ok := c.value.([]decimal.Decimal); ok {
		for i := 0; i < len(n); i++ {
			if i != len(n)-1 {
				s += n[i].String() + delimiter
			} else {
				s += n[i].String()
			}
		}
	}

	return s
}

func (c Collection) Combine(value []interface{}) Collection {
	var (
		m      = make(map[string]interface{}, 0)
		length = c.length
	)

	if length > len(value) {
		length = len(value)
	}

	if n, ok := c.value.([]string); ok {
		for i := 0; i < length; i++ {
			m[n[i]] = value[i]
		}
	}

	return Collect(m)
}

func (c Collection) Count() int {
	return c.length
}

func (c Collection) Mode(key ...string) []interface{} {
	panic("implement it")
}

func (c Collection) Only(keys []string) Collection {
	panic("implement it")
}

func (c Collection) Pluck(key string) []interface{} {
	panic("implement it")
}

func (c Collection) Prepend(key string, value interface{}) Collection {
	panic("implement it")
}

func (c Collection) Pull(key interface{}) Collection {
	panic("implement it")
}

func (c Collection) Put(key string, value interface{}) Collection {
	panic("implement it")
}

func (c Collection) SortBy(key string) Collection {
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

func (c Collection) ToMap() map[string]interface{} {
	if m, ok := c.value.(map[string]interface{}); ok {
		return m
	} else {
		return map[string]interface{}{}
	}
}

func (c Collection) Where(key string, value interface{}) Collection {
	panic("implement it")
}

func NewDecimalFromInterface(a interface{}) decimal.Decimal {
	var d decimal.Decimal

	switch a.(type) {
	case int:
		d = decimal.New(int64(a.(int)), 0)
	case int8:
		d = decimal.New(int64(a.(int8)), 0)
	case int16:
		d = decimal.New(int64(a.(int16)), 0)
	case int32:
		d = decimal.New(int64(a.(int32)), 0)
	case int64:
		d = decimal.New(a.(int64), 0)
	case float32:
		d = decimal.NewFromFloat32(a.(float32))
	case float64:
		d = decimal.NewFromFloat(a.(float64))
	}

	return d
}
