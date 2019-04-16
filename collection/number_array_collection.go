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

	d.value = append([]decimal.Decimal{NewDecimalFromInterface(values[0])}, d.value...)
	d.length = len(d.value)

	return d
}

func (c NumberArrayCollection) Splice(index, length int, new interface{}) Collection {
	var d NumberArrayCollection

	if value, ok := new.([]decimal.Decimal); ok {
		n := c.value
		n = append(n[:index], value...)
		n = append(n, n[index+length:]...)

		d.value = n
		d.length = len(n)
	} else {
		panic("new's type is wrong")
	}

	return d
}