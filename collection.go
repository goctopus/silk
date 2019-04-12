package silk

import (
	"github.com/shopspring/decimal"
)

// Collection is a data structure contains useful functions. It is immutable.
//
// interface{} must be one type of:
//
// []decimal.Decimal,
// []string,
// []map[string]interface{},
// map[string]interface{}
//
type Collection struct {
	value  interface{}
	length int
}

func Collect(src interface{}) Collection {
	var c Collection

	switch src.(type) {
	case []map[string]interface{}:
		c.value = src
		c.length = len(src.([]map[string]interface{}))
	case []string:
		c.value = src
		c.length = len(src.([]string))
	case map[string]interface{}:
		c.value = src
		c.length = len(src.(map[string]interface{}))
	case []int:
		var d = make([]decimal.Decimal, 0)
		for _, v := range src.([]int) {
			d = append(d, decimal.New(int64(v), 0))
		}
		c.value = d
		c.length = len(src.([]int))
	case []int8:
		var d = make([]decimal.Decimal, 0)
		for _, v := range src.([]int8) {
			d = append(d, decimal.New(int64(v), 0))
		}
		c.value = d
		c.length = len(src.([]int8))
	case []int16:
		var d = make([]decimal.Decimal, 0)
		for _, v := range src.([]int16) {
			d = append(d, decimal.New(int64(v), 0))
		}
		c.value = d
		c.length = len(src.([]int16))
	case []int32:
		var d = make([]decimal.Decimal, 0)
		for _, v := range src.([]int32) {
			d = append(d, decimal.New(int64(v), 0))
		}
		c.value = d
		c.length = len(src.([]int32))
	case []int64:
		var d = make([]decimal.Decimal, 0)
		for _, v := range src.([]int64) {
			d = append(d, decimal.New(v, 0))
		}
		c.value = d
		c.length = len(src.([]int64))
	case []float32:
		var f = make([]decimal.Decimal, 0)
		for _, v := range src.([]float32) {
			f = append(f, decimal.NewFromFloat32(v))
		}
		c.value = f
		c.length = len(src.([]float32))
	case []float64:
		var f = make([]decimal.Decimal, 0)
		for _, v := range src.([]float64) {
			f = append(f, decimal.NewFromFloat(v))
		}
		c.value = f
		c.length = len(src.([]float64))
	case []interface{}:
		s := src.([]interface{})
		if len(s) > 0 {
			switch s[0].(type) {
			case string:
				var ss = make([]string, 0)
				for _, v := range s {
					ss = append(ss, v.(string))
				}
				c.value = ss
				c.length = len(ss)
			default:
				var d = make([]decimal.Decimal, 0)
				for _, v := range s {
					d = append(d, NewDecimalFromInterface(v))
				}
				c.value = d
				c.length = len(d)
			}
		}
	default:
		panic("wrong type")
	}

	return c
}

func (c Collection) All() []interface{} {
	panic("implement it")
}

func (c Collection) Avg(key ...string) decimal.Decimal {
	return c.Sum(key...).Div(decimal.New(int64(c.length), 0))
}

func (c Collection) Sum(key ...string) decimal.Decimal {
	var sum = decimal.New(0, 0)

	if len(key) == 0 {
		if n, ok := c.value.([]decimal.Decimal); ok {
			for i := 0; i < len(n); i++ {
				sum = sum.Add(n[i])
			}
		}
	} else {
		if n, ok := c.value.([]map[string]interface{}); ok {
			for i := 0; i < len(n); i++ {
				sum = sum.Add(NewDecimalFromInterface(n[i][key[0]]))
			}
		}
	}

	return sum
}

func (c Collection) Min(key ...string) decimal.Decimal {

	var smallest = decimal.New(0, 0)

	if len(key) == 0 {
		if n, ok := c.value.([]decimal.Decimal); ok {
			for i := 0; i < len(n); i++ {
				if i == 0 {
					smallest = n[i]
					continue
				}
				if smallest.GreaterThan(n[i]) {
					smallest = n[i]
				}
			}
		}
	} else {

		var number decimal.Decimal

		if n, ok := c.value.([]map[string]interface{}); ok {
			for i := 0; i < len(n); i++ {
				number = NewDecimalFromInterface(n[i][key[0]])
				if i == 0 {
					smallest = number
					continue
				}
				if smallest.GreaterThan(number) {
					smallest = number
				}
			}
		}
	}

	return smallest
}

