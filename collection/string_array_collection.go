package collection

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
			panic("new's type is wrong")
		}
	} else {
		n = append(n[:index], n[index+length:]...)
	}

	d.value = n
	d.length = len(n)

	return d
}
