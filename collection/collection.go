package collection

import "github.com/shopspring/decimal"

func Collect(src interface{}) Collection {
	switch src.(type) {
	case []string:
		var c StringArrayCollection
		c.value = src.([]string)
		c.length = len(src.([]string))
		return c
	default:
		panic("wrong type")
	}
}

type Collection interface {
	// reference: https://laravel.com/docs/5.8/collections#method-all
	All() []interface{}

	// reference: https://laravel.com/docs/5.8/collections#method-avg
	Avg(key ...string) decimal.Decimal

	// reference: https://laravel.com/docs/5.8/collections#method-sum
	Sum(key ...string) decimal.Decimal

	// reference: https://laravel.com/docs/5.8/collections#method-min
	Min(key ...string) decimal.Decimal

	// reference: https://laravel.com/docs/5.8/collections#method-max
	Max(key ...string) decimal.Decimal

	// reference: https://laravel.com/docs/5.8/collections#method-join
	Join(delimiter string) string

	// reference: https://laravel.com/docs/5.8/collections#method-combine
	Combine(value []interface{}) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-count
	Count() int

	// reference: https://laravel.com/docs/5.8/collections#method-pluck
	Pluck(key string) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-mode
	Mode(key ...string) []interface{}

	// reference: https://laravel.com/docs/5.8/collections#method-only
	Only(keys []string) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-prepend
	Prepend(values ...interface{}) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-pull
	Pull(key interface{}) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-put
	Put(key string, value interface{}) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-sortby
	SortBy(key string) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-spice
	Spice(index int) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-take
	Take(num int) Collection

	// reference: https://laravel.com/docs/5.8/collections#method-average
	Average()

	// reference: https://laravel.com/docs/5.8/collections#method-chunk
	Chunk()

	// reference: https://laravel.com/docs/5.8/collections#method-collapse
	Collapse()

	// reference: https://laravel.com/docs/5.8/collections#method-concat
	Concat()

	// reference: https://laravel.com/docs/5.8/collections#method-contains
	Contains()

	// reference: https://laravel.com/docs/5.8/collections#method-containsStrict
	ContainsStrict()

	// reference: https://laravel.com/docs/5.8/collections#method-countBy
	CountBy()

	// reference: https://laravel.com/docs/5.8/collections#method-crossJoin
	CrossJoin()

	// reference: https://laravel.com/docs/5.8/collections#method-dd
	Dd()

	// reference: https://laravel.com/docs/5.8/collections#method-diff
	Diff()

	// reference: https://laravel.com/docs/5.8/collections#method-diffAssoc
	DiffAssoc()

	// reference: https://laravel.com/docs/5.8/collections#method-diffKeys
	DiffKeys()

	// reference: https://laravel.com/docs/5.8/collections#method-dump
	Dump()

	// reference: https://laravel.com/docs/5.8/collections#method-each
	Each()

	// reference: https://laravel.com/docs/5.8/collections#method-eachSpread
	EachSpread()

	// reference: https://laravel.com/docs/5.8/collections#method-every
	Every()

	// reference: https://laravel.com/docs/5.8/collections#method-except
	Except()

	// reference: https://laravel.com/docs/5.8/collections#method-filter
	Filter()

	// reference: https://laravel.com/docs/5.8/collections#method-first
	First()

	// reference: https://laravel.com/docs/5.8/collections#method-firstWhere
	FirstWhere()

	// reference: https://laravel.com/docs/5.8/collections#method-flatMap
	FlatMap()

	// reference: https://laravel.com/docs/5.8/collections#method-flatten
	Flatten()

	// reference: https://laravel.com/docs/5.8/collections#method-flip
	Flip()

	// reference: https://laravel.com/docs/5.8/collections#method-forget
	Forget()

	// reference: https://laravel.com/docs/5.8/collections#method-forPage
	ForPage()

	// reference: https://laravel.com/docs/5.8/collections#method-get
	Get()

	// reference: https://laravel.com/docs/5.8/collections#method-groupBy
	GroupBy()

	// reference: https://laravel.com/docs/5.8/collections#method-has
	Has()

	// reference: https://laravel.com/docs/5.8/collections#method-implode
	Implode()

	// reference: https://laravel.com/docs/5.8/collections#method-intersect
	Intersect()

	// reference: https://laravel.com/docs/5.8/collections#method-intersectByKeys
	IntersectByKeys()

	// reference: https://laravel.com/docs/5.8/collections#method-isEmpty
	IsEmpty()

	// reference: https://laravel.com/docs/5.8/collections#method-isNotEmpty
	IsNotEmpty()

	// reference: https://laravel.com/docs/5.8/collections#method-keyBy
	KeyBy()

	// reference: https://laravel.com/docs/5.8/collections#method-keys
	Keys()

	// reference: https://laravel.com/docs/5.8/collections#method-last
	Last()

	// reference: https://laravel.com/docs/5.8/collections#method-macro
	Macro()

	// reference: https://laravel.com/docs/5.8/collections#method-make
	Make()

	// reference: https://laravel.com/docs/5.8/collections#method-map
	Map()

	// reference: https://laravel.com/docs/5.8/collections#method-mapInto
	MapInto()

	// reference: https://laravel.com/docs/5.8/collections#method-mapSpread
	MapSpread()

	// reference: https://laravel.com/docs/5.8/collections#method-mapToGroups
	MapToGroups()

	// reference: https://laravel.com/docs/5.8/collections#method-mapWithKeys
	MapWithKeys()

	// reference: https://laravel.com/docs/5.8/collections#method-median
	Median()

	// reference: https://laravel.com/docs/5.8/collections#method-merge
	Merge()

	// reference: https://laravel.com/docs/5.8/collections#method-nth
	Nth()

	// reference: https://laravel.com/docs/5.8/collections#method-pad
	Pad()

	// reference: https://laravel.com/docs/5.8/collections#method-partition
	Partition()

	// reference: https://laravel.com/docs/5.8/collections#method-pipe
	Pipe()

	// reference: https://laravel.com/docs/5.8/collections#method-pop
	Pop()

	// reference: https://laravel.com/docs/5.8/collections#method-push
	Push()

	// reference: https://laravel.com/docs/5.8/collections#method-random
	Random()

	// reference: https://laravel.com/docs/5.8/collections#method-reduce
	Reduce()

	// reference: https://laravel.com/docs/5.8/collections#method-reject
	Reject()

	// reference: https://laravel.com/docs/5.8/collections#method-reverse
	Reverse()

	// reference: https://laravel.com/docs/5.8/collections#method-search
	Search()

	// reference: https://laravel.com/docs/5.8/collections#method-shift
	Shift()

	// reference: https://laravel.com/docs/5.8/collections#method-shuffle
	Shuffle()

	// reference: https://laravel.com/docs/5.8/collections#method-slice
	Slice()

	// reference: https://laravel.com/docs/5.8/collections#method-some
	Some()

	// reference: https://laravel.com/docs/5.8/collections#method-sort
	Sort()

	// reference: https://laravel.com/docs/5.8/collections#method-sortByDesc
	SortByDesc()

	// reference: https://laravel.com/docs/5.8/collections#method-sortKeys
	SortKeys()

	// reference: https://laravel.com/docs/5.8/collections#method-sortKeysDesc
	SortKeysDesc()

	// reference: https://laravel.com/docs/5.8/collections#method-splice
	Splice()

	// reference: https://laravel.com/docs/5.8/collections#method-split
	Split()

	// reference: https://laravel.com/docs/5.8/collections#method-tap
	Tap()

	// reference: https://laravel.com/docs/5.8/collections#method-times
	Times()

	// reference: https://laravel.com/docs/5.8/collections#method-transform
	Transform()

	// reference: https://laravel.com/docs/5.8/collections#method-union
	Union()

	// reference: https://laravel.com/docs/5.8/collections#method-unique
	Unique()

	// reference: https://laravel.com/docs/5.8/collections#method-uniqueStrict
	UniqueStrict()

	// reference: https://laravel.com/docs/5.8/collections#method-unless
	Unless()

	// reference: https://laravel.com/docs/5.8/collections#method-unlessEmpty
	UnlessEmpty()

	// reference: https://laravel.com/docs/5.8/collections#method-unlessNotEmpty
	UnlessNotEmpty()

	// reference: https://laravel.com/docs/5.8/collections#method-unwrap
	Unwrap()

	// reference: https://laravel.com/docs/5.8/collections#method-values
	Values()

	// reference: https://laravel.com/docs/5.8/collections#method-when
	When()

	// reference: https://laravel.com/docs/5.8/collections#method-whenEmpty
	WhenEmpty()

	// reference: https://laravel.com/docs/5.8/collections#method-whenNotEmpty
	WhenNotEmpty()

	// reference: https://laravel.com/docs/5.8/collections#method-whereStrict
	WhereStrict()

	// reference: https://laravel.com/docs/5.8/collections#method-whereBetween
	WhereBetween()

	// reference: https://laravel.com/docs/5.8/collections#method-whereIn
	WhereIn()

	// reference: https://laravel.com/docs/5.8/collections#method-whereInStrict
	WhereInStrict()

	// reference: https://laravel.com/docs/5.8/collections#method-whereInstanceOf
	WhereInstanceOf()

	// reference: https://laravel.com/docs/5.8/collections#method-whereNotBetween
	WhereNotBetween()

	// reference: https://laravel.com/docs/5.8/collections#method-whereNotIn
	WhereNotIn()

	// reference: https://laravel.com/docs/5.8/collections#method-whereNotInStrict
	WhereNotInStrict()

	// reference: https://laravel.com/docs/5.8/collections#method-wrap
	Wrap()

	// reference: https://laravel.com/docs/5.8/collections#method-zip
	Zip()

	ToJson() string

	ToNumberArray() []decimal.Decimal

	ToStringArray() []string

	ToMap() map[string]interface{}

	ToMapArray() []map[string]interface{}

	Where(key string, value interface{}) Collection
}

func NewDecimalFromInterface(a interface{}) decimal.Decimal {
	var d decimal.Decimal

	switch a.(type) {
	case uint:
		d = decimal.New(int64(a.(uint)), 0)
	case uint8:
		d = decimal.New(int64(a.(uint8)), 0)
	case uint16:
		d = decimal.New(int64(a.(uint16)), 0)
	case uint32:
		d = decimal.New(int64(a.(uint32)), 0)
	case uint64:
		d = decimal.New(int64(a.(uint64)), 0)
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
	default:
		panic("wrong type")
	}

	return d
}
