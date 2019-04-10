package silk

import "github.com/shopspring/decimal"

// interface{} must be one type of:
// decimal.Decimal, string, map[string]string, map[string]decimal.Decimal
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
			c = append(c, decimal.New(int64(v), 0))
		}
	case []int8:
		for _, v := range src.([]int8) {
			c = append(c, decimal.New(int64(v), 0))
		}
	case []int16:
		for _, v := range src.([]int16) {
			c = append(c, decimal.New(int64(v), 0))
		}
	case []int32:
		for _, v := range src.([]int32) {
			c = append(c, decimal.New(int64(v), 0))
		}
	case []int64:
		for _, v := range src.([]int64) {
			c = append(c, decimal.New(v, 0))
		}
	case []float32:
		for _, v := range src.([]float32) {
			c = append(c, decimal.NewFromFloat32(v))
		}
	case []float64:
		for _, v := range src.([]float64) {
			c = append(c, decimal.NewFromFloat(v))
		}
	default:
		panic("wrong type")
	}

	return c
}

func (c Collection) All() []interface{} {
	return []interface{}(c)
}

func (c Collection) Avg(key ...string) decimal.Decimal {
	return c.Sum(key...).Div(decimal.New(int64(len(c)), 0))
}

func (c Collection) Sum(key ...string) decimal.Decimal {
	var sum = decimal.New(0, 0)

	if len(key) == 0 {
		for i := 0; i < len(c); i++ {
			switch c[i].(type) {
			case decimal.Decimal:
				sum = sum.Add(c[i].(decimal.Decimal))
			default:
				continue
			}
		}
	} else {
		for i := 0; i < len(c); i++ {
			switch c[i].(type) {
			case map[string]interface{}:
				sum = sum.Add(NewDecimalFromInterface(c[i].(map[string]interface{})[key[0]]))
			default:
				continue
			}
		}
	}

	return sum
}

func (c Collection) Min(key ...string) decimal.Decimal {

	var smallest = decimal.New(0, 0)

	if len(key) == 0 {
		for i := 0; i < len(c); i++ {
			switch c[i].(type) {
			case decimal.Decimal:
				if i == 0 {
					smallest = c[i].(decimal.Decimal)
					continue
				}
				if smallest.GreaterThan(c[i].(decimal.Decimal)) {
					smallest = c[i].(decimal.Decimal)
				}
			default:
				continue
			}
		}
	} else {
		for i := 0; i < len(c); i++ {
			switch c[i].(type) {
			case map[string]interface{}:
				number := NewDecimalFromInterface(c[i].(map[string]interface{})[key[0]])
				if i == 0 {
					smallest = number
					continue
				}
				if smallest.GreaterThan(number) {
					smallest = number
				}
			default:
				continue
			}
		}
	}

	return smallest
}

// 以上几个函数都是同样的模板，能不能抽出来变一个函数呢
func (c Collection) Max(key ...string) decimal.Decimal {
	var biggest = decimal.New(0, 0)

	if len(key) == 0 {
		for i := 0; i < len(c); i++ {
			switch c[i].(type) {
			case decimal.Decimal:
				if i == 0 {
					biggest = c[i].(decimal.Decimal)
					continue
				}
				if biggest.LessThan(c[i].(decimal.Decimal)) {
					biggest = c[i].(decimal.Decimal)
				}
			default:
				continue
			}
		}
	} else {
		for i := 0; i < len(c); i++ {
			switch c[i].(type) {
			case map[string]interface{}:
				number := NewDecimalFromInterface(c[i].(map[string]interface{})[key[0]])
				if i == 0 {
					biggest = number
					continue
				}
				if biggest.LessThan(number) {
					biggest = number
				}
			default:
				continue
			}
		}
	}

	return biggest
}

func (c Collection) Mode(key ...string) []interface{} {
	panic("implement it")
}

func (c Collection) Only(keys []interface{}) Collection {
	panic("implement it")
}

func (c Collection) Pluck(key interface{}) []interface{} {
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