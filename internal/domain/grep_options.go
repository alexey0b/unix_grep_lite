package domain

type GrepOptions struct {
	NumAfter      int
	NumBefore     int
	NumAround     int
	AfterContext  bool
	BeforeContext bool
	AroundContext bool
	Count         bool
	IgnoreCase    bool
	InvertMatch   bool
	FixedStrings  bool
	LineNumber    bool
}