func (c Collection) Max(key ...string) decimal.Decimal {

	var biggest = decimal.New(0, 0)

	if len(key) == 0 {
		if n, ok := c.value.([]decimal.Decimal); ok {
			for i := 0; i < len(n); i++ {
				if i == 0 {
					biggest = n[i]
					continue
				}
				if biggest.LessThan(n[i]) {
					biggest = n[i]
				}
			}
		}
	} else {

		var number decimal.Decimal

		if n, ok := c.value.([]map[string]interface{}); ok {
			for i := 0; i < len(n); i++ {
				number = NewDecimalFromInterface(n[i][key[0]])
				if i == 0 {
					biggest = number
					continue
				}
				if biggest.LessThan(number) {
					biggest = number
				}
			}
		}
	}

	return biggest
}

func (c Collection) Join(delimiter string) string {
	s := ""

	if n, ok := c.value.([]string); ok {
		for i := 0; i < len(n); i++ {
			if i != len(n)-1 {
				s += n[i] + delimiter
			} else {
				s += n[i]
			}
		}
	} else if n, ok := c.value.([]decimal.Decimal); ok {
		for i := 0; i < len(n); i++ {
			if i != len(n)-1 {
				s += n[i].String() + delimiter
			} else {
				s += n[i].String()
			}
		}
	}

	return s
}

func (c Collection) Combine(value []interface{}) Collection {
	var (
		m      = make(map[string]interface{}, 0)
		length = c.length
		d      Collection
	)

	if length > len(value) {
		length = len(value)
	}

	if n, ok := c.value.([]string); ok {
		for i := 0; i < length; i++ {
			m[n[i]] = value[i]
		}
	}

	d.value = m
	d.length = len(m)

	return d
}

func (c Collection) Count() int {
	return c.length
}

func (c Collection) Pluck(key string) Collection {
	var s = make([]interface{}, 0)
	if n, ok := c.value.([]map[string]interface{}); ok {
		for i := 0; i < len(n); i++ {
			s = append(s, n[i][key])
		}
	}
	return Collect(s)
}

