package collection

import "fmt"

type StringArrayCollection struct {
	value []string
	BaseCollection
}

func (c StringArrayCollection) Join(delimiter string) string {
	s := ""
	for i := 0; i < len(c.value); i++ {
		if i != len(c.value)-1 {
			s += c.value[i] + delimiter
		} else {
			s += c.value[i]
		}
	}
	return s
}

func (c StringArrayCollection) Combine(value []interface{}) Collection {
	var (
		m      = make(map[string]interface{}, 0)
		length = c.length
		d      MapCollection
	)

	if length > len(value) {
		length = len(value)
	}

	for i := 0; i < length; i++ {
		m[c.value[i]] = value[i]
	}

	d.value = m
	d.length = len(m)

	return d
}

func (c StringArrayCollection) Prepend(values ...interface{}) Collection {
	var d StringArrayCollection

	d.value = append([]string{values[0].(string)}, c.value...)
	d.length = len(d.value)

	return d
}

func (c StringArrayCollection) Splice(index, length int, new interface{}) Collection {
	var d StringArrayCollection

	n := c.value
	if new != nil {
		if value, ok := new.([]string); ok {
			m := n[index+length:]
			n = append(n[:index], value...)
			n = append(n, m...)
		} else {
			panic(fmt.Sprintf("invalid argument: %v", new))
		}
	} else {
		n = append(n[:index], n[index+length:]...)
	}

	d.value = n
	d.length = len(n)

	return d
}

func (c StringArrayCollection) Take(num int) Collection {
	var d StringArrayCollection
	if num > c.length {
		panic("not enough elements to take")
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

func (c StringArrayCollection) All() []interface{} {
	s := make([]interface{}, len(c.value))
	for i := 0; i < len(c.value); i++ {
		s[i] = c.value[i]
	}

	return s
}

// Type of slice use "" as parameter
func (c StringArrayCollection) Mode(key ...string) []interface{} {
	valueCount := make(map[string]int)
	for _, v := range c.value {
		valueCount[v]++
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

func (c StringArrayCollection) ToStringArray() []string {
	return c.value
}
