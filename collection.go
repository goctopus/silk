package silk

type Collection map[string]interface{}

type Collections []Collection

func GetCollections(src []map[string]interface{}) Collections {
	var c Collections
	for _,v := range src {
		c = append(c, v)
	}
	return c
}

func (c Collections) Min(key string) interface{} {
	panic("implement it")
}

func (c Collections) Max(key string) interface{} {
	panic("implement it")
}

func (c Collections) Mode(key string) []interface{} {
	panic("implement it")
}

func (c Collections) Only(keys []string) Collections {
	panic("implement it")
}

func (c Collections) Pluck(key string) []interface{} {
	panic("implement it")
}

func (c Collections) Prepend(key string, value interface{}) Collections {
	panic("implement it")
}

func (c Collections) Pull(key string) Collections {
	panic("implement it")
}

func (c Collections) Put(key string, value interface{}) Collections {
	panic("implement it")
}

func (c Collections) SortBy(key string) Collections {
	panic("implement it")
}

func (c Collections) Spice(index int) Collections {
	panic("implement it")
}

func (c Collections) Sum(key string) interface{} {
	panic("implement it")
}

func (c Collections) Take(num int) Collections {
	panic("implement it")
}

func (c Collections) Tojson() string {
	panic("implement it")
}

func (c Collections) Where(key string, value interface{}) Collections {
	panic("implement it")
}