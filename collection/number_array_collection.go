package collection

import "github.com/shopspring/decimal"

type NumberArrayCollection struct {
	value []decimal.Decimal
	BaseCollection
}

func (c NumberArrayCollection) Sum(key ...string) decimal.Decimal {

	var sum = decimal.New(0, 0)

	for i := 0; i < len(c.value); i++ {
		sum = sum.Add(c.value[i])
	}

	return sum
}

func (c NumberArrayCollection) Min(key ...string) decimal.Decimal {

	var smallest = decimal.New(0, 0)

	for i := 0; i < len(c.value); i++ {
		if i == 0 {
			smallest = c.value[i]
			continue
		}
		if smallest.GreaterThan(c.value[i]) {
			smallest = c.value[i]
		}
	}

	return smallest
}

func (c NumberArrayCollection) Max(key ...string) decimal.Decimal {

	var biggest = decimal.New(0, 0)

	for i := 0; i < len(c.value); i++ {
		if i == 0 {
			biggest = c.value[i]
			continue
		}
		if biggest.LessThan(c.value[i]) {
			biggest = c.value[i]
		}
	}

	return biggest
}

func (c NumberArrayCollection) Prepend(values ...interface{}) Collection {
	var d NumberArrayCollection

	d.value = append([]decimal.Decimal{newDecimalFromInterface(values[0])}, d.value...)
	d.length = len(d.value)

	return d
}

func (c NumberArrayCollection) Splice(index, length int, new interface{}) Collection {
	var d NumberArrayCollection

	n := c.value
	if new != nil {
		if value, ok := new.([]decimal.Decimal); ok {
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

func (c NumberArrayCollection) Take(num int) Collection {
	var d NumberArrayCollection
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

func (c NumberArrayCollection) All() []interface{} {
	s := make([]interface{}, len(c.value))
	for i := 0; i < len(c.value); i++ {
		s[i] = c.value[i]
	}

	return s
}

// Type of slice use "" as parameter
func (c NumberArrayCollection) Mode(key ...string) []interface{} {
	valueCount := make(map[float64]int)
	for _, v := range c.value {
		f, _ := v.Float64()
		valueCount[f]++
	}

	maxCount := 0
	maxValue := make([]interface{}, len(valueCount))
	for v, c := range valueCount {
		switch {
		case c < maxCount:
			continue
		case c == maxCount:
			maxValue = append(maxValue, newDecimalFromInterface(v))
		case c > maxCount:
			maxValue = append([]interface{}{}, newDecimalFromInterface(v))
			maxCount = c
		}
	}
	return maxValue
}

func (c NumberArrayCollection) ToNumberArray() []decimal.Decimal {
	return c.value
}

func (c NumberArrayCollection) Chunk(num int) MultiDimensionalArrayCollection {
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