// reference: https://laravel.com/docs/5.8/collections#method-mode
func (c Collection) Mode(key ...string) []interface{} {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-only
func (c Collection) Only(keys []string) Collection {
	var d Collection
	if n, ok := c.value.(map[string]interface{}); ok {
		var m = make(map[string]interface{}, 0)
		for _, k := range keys {
			m[k] = n[k]
		}
		d.value = m
		d.length = len(m)
	} else if n, ok := c.value.([]map[string]interface{}); ok {
		var ma = make([]map[string]interface{}, 0)
		for _, k := range keys {
			m := make(map[string]interface{}, 0)
			for _, v := range n {
				m[k] = v[k]
			}
			ma = append(ma, m)
		}
		d.value = ma
		d.length = len(ma)
	}

	return d
}

// reference: https://laravel.com/docs/5.8/collections#method-prepend
func (c Collection) Prepend(values ...interface{}) Collection {
	var d Collection
	if len(values) == 1 {
		if n, ok := c.value.([]string); ok {
			n = append([]string{values[0].(string)}, n...)
			d.value = n
			d.length = len(n)
		} else if n, ok := c.value.([]decimal.Decimal); ok {
			n = append([]decimal.Decimal{NewDecimalFromInterface(values[0])}, n...)
			d.value = n
			d.length = len(n)
		}
	} else if len(values) == 2 {
		if n, ok := c.value.(map[string]interface{}); ok {
			n[values[0].(string)] = values[1]
			d.value = n
			d.length = len(n)
		}
	} else {
		panic("wrong parameter")
	}

	return d
}

// reference: https://laravel.com/docs/5.8/collections#method-pull
func (c Collection) Pull(key interface{}) Collection {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-put
func (c Collection) Put(key string, value interface{}) Collection {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-sortby
func (c Collection) SortBy(key string) Collection {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-spice
func (c Collection) Spice(index int) Collection {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-take
func (c Collection) Take(num int) Collection {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-average
func (c Collection) Average() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-chunk
func (c Collection) Chunk() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-collapse
func (c Collection) Collapse() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-concat
func (c Collection) Concat() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-contains
func (c Collection) Contains() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-containsStrict
func (c Collection) ContainsStrict() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-countBy
func (c Collection) CountBy() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-crossJoin
func (c Collection) CrossJoin() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-dd
func (c Collection) Dd() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-diff
func (c Collection) Diff() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-diffAssoc
func (c Collection) DiffAssoc() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-diffKeys
func (c Collection) DiffKeys() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-dump
func (c Collection) Dump() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-each
func (c Collection) Each() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-eachSpread
func (c Collection) EachSpread() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-every
func (c Collection) Every() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-except
func (c Collection) Except() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-filter
func (c Collection) Filter() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-first
func (c Collection) First() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-firstWhere
func (c Collection) FirstWhere() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-flatMap
func (c Collection) FlatMap() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-flatten
func (c Collection) Flatten() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-flip
func (c Collection) Flip() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-forget
func (c Collection) Forget() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-forPage
func (c Collection) ForPage() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-get
func (c Collection) Get() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-groupBy
func (c Collection) GroupBy() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-has
func (c Collection) Has() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-implode
func (c Collection) Implode() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-intersect
func (c Collection) Intersect() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-intersectByKeys
func (c Collection) IntersectByKeys() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-isEmpty
func (c Collection) IsEmpty() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-isNotEmpty
func (c Collection) IsNotEmpty() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-keyBy
func (c Collection) KeyBy() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-keys
func (c Collection) Keys() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-last
func (c Collection) Last() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-macro
func (c Collection) Macro() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-make
func (c Collection) Make() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-map
func (c Collection) Map() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-mapInto
func (c Collection) MapInto() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-mapSpread
func (c Collection) MapSpread() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-mapToGroups
func (c Collection) MapToGroups() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-mapWithKeys
func (c Collection) MapWithKeys() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-median
func (c Collection) Median() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-merge
func (c Collection) Merge() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-nth
func (c Collection) Nth() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-pad
func (c Collection) Pad() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-partition
func (c Collection) Partition() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-pipe
func (c Collection) Pipe() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-pop
func (c Collection) Pop() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-push
func (c Collection) Push() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-random
func (c Collection) Random() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-reduce
func (c Collection) Reduce() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-reject
func (c Collection) Reject() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-reverse
func (c Collection) Reverse() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-search
func (c Collection) Search() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-shift
func (c Collection) Shift() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-shuffle
func (c Collection) Shuffle() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-slice
func (c Collection) Slice() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-some
func (c Collection) Some() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-sort
func (c Collection) Sort() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-sortByDesc
func (c Collection) SortByDesc() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-sortKeys
func (c Collection) SortKeys() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-sortKeysDesc
func (c Collection) SortKeysDesc() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-splice
func (c Collection) Splice() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-split
func (c Collection) Split() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-tap
func (c Collection) Tap() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-times
func (c Collection) Times() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-transform
func (c Collection) Transform() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-union
func (c Collection) Union() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-unique
func (c Collection) Unique() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-uniqueStrict
func (c Collection) UniqueStrict() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-unless
func (c Collection) Unless() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-unlessEmpty
func (c Collection) UnlessEmpty() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-unlessNotEmpty
func (c Collection) UnlessNotEmpty() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-unwrap
func (c Collection) Unwrap() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-values
func (c Collection) Values() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-when
func (c Collection) When() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-whenEmpty
func (c Collection) WhenEmpty() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-whenNotEmpty
func (c Collection) WhenNotEmpty() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-whereStrict
func (c Collection) WhereStrict() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-whereBetween
func (c Collection) WhereBetween() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-whereIn
func (c Collection) WhereIn() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-whereInStrict
func (c Collection) WhereInStrict() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-whereInstanceOf
func (c Collection) WhereInstanceOf() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-whereNotBetween
func (c Collection) WhereNotBetween() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-whereNotIn
func (c Collection) WhereNotIn() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-whereNotInStrict
func (c Collection) WhereNotInStrict() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-wrap
func (c Collection) Wrap() {
	panic("implement it")
}

// reference: https://laravel.com/docs/5.8/collections#method-zip
func (c Collection) Zip() {
	panic("implement it")
}

func (c Collection) ToJson() string {
	panic("implement it")
}

func (c Collection) ToNumberArray() []decimal.Decimal {
	if m, ok := c.value.([]decimal.Decimal); ok {
		return m
	} else {
		return []decimal.Decimal{}
	}
}

func (c Collection) ToStringArray() []string {
	if m, ok := c.value.([]string); ok {
		return m
	} else {
		return []string{}
	}
}

func (c Collection) ToMap() map[string]interface{} {
	if m, ok := c.value.(map[string]interface{}); ok {
		return m
	} else {
		return map[string]interface{}{}
	}
}

func (c Collection) ToMapArray() []map[string]interface{} {
	if m, ok := c.value.([]map[string]interface{}); ok {
		return m
	} else {
		return []map[string]interface{}{}
	}
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
