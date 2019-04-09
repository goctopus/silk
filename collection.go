package silk

type Collection []interface{}

func Collect(src interface{}) Collection {
	var c Collection

	switch src.(type) {
	case []map[string]interface{}:
		for _,v := range src.([]map[string]interface{}) {
			c = append(c, v)
		}
	default:
		panic("wrong type")
	}

	return c
}

func (c Collection) Min(key interface{}) interface{} {
	panic("implement it")
}

func (c Collection) Max(key interface{}) interface{} {
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

func (c Collection) Sum(key interface{}) interface{} {
	panic("implement it")
}

func (c Collection) Take(num int) Collection {
	panic("implement it")
}

func (c Collection) Tojson() string {
	panic("implement it")
}

func (c Collection) Where(key interface{}, value interface{}) Collection {
	panic("implement it")
}