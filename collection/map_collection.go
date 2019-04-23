package collection

import (
	"fmt"
	"strconv"
)

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

	var m = copyMap(c.value)
	m[values[0].(string)] = values[1]

	return MapCollection{m, BaseCollection{length: len(m)}}
}

func (c MapCollection) ToMap() map[string]interface{} {
	return c.value
}

func (c MapCollection) Contains(value interface{}, key ...interface{}) bool {
	t := fmt.Sprintf("%T&%T", c.value, value)
	switch {
	case t == "[]map[string]string&int":
		return parseContainsKey(c.value, strconv.Itoa(value.(int)), key)
	case t == "[]map[string]string&int64":
		return parseContainsKey(c.value, strconv.FormatInt(value.(int64), 10), key)
	default:
		return parseContainsKey(c.value, value, key)
	}
}

func (c MapCollection) ContainsStrict(value interface{}, key ...interface{}) bool {
	return parseContainsKey(c.value, value, key)
}
