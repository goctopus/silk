package collection

type MapCollection struct {
	value map[string]interface{}
	BaseCollection
}

func (c MapCollection) Only(keys []string) Collection {
	var (
		d MapCollection
		m = make(map[string]interface{}, 0)
	)

	for _, k := range keys {
		m[k] = c.value[k]
	}
	d.value = m
	d.length = len(m)

	return d
}

func (c MapCollection) Prepend(values ...interface{}) Collection {
	var d MapCollection

	d.value = c.value
	d.value[values[0].(string)] = values[1]
	d.length = len(d.value)

	return d
}

func (c MapCollection) Take(num int) Collection {
	var d MapCollection
	if num > c.length {
		panic("Not enough elements to take")
	}

	if num < 0 {
		num = 0 - num
	}
	m := make(map[string]interface{})
	i := 0
	for k, v := range c.value {
		if i == num {
			break
		}
		m[k] = v
		i++
	}
	d.value = m
	d.length = len(m)

	return d
}
