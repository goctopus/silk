package collection

import "github.com/shopspring/decimal"

type MapArrayCollection struct {
	value []map[string]interface{}
	BaseCollection
}

func (c MapArrayCollection) Sum(key ...string) decimal.Decimal {
	var sum = decimal.New(0, 0)

	for i := 0; i < len(c.value); i++ {
		sum = sum.Add(newDecimalFromInterface(c.value[i][key[0]]))
	}

	return sum
}

func (c MapArrayCollection) Min(key ...string) decimal.Decimal {

	var (
		smallest = decimal.New(0, 0)
		number   decimal.Decimal
	)

	for i := 0; i < len(c.value); i++ {
		number = newDecimalFromInterface(c.value[i][key[0]])
		if i == 0 {
			smallest = number
			continue
		}
		if smallest.GreaterThan(number) {
			smallest = number
		}
	}

	return smallest
}

func (c MapArrayCollection) Max(key ...string) decimal.Decimal {

	var (
		biggest = decimal.New(0, 0)
		number  decimal.Decimal
	)

	for i := 0; i < len(c.value); i++ {
		number = newDecimalFromInterface(c.value[i][key[0]])
		if i == 0 {
			biggest = number
			continue
		}
		if biggest.LessThan(number) {
			biggest = number
		}
	}

	return biggest
}

func (c MapArrayCollection) Pluck(key string) Collection {
	var s = make([]interface{}, 0)
	for i := 0; i < len(c.value); i++ {
		s = append(s, c.value[i][key])
	}
	return Collect(s)
}

func (c MapArrayCollection) Only(keys []string) Collection {
	var d MapArrayCollection

	var ma = make([]map[string]interface{}, 0)
	for _, k := range keys {
		m := make(map[string]interface{}, 0)
		for _, v := range c.value {
			m[k] = v[k]
		}
		ma = append(ma, m)
	}
	d.value = ma
	d.length = len(ma)

	return d
}

func (c MapArrayCollection) Splice(index, length int, new interface{}) Collection {
	var d MapArrayCollection

	n := c.value
	if new != nil {
		if value, ok := new.([]map[string]interface{}); ok {
			m := n[index+length:]
			n = append(n[:index], value...)
			n = append(n, m...)
		} else {
			panic("new's type is wrong")
		}
	} else {
		n = append(n[:index], n[index+length:]...)
	}

	d.value = n
	d.length = len(n)

	return d
}

func (c MapArrayCollection) Take(num int) Collection {
	var d MapArrayCollection
	if num > c.length {
		panic("Not enough elements to take")
	}

	if num >= 0 {
		d.value = c.value[:num]
		d.length = num
	} else {
		d.value = c.value[len(c.value)+num:]
		d.length = 0 - num
	}

	return d
}

func (c MapArrayCollection) All() []interface{} {
	s := make([]interface{}, len(c.value))
	for i := 0; i < len(c.value); i++ {
		s[i] = c.value[i]
	}

	return s
}

func (c MapArrayCollection) Mode(key ...string) []interface{} {
	valueCount := make(map[interface{}]int)
	for i := 0; i < c.length; i++ {
		if v, ok := c.value[i][key[0]]; ok {
			valueCount[v]++
		}
	}

	maxCount := 0
	maxValue := make([]interface{}, len(valueCount))
	for v, c := range valueCount {
		switch {
		case c < maxCount:
			continue
		case c == maxCount:
			maxValue = append(maxValue, v)
		case c > maxCount:
			maxValue = append([]interface{}{}, v)
			maxCount = c
		}
	}
	return maxValue
}

func (c MapArrayCollection) ToMapArray() []map[string]interface{} {
	return c.value
}

func (c MapArrayCollection) Chunk(num int) MultiDimensionalArrayCollection {
	var d MultiDimensionalArrayCollection
	d.length = c.length/num + 1
	d.value = make([][]interface{}, d.length)

	count := 0
	for i := 1; i <= c.length; i++ {
		switch {
		case i == c.length:
			if i%num == 0 {
				d.value[count] = c.All()[i-num:]
				d.value = d.value[:d.length-1]
			} else {
				d.value[count] = c.All()[i-i%num:]
			}
		case i%num != 0 || i < num:
			continue
		default:
			d.value[count] = c.All()[i-num : i]
			count++
		}
	}

	return d
}
