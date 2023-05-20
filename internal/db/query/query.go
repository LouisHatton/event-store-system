package query

type Matcher string

const (
	EqualTo       Matcher = "=="
	NotEqualTo    Matcher = "!="
	LessThan      Matcher = "<"
	LessThanEQ    Matcher = "<="
	GreaterThan   Matcher = ">"
	GreaterThanEQ Matcher = ">="
	Contains      Matcher = "array-contains"
)

type Where struct {
	Key     string
	Matcher Matcher
	Value   interface{}
}

type Options struct {
	Limit   *int
	Offset  *int
	OrderBy *OrderBy
}

type OrderByDirection string

const (
	OrderAsc  OrderByDirection = "Asc"
	OrderDesc OrderByDirection = "Desc"
)

type OrderBy struct {
	Value     string
	Direction OrderByDirection
}
